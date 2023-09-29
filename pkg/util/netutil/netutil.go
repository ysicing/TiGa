// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package netutil

import (
	"fmt"
	"io"
	"io/fs"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/ysicing/tiga/pkg/log"

	"github.com/ysicing/tiga/common"

	"github.com/cockroachdb/errors"
	"github.com/go-ping/ping"
	"github.com/jackpal/gateway"
	"github.com/miekg/dns"
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
	pinger.SetPrivileged(true)
	pinger.Count = 3
	pinger.Debug = true
	pinger.Interval = 200 * time.Millisecond
	pinger.Timeout = 3 * time.Second
	err = pinger.Run()
	if err != nil {
		log.GetInstance().Debugf("ping %s error: %s", host, err)
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

// CheckCaptivePortal checks if a captive portal is present
func CheckCaptivePortal(p ...string) error {
	checkURL := common.DefaultGenerate204URL
	if len(p) > 0 && strings.HasPrefix(p[0], "https://") {
		checkURL = p[0]
	}
	// nolint:gosec
	resp, err := http.Get(checkURL)
	if err != nil || resp.StatusCode != 204 {
		return err
	}
	return nil
}

// GetCloudflareEdgeTrace returns the edge POP
func GetCloudflareEdgeTrace() (string, error) {
	resp, err := http.Get("https://www.cloudflare.com/cdn-cgi/trace")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	body := string(bodyBytes)
	re := regexp.MustCompile(`colo=(.*?)\n`)
	match := re.FindStringSubmatch(body)
	if len(match) < 2 {
		return "", fmt.Errorf("could not determine edge pop")
	}
	return match[1], nil
}

// GetSkkMoeTrace returns the edge POP
func GetSkkMoeTrace() (string, error) {
	resp, err := http.Get("https://ip.skk.moe/cdn-cgi/trace")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	body := string(bodyBytes)
	re := regexp.MustCompile(`colo=(.*?)\n`)
	match := re.FindStringSubmatch(body)
	if len(match) < 2 {
		return "", fmt.Errorf("could not determine edge pop")
	}
	return match[1], nil
}
