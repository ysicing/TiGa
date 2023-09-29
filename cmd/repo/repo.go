// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package repo

import (
	"os"
	"strings"

	"github.com/ergoapi/util/color"

	"github.com/cockroachdb/errors"

	"github.com/ergoapi/util/output"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/common"
	"github.com/ysicing/tiga/internal/pkg/repo"
	"github.com/ysicing/tiga/pkg/factory"
)

func IndexListCmd(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List configured indexes",
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			indexes, err := repo.LoadIndex()
			if err != nil {
				return err
			}
			switch strings.ToLower(common.ListOutput) {
			case "json":
				return output.EncodeJSON(os.Stdout, indexes)
			case "yaml":
				return output.EncodeYAML(os.Stdout, indexes)
			default:
				f.GetLog().Infof("last generated: %v", indexes.Generated)
				table := uitable.New()
				table.AddRow("INDEX", "URL", "DESC")
				for _, index := range indexes.Index {
					table.AddRow(index.Name, index.Url, index.Desc)
				}
				return output.EncodeTable(os.Stdout, table)
			}
		},
	}
	cmd.Flags().StringVarP(&common.ListOutput, "output", "o", "",
		"prints the output in the specified format. Allowed values: table, json, yaml (default table)")
	return cmd
}

func IndexAddCmd(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add",
		Short:   "Add a new index",
		Args:    cobra.ExactArgs(2),
		Example: `tiga repo add <name> <url>`,
		RunE: func(_ *cobra.Command, args []string) error {
			name := args[0]
			exist, err := repo.IsValidIndexName(name)
			if err != nil {
				return err
			}
			if exist && name == "default" {
				return errors.New("built-in not allow modify")
			}
			url := args[1]
			if strings.HasPrefix(url, "http") || strings.HasPrefix(url, "file://") {
				return errors.New("url must start with http or file://")
			}
			err = repo.AddIndex(name, url)
			if err == nil {
				f.GetLog().Donef("rebuild index %s success", color.SGreen(name))
				return nil
			}
			return err
		},
	}
	return cmd
}

func IndexDeleteCmd(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"remove"},
		Short:   "Delete an existing index",
		Example: `tiga repo delete <name>`,
		Args:    cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			name := args[0]
			exist, err := repo.IsValidIndexName(name)
			if err != nil {
				return err
			}
			if !exist {
				return errors.New("index not found")
			}
			if exist && name == "default" {
				return errors.New("built-in not allow remove")
			}
			err = repo.DeleteIndex(name)
			if err == nil {
				f.GetLog().Donef("remove index %s success", color.SBlue(name))
				return nil
			}
			return err
		},
	}
	return cmd
}

func IndexUpdateCmd(f factory.Factory) *cobra.Command {
	var force bool
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update index",
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			return repo.UpdateIndexs(force)
		},
	}
	cmd.Flags().BoolVarP(&force, "force", "f", false, "force update")
	return cmd
}
