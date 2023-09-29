// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/cmd/repo"
	"github.com/ysicing/tiga/pkg/factory"
)

func newCmdRepo(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repo",
		Short: "manage custom plugin indexes",
		Args:  cobra.NoArgs,
	}
	cmd.AddCommand(repo.IndexListCmd(f))
	cmd.AddCommand(repo.IndexAddCmd(f))
	cmd.AddCommand(repo.IndexDeleteCmd(f))
	cmd.AddCommand(repo.IndexUpdateCmd(f))
	return cmd
}
