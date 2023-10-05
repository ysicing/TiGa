// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package nnr

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/ergoapi/util/output"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/internal/pkg/nnr"
	"github.com/ysicing/tiga/internal/util"
	"github.com/ysicing/tiga/pkg/factory"
)

var token string

func NewCmdNNR(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nnr",
		Short:   "nnr tools",
		Version: "2023.10.0519",
	}
	cmd.AddCommand(listServers(f))
	cmd.AddCommand(listRules(f))
	cmd.AddCommand(addRule(f))
	cmd.AddCommand(delRule(f))
	cmd.AddCommand(updateRule(f))
	cmd.PersistentFlags().StringVarP(&token, "token", "t", os.Getenv("NNR_TOKEN"), "token")
	return cmd
}

func listServers(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nodes",
		Aliases: []string{"servers"},
		Short:   "list servers",
		RunE: func(cmd *cobra.Command, args []string) error {
			api := nnr.New(token)
			s, err := api.ListServers()
			if err != nil {
				return err
			}
			if len(s) == 0 {
				f.GetLog().Infof("no servers found")
				return nil
			}
			f.GetLog().Infof("found %d servers", len(s))
			sort.Slice(s, func(i, j int) bool {
				return s[i].Mf >= s[j].Mf
			})
			table := uitable.New()
			table.AddRow("标识", "节点", "IP", "倍率", "支持TCP+UDP", "描述")
			for _, index := range s {
				table.AddRow(index.Sid, index.Name, index.Host, index.Mf, len(index.Types) == 3, strings.ReplaceAll(index.Detail, "\n", " "))
			}
			return output.EncodeTable(os.Stdout, table)
		},
	}
	return cmd
}

func listRules(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rules",
		Short: "list rules",
		RunE: func(cmd *cobra.Command, args []string) error {
			api := nnr.New(token)
			s, err := api.ListRules()
			if err != nil {
				return err
			}
			if len(s) == 0 {
				f.GetLog().Infof("no rules found")
				return nil
			}
			f.GetLog().Infof("found %d rules", len(s))
			sort.Slice(s, func(i, j int) bool {
				return s[i].Traffic >= s[j].Traffic
			})
			table := uitable.New()
			table.AddRow("规则标识", "节点标识", "转发地址", "远程地址", "类型", "流量")
			for _, index := range s {
				table.AddRow(index.Rid, index.Sid, fmt.Sprintf("%v:%v", index.Host, index.Port), fmt.Sprintf("%v:%v", index.Remote, index.RPort), index.Type, util.Traffic(index.Traffic))
			}
			return output.EncodeTable(os.Stdout, table)
		},
	}
	return cmd
}

func addRule(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add rule",
	}
	return cmd
}

func delRule(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "del",
		Short: "del rule",
	}
	return cmd
}

func updateRule(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "update rule",
	}
	return cmd
}
