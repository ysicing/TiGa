// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package debug

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/pkg/factory"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	tcpingExample = templates.Examples(`
  # simple tcping
  tiga debug tcping [-4] [-6] [-n count] [-t timeout] address port
  `)
)

func TcpingCommand(f factory.Factory) *cobra.Command {
	var ipv4, ipv6 bool
	var count, timeout int
	var stopPing chan bool
	cmd := &cobra.Command{
		Use:     "tcping",
		Short:   "tcping",
		Example: tcpingExample,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if ipv4 && ipv6 {
				return errors.New("ipv4 and ipv6 can't be used together")
			}
			if len(args) < 2 {
				return errors.New("address and port is required")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			address := args[0]
			port := args[1]
			if ipv4 || (!ipv6 && isIPv4(address)) {
				address, err = resolveAddress(address, "ipv4")
			} else if ipv6 || isIPv6(address) {
				address, err = resolveAddress(address, "ipv6")
			} else {
				// Default to IPv4 if no -4 or -6 flags specified and address is not explicitly IPv6
				address, err = resolveAddress(address, "ipv4")
			}
			if err != nil {
				return err
			}
			f.GetLog().Infof("start ping %s:%s...", address, port)
			stopPing = make(chan bool, 1)
			interrupt := make(chan os.Signal, 1)
			signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
			var sentCount int
			var respondedCount int
			var minTime, maxTime, totalResponseTime int64
			go func() {
				for i := 0; count == 0 || i < count; i++ {
					select {
					case <-stopPing:
						return
					default:
						start := time.Now()
						conn, err := net.DialTimeout("tcp", address+":"+port, time.Duration(timeout)*time.Second)
						elapsed := time.Since(start).Milliseconds()

						sentCount++
						if err != nil {
							f.GetLog().Warnf("Failed to connect to %s:%s: %v", address, port, err)
						} else {
							conn.Close()
							respondedCount++
							if respondedCount == 1 || elapsed < minTime {
								minTime = elapsed
							}
							if elapsed > maxTime {
								maxTime = elapsed
							}
							totalResponseTime += elapsed
							fmt.Printf("tcping %s:%s in %dms\n", address, port, elapsed)
						}

						if count != 0 && i == count-1 {
							break
						}

						time.Sleep(time.Duration(timeout) * time.Second)
					}
				}
				stopPing <- true
			}()

			select {
			case <-interrupt:
				f.GetLog().Warn("ping interrupted.")
				stopPing <- true
			case <-stopPing:
				f.GetLog().Done("ping stopped.")
			}
			printTcpingStatistics(sentCount, respondedCount, minTime, maxTime, totalResponseTime)
			return nil
		},
	}
	cmd.Flags().BoolVarP(&ipv4, "ipv4", "4", false, "ipv4")
	cmd.Flags().BoolVarP(&ipv6, "ipv6", "6", false, "ipv6")
	cmd.Flags().IntVarP(&count, "count", "c", 3, "number of pings")
	cmd.Flags().IntVarP(&timeout, "timeout", "t", 3, "time interval between pings in seconds ")
	return cmd
}

func isIPv4(address string) bool {
	return strings.Count(address, ":") == 0
}

// resolveAddress resolves the address to the specified network type
func isIPv6(address string) bool {
	return strings.Count(address, ":") >= 2
}

func resolveAddress(address, version string) (string, error) {
	ipList, err := net.LookupIP(address)
	if err != nil {
		return "", errors.Errorf("failed to resolve %s: %v", address, err)
	}

	for _, ip := range ipList {
		if version == "ipv4" && ip.To4() != nil {
			return ip.String(), nil
		} else if version == "ipv6" && ip.To16() != nil && ip.To4() == nil {
			return "[" + ip.String() + "]", nil
		}
	}
	return "", errors.Errorf("no %s addresses found for %s", version, address)
}

func printTcpingStatistics(sentCount, respondedCount int, minTime, maxTime, totalResponseTime int64) {
	fmt.Println("")
	fmt.Println("--- Tcping Statistics ---")
	fmt.Printf("%d tcp ping sent, %d tcp ping responsed, %.2f%% loss\n", sentCount, respondedCount, float64(sentCount-respondedCount)/float64(sentCount)*100)
	if respondedCount > 0 {
		fmt.Printf("min/avg/max = %dms/%dms/%dms\n", minTime, totalResponseTime/int64(respondedCount), maxTime)
	} else {
		fmt.Println("No responses received.")
	}
}
