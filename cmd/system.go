// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package cmd

import (
	"github.com/ergoapi/util/zos"
	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/cmd/system"
	"github.com/ysicing/tiga/pkg/factory"
)

func newCmdSystem(f factory.Factory) *cobra.Command {
	systemCmd := &cobra.Command{
		Use:   "system",
		Short: "system ops tools",
		Long:  "system ops tools",
	}
	if zos.IsLinux() {
		systemCmd.AddCommand(system.DebianCommand(f))
	}
	return systemCmd
}
