// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package app

import (
	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/pkg/factory"
)

func installGost(f factory.Factory) *cobra.Command {
	logpkg := f.GetLog()
	return &cobra.Command{
		Use:   "gost",
		Short: "gost",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			logpkg.Infof("precheck gost")
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
