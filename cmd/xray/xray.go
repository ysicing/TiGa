// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package xray

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/internal/pkg/xray"
	"github.com/ysicing/tiga/pkg/factory"
)

var api string

func NewCmdXray(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "xray",
		Short:   "xray",
		Version: "0.2.0",
	}
	cmd.AddCommand(trafficXray(f))
	cmd.PersistentFlags().StringVar(&api, "api", "127.0.0.1:10086", "api")
	return cmd
}

func trafficXray(f factory.Factory) *cobra.Command {
	log := f.GetLog()
	xrayCtl := new(xray.XrayController)
	t := &cobra.Command{
		Use:     "traffic",
		Short:   "traffic",
		Version: "0.2.0",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := xrayCtl.Init(api); err != nil {
				return fmt.Errorf("init xray controller failed: %v", err)
			}
			defer xrayCtl.CmdConn.Close()
			log.Info("xray exporter start")
			return xrayCtl.QueryTraffic(false)
		},
	}
	return t
}
