// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package debug

import (
	"github.com/ergoapi/util/color"
	"github.com/ergoapi/util/exnet"
	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/pkg/factory"
	"github.com/ysicing/tiga/pkg/util/ipdb"
)

func IPMMDBCommand(f factory.Factory) *cobra.Command {
	logpkg := f.GetLog()
	var ip string
	cmd := &cobra.Command{
		Use:   "mmdb",
		Short: "mmdb",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return ipdb.InitMMDB()
		},
		Run: func(cmd *cobra.Command, args []string) {
			if len(ip) == 0 {
				ip, _ = exnet.OutboundIPv2()
				if len(ip) == 0 {
					logpkg.Warnf("ip is empty")
					return
				}
			}
			if ipdb.MatchCN(ip) {
				logpkg.Infof("China ip %s", color.SGreen(ip))
				return
			}
			logpkg.Infof("Global ip %s", color.SBlue(ip))
		},
	}
	cmd.Flags().StringVar(&ip, "ip", "", "ip")
	return cmd
}
