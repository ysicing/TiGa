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
