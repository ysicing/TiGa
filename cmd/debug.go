// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"

	"github.com/ergoapi/util/zos"
	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/cmd/clash"
	"github.com/ysicing/tiga/cmd/debug"
	"github.com/ysicing/tiga/cmd/xray"
	"github.com/ysicing/tiga/pkg/factory"
)

func newCmdDebug(f factory.Factory) *cobra.Command {
	debugCmd := &cobra.Command{
		Use:   "debug",
		Short: "debug, not a stable interface, contains misc debug facilities",
		Long:  fmt.Sprintf("\"%s debug\" contains misc debug facilities; it is not a stable interface.", os.Args[0]),
	}
	debugCmd.AddCommand(debug.HostInfoCommand(f))
	debugCmd.AddCommand(debug.DownloadCommand(f))
	debugCmd.AddCommand(debug.CleanCacheCommand(f))
	debugCmd.AddCommand(debug.GOpsCommand(f))
	debugCmd.AddCommand(debug.NetCheckCommand(f))
	debugCmd.AddCommand(debug.IPMMDBCommand(f))
	if zos.IsLinux() {
		debugCmd.AddCommand(debug.ChinaRouteCommand(f))
	}
	// Deprecated commands
	xray := xray.NewCmdXray(f)
	xray.Deprecated = "please use xray"
	debugCmd.AddCommand(xray)
	clash := clash.NewCmdClash(f)
	clash.Deprecated = "please use clash"
	debugCmd.AddCommand(clash)
	return debugCmd
}
