// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package netutil

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/go-ping/ping"
	"github.com/jackpal/gateway"
	"github.com/miekg/dns"
	"io/fs"
	"net"
	"os"
	"regexp"
	"time"
)

// CheckDefaultRoute checks if the default route is reachable
func CheckDefaultRoute() (net.IP, error) {
	// check default gateway
	gw, err := gateway.DiscoverGateway()

	if err != nil {
		return nil, fmt.Errorf("error reading default route: %s", err)
	}

	if CheckReachabilityWithICMP(gw.String()) {
		return gw, nil
	}

	return gw, fmt.Errorf("default route is unreachable")
}

// CheckReachabilityWithICMP checks if a host is reachable using ICMP
func CheckReachabilityWithICMP(host string) bool {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		return false
	}

	pinger.Count = 3
	pinger.Debug = true
	pinger.Interval = 200 * time.Millisecond
	pinger.Timeout = 3 * time.Second

	err = pinger.Run()

	if err != nil {
		return false
	}

	stats := pinger.Statistics()

	return stats.PacketsRecv != 0
}

// GetDefaultNameserver returns the default nameserver
func GetDefaultNameserver() (string, error) {
	// get default ns from /etc/resolv.conf
	byteString, err := fs.ReadFile(os.DirFS("/etc"), "resolv.conf")
	if err != nil {
		return "", err
	}
	s := string(byteString)
	re := regexp.MustCompile(`(?m)^nameserver( *|\t*)(.*?)$`)
	match := re.FindStringSubmatch(s)

	if len(match) < 2 {
		return "", errors.New("nameserver not found")
	}
	return match[2], nil
}

// CheckNameserverAvailability checks if a nameserver is reachable using DNS
func CheckNameserverAvailability(s string) error {
	c := new(dns.Client)
	c.Dialer = &net.Dialer{
		Timeout: 3 * time.Second,
	}
	m := new(dns.Msg)
	m.SetQuestion("apple.com.", dns.TypeA)
	_, _, err := c.Exchange(m, s)

	if err != nil {
		return err
	}
	return nil
}
