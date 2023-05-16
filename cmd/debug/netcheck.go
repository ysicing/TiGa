// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package debug

import (
	"github.com/spf13/cobra"
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
				logpkg.Donef("default route %s reachable via icmp", gw.String())
			} else {
				logpkg.Warnf("default route %s unreachable via icmp", gw.String())
			}
			if ns, err := netutil.GetDefaultNameserver(); err == nil {
				a := ""

				if !netutil.CheckReachabilityWithICMP(ns) {
					a = "un"
				}
				if err := netutil.CheckNameserverAvailability(ns + ":53"); err != nil {
					logpkg.Warnf("Default nameserver: %s (ICMP %s reachable, DNS unreachable: %s)", ns, a, err)
				} else {
					logpkg.Donef("Default nameserver: %s (ICMP %s reachable, DNS reachable)", ns, a)
				}
			} else {
				logpkg.Warnf("failed to reading default nameserver from system: %s", err)
			}
		},
	}
	return cmd
}
