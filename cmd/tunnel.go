package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/cmd/cfd"
	"github.com/ysicing/tiga/pkg/factory"
)

func newCmdCfdTunnel(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tunnel",
		Short: "manage cfd tunnel",
		Args:  cobra.NoArgs,
	}
	cmd.AddCommand(cfd.TunnelListCmd(f))
	return cmd
}
