// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
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
	"github.com/ysicing/tiga/common"
	"github.com/ysicing/tiga/pkg/exec"
	"github.com/ysicing/tiga/pkg/factory"
)

func swapCommand(f factory.Factory) *cobra.Command {
	var swapSize int64
	debianCmd := &cobra.Command{
		Use:     "swap",
		Short:   "swap op",
		Long:    "add swap to linux",
		Version: "0.2.2",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			swap, err := mem.SwapMemory()
			if err != nil {
				return fmt.Errorf("check swap info failed: %v", err)
			}
			if swap.Total > 0 {
				f.GetLog().Info("swap already exists")
				os.Exit(0)
			}
			if swapSize < 0 {
				swapSize = 1
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			f.GetLog().Infof("not found swap, will create swap %dG", swapSize)
			return exec.CommandRun("bash", "-c", common.GetCustomScriptFile("hack/manifests/system/swap.sh"), fmt.Sprintf("%d", swapSize))
		},
	}
	debianCmd.Flags().Int64VarP(&swapSize, "size", "s", 1, "swap size 1GB")
	return debianCmd
}
