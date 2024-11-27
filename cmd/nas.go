// Copyright (c) 2024 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/cmd/nas"
	"github.com/ysicing/tiga/pkg/factory"
)

func newCmdNas(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nas",
		Short: "nas tools",
		Args:  cobra.NoArgs,
	}
	cmd.AddCommand(nas.WolCmd(f))
	return cmd
}
