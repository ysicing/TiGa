package cfd

import (
	"os"

	"github.com/cockroachdb/errors"

	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/common"
	"github.com/ysicing/tiga/pkg/factory"
)

var cfkey, cfmail, cftoken string

func precheckCfApi(_ *cobra.Command, _ []string) error {
	cfkey = os.Getenv("CLOUDFLARE_API_KEY")
	cfmail = os.Getenv("CLOUDFLARE_API_EMAIL")
	cftoken = os.Getenv("CLOUDFLARE_API_TOKEN")
	if cftoken != "" {
		return nil
	}
	if cfkey == "" || cfmail == "" {
		return errors.New("CLOUDFLARE_API_KEY or CLOUDFLARE_API_EMAIL not found")
	}
	return nil
}

func TunnelListCmd(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List tunnel",
		Args:    cobra.NoArgs,
		PreRunE: precheckCfApi,
		RunE: func(_ *cobra.Command, _ []string) error {
			return nil
		},
	}
	cmd.Flags().StringVarP(&common.ListOutput, "output", "o", "",
		"prints the output in the specified format. Allowed values: table, json, yaml (default table)")
	return cmd
}
