// Copyright (c) 2024 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package nas

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/internal/pkg/wol"
	"github.com/ysicing/tiga/pkg/factory"
)

func WolCmd(f factory.Factory) *cobra.Command {
	var mac string
	cmd := &cobra.Command{
		Use:   "wol",
		Short: "wol tools",
		RunE: func(_ *cobra.Command, _ []string) error {
			if mac == "" {
				return fmt.Errorf("mac address is required")
			}
			if err := wol.Wake(mac); err != nil {
				f.GetLog().Warnf("wol %s failed: %v", mac, err)
				return nil
			}
			f.GetLog().Donef("wol %s success", mac)
			return nil
		},
	}
	cmd.Flags().StringVarP(&mac, "mac", "m", "", "mac address")
	return cmd
}
