// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package debug

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/internal/pkg/hostinfo"
	"github.com/ysicing/tiga/pkg/factory"
)

func HostInfoCommand(f factory.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "hostinfo",
		Short: "print hostinfo",
		RunE: func(cmd *cobra.Command, args []string) error {
			hi := hostinfo.New()
			j, _ := json.MarshalIndent(hi, "", "  ")
			os.Stdout.Write(j)
			return nil
		},
	}
}
