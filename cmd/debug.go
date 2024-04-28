// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/cmd/debug"
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
	debugCmd.AddCommand(debug.TcpingCommand(f))
	debugCmd.AddCommand(debug.ChinaRouteCommand(f))
	return debugCmd
}
