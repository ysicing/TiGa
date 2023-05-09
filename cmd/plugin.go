// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/cmd/plugin"
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
	cmd.AddCommand(NewCmdPluginList())
	cmd.AddCommand(NewPluginInstall())
	cmd.AddCommand(NewPluginSearch())
	return cmd
}

func NewCmdPluginList() *cobra.Command {
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

func NewPluginInstall() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "install [flags]",
		Aliases: []string{"i"},
		Short:   "install plugin",
		Example: pluginInstallExample,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}

func NewPluginSearch() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "search [flags]",
		Short:   "search plugin from repository",
		Example: pluginSearchExample,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}
