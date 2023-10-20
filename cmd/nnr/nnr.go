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

	"github.com/ergoapi/util/color"
	"github.com/ergoapi/util/output"
	"github.com/gosuri/uitable"
	"github.com/manifoldco/promptui"
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
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(token) == 0 {
				return fmt.Errorf("token is empty")
			}
			return nil
		},
	}
	cmd.AddCommand(listServers(f))
	cmd.AddCommand(listRules(f))
	cmd.AddCommand(listRemoteNode(f))
	cmd.AddCommand(addRule(f))
	// cmd.AddCommand(delRule(f))
	// cmd.AddCommand(updateRule(f))
	cmd.PersistentFlags().StringVarP(&token, "token", "t", os.Getenv("NNR_TOKEN"), "token")
	return cmd
}

func listServers(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "servers",
		Short: "list servers",
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

var all bool

func listRules(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rules",
		Short: "list rules",
		RunE: func(cmd *cobra.Command, args []string) error {
			api := nnr.New(token)
			servers := api.ServersMap()
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
			table.AddRow("规则", "节点", "转发地址", "远程地址", "类型", "流量")
			for _, index := range s {
				rulename := index.Name
				sname := index.Sid
				value, exist := servers[index.Sid]
				if exist {
					if all {
						sname = fmt.Sprintf("%s(%s)", value.Name, color.SGreen(index.Sid))
					} else {
						sname = value.Name
					}
				}
				if all {
					rulename = fmt.Sprintf("%s(%s)", index.Name, color.SGreen(index.Rid))
				}
				table.AddRow(rulename, sname, fmt.Sprintf("%v:%v", index.Host, index.Port), fmt.Sprintf("%v:%v", index.Remote, index.RPort), index.Type, util.Traffic(index.Traffic, 1000.0))
			}
			return output.EncodeTable(os.Stdout, table)
		},
	}
	cmd.Flags().BoolVar(&all, "all", false, "all output")
	return cmd
}

func listRemoteNode(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nodes",
		Aliases: []string{"remotes"},
		Short:   "list remote nodes",
		RunE: func(cmd *cobra.Command, args []string) error {
			api := nnr.New(token)
			s, err := api.SortRemoteNode()
			if err != nil {
				return err
			}
			if len(s) == 0 {
				f.GetLog().Infof("no nodes found")
				return nil
			}
			f.GetLog().Infof("found %d node", len(s))
			sort.Slice(s, func(i, j int) bool {
				return s[i].Traffic >= s[j].Traffic
			})
			table := uitable.New()
			table.AddRow("远程地址", "流量")
			for _, index := range s {
				table.AddRow(index.Remote, util.Traffic(index.Traffic, 1000.0))
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
			searcher := func(input string, index int) bool {
				pepper := s[index]
				name := strings.Replace(strings.ToLower(pepper.Name), " ", "", -1)
				input = strings.Replace(strings.ToLower(input), " ", "", -1)
				return strings.Contains(name, input)
			}
			selectApp := promptui.Select{
				Label: "select app",
				Items: s,
				Templates: &promptui.SelectTemplates{
					Label:    "{{ . }}?",
					Active:   "\U0001F449 {{ .Name | cyan }} ({{ .Mf }}倍率)",
					Inactive: "  {{ .Name | cyan }}",
					Selected: "\U0001F389 {{ .Name | green | cyan }} ({{ .Mf }}倍率)",
					Details: `
{{ "Detail:" | faint }} {{ .Name | green }} {{ .Detail }}
`,
				},
				Size:     10,
				Searcher: searcher,
			}
			it, _, _ := selectApp.Run()
			f.GetLog().Debugf("select server: %s", s[it].Name)
			prompt := promptui.Prompt{
				Label: "remote",
				Validate: func(input string) error {
					s := strings.Split(input, ":")
					if len(s) != 2 {
						return fmt.Errorf("invalid remote")
					}
					return nil
				},
			}
			result, err := prompt.Run()
			if err != nil {
				return err
			}
			t := "tcp"
			if strings.Contains(s[it].Detail, "udp") || strings.Contains(s[it].Detail, "UDP") || strings.Contains(s[it].Name, "IEPL") {
				t = "tcp+udp"
			}
			r, err := api.AddRule(s[it].Sid, result, t)
			if err != nil {
				return err
			}
			f.GetLog().Infof("add rule success, rule: %s %s:%v -> %s:%v", r.Rid, r.Host, r.Port, r.Remote, r.RPort)
			return nil
		},
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
