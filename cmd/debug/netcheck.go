// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package debug

import (
	"sync"

	"github.com/ysicing/tiga/pkg/util/ts"

	"github.com/ergoapi/util/color"

	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/common"
	"github.com/ysicing/tiga/pkg/factory"
	"github.com/ysicing/tiga/pkg/util/netutil"
)

func NetCheckCommand(f factory.Factory) *cobra.Command {
	logpkg := f.GetLog()
	cmd := &cobra.Command{
		Use:   "netcheck",
		Short: "netcheck",
		Run: func(cmd *cobra.Command, args []string) {
			if gw, err := netutil.CheckDefaultRoute(); err == nil {
				logpkg.Donef("default route %s reachable via icmp", color.SGreen(gw.String()))
			} else {
				logpkg.Warnf("default route %s unreachable via icmp", color.SRed(gw.String()))
			}
			if ns, err := netutil.GetDefaultNameserver(); err == nil {
				a := "reachable"
				if !netutil.CheckReachabilityWithICMP(ns) {
					a = "unreachable"
				}
				if err := netutil.CheckNameserverAvailability(ns + ":53"); err != nil {
					logpkg.Warnf("nameserver: %s (ICMP %s, DNS unreachable: %s)", color.SRed(ns), a, err)
				} else {
					logpkg.Donef("nameserver: %s (ICMP %s, DNS reachable)", color.SGreen(ns), a)
				}
			} else {
				logpkg.Warnf("failed to reading default nameserver from system: %s", err)
			}
			var wg sync.WaitGroup
			wg.Add(6)
			go func() {
				defer wg.Done()
				if err := netutil.CheckCaptivePortal(); err == nil {
					logpkg.Donef("captive portal %s detected success", color.SGreen(common.DefaultGenerate204URL))
				} else {
					logpkg.Warnf("captive portal %s detected failed: %s", color.SRed(common.DefaultGenerate204URL), err)
				}
			}()
			go func() {
				defer wg.Done()
				if err := netutil.CheckCaptivePortal(common.MiuiGenerate204URL); err == nil {
					logpkg.Donef("captive portal %s detected success", color.SGreen(common.MiuiGenerate204URL))
				} else {
					logpkg.Warnf("captive portal %s detected failed: %s", color.SRed(common.MiuiGenerate204URL), err)
				}
			}()
			go func() {
				defer wg.Done()
				if err := netutil.CheckNameserverAvailability("8.8.8.8:53"); err != nil {
					logpkg.Warnf("remote dns %s is unavailable: %s", color.SRed("8.8.8.8"), err)
				} else {
					logpkg.Donef("remote dns %s is available", color.SGreen("8.8.8.8"))
				}
			}()
			go func() {
				defer wg.Done()
				if err := netutil.CheckNameserverAvailability("114.114.114.114:53"); err != nil {
					logpkg.Warnf("remote dns %s is unavailable: %s", color.SRed("114.114.114.114"), err)
				} else {
					logpkg.Donef("remote dns %s is available", color.SGreen("114.114.114.114"))
				}
			}()
			go func() {
				defer wg.Done()
				if loc, err := netutil.GetCloudflareEdgeTrace(); err == nil {
					logpkg.Donef("match Cloudflare CDN: %s", color.SGreen(loc))
				} else {
					logpkg.Warnf("miss Cloudflare CDN failed")
				}
			}()
			go func() {
				defer wg.Done()
				ts.GetTailscaleStatus()
			}()
			wg.Wait()
		},
	}
	return cmd
}
