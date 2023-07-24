// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package debug

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/pkg/factory"
)

func ClashCommand(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clash",
		Short: "clash",
	}
	cmd.AddCommand(clashSubLink(f))
	return cmd
}

func clashSubLink(f factory.Factory) *cobra.Command {
	var sublink []string
	var outfile string
	cmd := &cobra.Command{
		Use:   "sub",
		Short: "get proxies from subscribe link",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(sublink) == 0 {
				return errors.New("sublink is empty")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	cmd.Flags().StringArrayVarP(&sublink, "sub", "s", nil, "subscribe link")
	cmd.Flags().StringVarP(&outfile, "outfile", "o", "clash.yaml", "output file")
	return cmd
}
