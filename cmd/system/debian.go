// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package system

import (
	"fmt"

	"github.com/ergoapi/util/zos"
	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/pkg/factory"
)

func DebianCommand(f factory.Factory) *cobra.Command {
	debianCmd := &cobra.Command{
		Use:   "debian",
		Short: "debian ops tools",
		Long:  "debian ops tools",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if !zos.IsDebian() {
				return fmt.Errorf("now only support debian linux")
			}
			return nil
		},
	}
	debianCmd.AddCommand(swapCommand(f))
	return debianCmd
}
