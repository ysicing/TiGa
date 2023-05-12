// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/ergoapi/util/excmd"
	mcobra "github.com/muesli/mango-cobra"
	"github.com/muesli/roff"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/cmd/flags"
	"github.com/ysicing/tiga/common"
	"github.com/ysicing/tiga/pkg/factory"
	"github.com/ysicing/tiga/pkg/log"
)

var (
	globalFlags *flags.GlobalFlags
)

func Execute() {
	// create a new factory
	f := factory.DefaultFactory()
	// build the root command
	rootCmd := BuildRoot(f)
	// before hook
	// execute command
	err := rootCmd.Execute()
	// after hook
	if err != nil {
		if globalFlags.Debug {
			f.GetLog().Fatalf("%v", err)
		} else {
			f.GetLog().Fatal(err)
		}
		if !strings.Contains(err.Error(), "unknown command") {
			f.GetLog().Info("----------------------------")
			bugMsg := "found bug: submit the error message to Github \n\t Github: https://github.com/ysicing/tiga"
			f.GetLog().Info(bugMsg)
		}
	}
}

// BuildRoot creates a new root command from the
func BuildRoot(f factory.Factory) *cobra.Command {
	// build the root cmd
	rootCmd := NewRootCmd(f)
	persistentFlags := rootCmd.PersistentFlags()
	globalFlags = flags.SetGlobalFlags(persistentFlags)
	rootCmd.AddCommand(newCmdVersion(f))
	rootCmd.AddCommand(newCmdUpgrade(f))
	rootCmd.AddCommand(NewCmdPlugin())
	rootCmd.AddCommand(newCmdApp(f))
	rootCmd.AddCommand(newCmdDebug(f))

	rootCmd.AddCommand(newManCmd())

	args := os.Args
	if len(args) > 1 {
		pluginHandler := excmd.NewDefaultPluginHandler(common.GetDefaultBinDir(), common.ValidPrefixes)
		cmdPathPieces := args[1:]
		if _, _, err := rootCmd.Find(cmdPathPieces); err != nil {
			var cmdName string // first "non-flag" arguments
			for _, arg := range cmdPathPieces {
				if !strings.HasPrefix(arg, "-") {
					cmdName = arg
					break
				}
			}
			switch cmdName {
			case "help", cobra.ShellCompRequestCmd, cobra.ShellCompNoDescRequestCmd:
				// Don't search for a plugin
			default:
				if err := excmd.HandlePluginCommand(pluginHandler, cmdPathPieces); err != nil {
					fmt.Fprintf(os.Stdout, "Error: %v\n", err)
					os.Exit(1)
				}
			}
		}
	}
	return rootCmd
}

// NewRootCmd returns a new root command
func NewRootCmd(f factory.Factory) *cobra.Command {
	return &cobra.Command{
		Use:           "tiga",
		SilenceUsage:  true,
		SilenceErrors: true,
		Short:         "Tiga is a cli tool for senior restart engineer",
		PersistentPreRunE: func(cobraCmd *cobra.Command, args []string) error {
			if cobraCmd.Annotations != nil {
				return nil
			}
			qlog := f.GetLog()
			if globalFlags.Silent {
				qlog.SetLevel(logrus.FatalLevel)
			} else if globalFlags.Debug {
				qlog.SetLevel(logrus.DebugLevel)
			}

			log.StartFileLogging()
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			common.ShowLogo()
			cmd.Help()
		},
	}
}

func newManCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "man",
		Short:                 "Generates tiga's command line manpages",
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
		Hidden:                true,
		Args:                  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			manPage, err := mcobra.NewManPage(1, cmd.Root())
			if err != nil {
				return err
			}
			_, err = fmt.Fprint(os.Stdout, manPage.Build(roff.NewDocument()))
			return err
		},
	}
	return cmd
}
