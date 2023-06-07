// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package debug

import (
	"os"
	"strconv"

	"github.com/spf13/cobra"
	tops "github.com/ysicing/tiga/internal/pkg/gops"
	"github.com/ysicing/tiga/pkg/factory"
	"k8s.io/kubectl/pkg/util/templates"
)

var (
	processExample = templates.Examples(`
  # simple process info
  tiga debug gops process <pid>
  # simple process info with period time(seconds)
  tiga debug gops process <pid> <time>
  `)
)

func GOpsCommand(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gops",
		Short: "gops is a tool to list and diagnose Go processes.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(os.Args) > 3 {
				_, err := strconv.Atoi(os.Args[3])
				if err == nil {
					f.GetLog().Infof("fetch pid %s process info", os.Args[3])
					tops.ProcessInfo(os.Args[3:]) // shift off the command name
					return
				}
			}
			f.GetLog().Info("fetch all process info")
			tops.Processes()
		},
	}
	cmd.AddCommand(treeCommand())
	cmd.AddCommand(processCommand())
	return cmd
}

// treeCommand displays a process tree.
func treeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "tree",
		Short: "Display parent-child tree for Go processes.",
		Run: func(cmd *cobra.Command, args []string) {
			tops.DisplayProcessTree()
		},
	}
}

// processCommand displays information about a Go process.
func processCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "process",
		Aliases: []string{"pid", "proc"},
		Short:   "Prints information about a Go process.",
		Example: processExample,
		Args:    cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			tops.ProcessInfo(args)
			return nil
		},
	}
}
