// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package app

import (
	"github.com/ergoapi/util/zos"
	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/pkg/factory"
)

const (
	appInstall = `
  # Install app
  tiga app install --help
`
)

func InstallApp(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "install",
		Aliases: []string{"i"},
		Short:   "Install app",
		Example: appInstall,
	}
	if zos.IsLinux() {
		cmd.AddCommand(installGo(f))
	}
	return cmd
}
