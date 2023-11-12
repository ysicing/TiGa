// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package cmd

import (
	"os"
	"strings"

	"github.com/ergoapi/util/output"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/cmd/plugin"
	"github.com/ysicing/tiga/common"
	"github.com/ysicing/tiga/pkg/log"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	pluginListExample = templates.Examples(`
          # List all available plugins
          tiga plugin list`)
	pluginInstallExample = templates.Examples(`
          # install a plugin from repository
          tiga plugin install [options] [flags]`)
	pluginSearchExample = templates.Examples(`
          # search plugin from repository
          tiga plugin search [options] [flags]`)
)

func NewCmdPlugin() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "plugin [flags]",
		DisableFlagsInUseLine: true,
		Short:                 "provides utilities for interacting with plugins",
	}
	cmd.AddCommand(pluginListCmd())
	cmd.AddCommand(pluginInstallCmd())
	cmd.AddCommand(pluginSearchCmd())
	return cmd
}

func pluginListCmd() *cobra.Command {
	o := plugin.ListOptions{}
	cmd := &cobra.Command{
		Use:                   "list [flags]",
		Aliases:               []string{"ls"},
		DisableFlagsInUseLine: true,
		Short:                 "list all visible plugins executable by tiga on your PATH",
		Example:               pluginListExample,
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(o.Complete(cmd))
			cmdutil.CheckErr(o.Run())
		},
	}
	cmd.Flags().BoolVar(&o.NameOnly, "name-only", false, "print only the plugin names")
	return cmd
}

func pluginInstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "install [flags]",
		Aliases: []string{"i"},
		Short:   "install plugin from repository",
		Example: pluginInstallExample,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}

func pluginSearchCmd() *cobra.Command {
	o := plugin.SearchOptions{
		Log: log.GetInstance(),
	}
	cmd := &cobra.Command{
		Use:     "search [flags]",
		Aliases: []string{"ls-remote"},
		Short:   "search plugin from repository",
		Example: pluginSearchExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			name := ""
			if len(args) > 0 {
				name = args[0]
			}
			plugins, err := o.Search(name)
			if err != nil {
				o.Log.Warnf("search plugin error: %v", err)
				return nil
			}
			o.Log.Infof("search plugin result count: %v", len(plugins))
			// spew.Dump(plugins)
			switch strings.ToLower(common.ListOutput) {
			case "json":
				return output.EncodeJSON(os.Stdout, plugins)
			case "yaml":
				return output.EncodeYAML(os.Stdout, plugins)
			default:
				table := uitable.New()
				table.AddRow("index", "name", "version", "url")
				for _, p := range plugins {
					for _, v := range p.Plugins {
						if len(v.Url) == 0 {
							if p.Index != "ysicing" {
								continue
							}
							v.Url = common.GetCustomBinary(v.Name)
						}
						table.AddRow(p.Index, v.Name, v.Version, v.Url)
					}
				}
				return output.EncodeTable(os.Stdout, table)
			}
		},
	}
	cmd.Flags().StringVarP(&common.ListOutput, "output", "o", "", "prints the output in the specified format. Allowed values: table, json, yaml (default table)")
	return cmd
}
