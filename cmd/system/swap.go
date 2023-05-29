// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package system

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/v3/mem"
	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/pkg/factory"
)

func swapCommand(f factory.Factory) *cobra.Command {
	debianCmd := &cobra.Command{
		Use:   "swap",
		Short: "swap op",
		Long:  "add swap to linux",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			swap, err := mem.SwapMemory()
			if err != nil {
				return fmt.Errorf("check swap info failed: %v", err)
			}
			if swap.Total > 0 {
				f.GetLog().Info("swap already exists")
				os.Exit(0)
			}
			f.GetLog().Info("not found swap, will create swap")
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return debianCmd
}
