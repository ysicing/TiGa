// Copyright (c) 2024 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package cfd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/ergoapi/util/output"
	"github.com/gosuri/uitable"

	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/common"
	"github.com/ysicing/tiga/internal/pkg/cft"
	"github.com/ysicing/tiga/internal/util"
	"github.com/ysicing/tiga/pkg/factory"
)

var cfkey, cfmail, cftoken string

func precheckCfApi(_ *cobra.Command, _ []string) error {
	cfkey = os.Getenv("CLOUDFLARE_API_KEY")
	if len(cfkey) == 0 {
		cfkey = os.Getenv("CF_Key")
	}
	cfmail = os.Getenv("CLOUDFLARE_API_EMAIL")
	if len(cfmail) == 0 {
		cfmail = os.Getenv("CF_Email")
	}
	cftoken = os.Getenv("CLOUDFLARE_API_TOKEN")
	if len(cftoken) == 0 {
		cftoken = os.Getenv("CF_Token")
	}
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
			authtype := "email"
			token := fmt.Sprintf("%s:%s", cfmail, cfkey)
			if len(cftoken) > 0 {
				authtype = "token"
				token = cftoken
			}
			client, err := cft.NewClient(authtype, token)
			if err != nil {
				return fmt.Errorf("create cloudflare client error: %w", err)
			}
			cfd, err := client.ListTunnels()
			if err != nil {
				return fmt.Errorf("list cloudflare tunnels error: %w", err)
			}
			switch strings.ToLower(common.ListOutput) {
			case "json":
				return output.EncodeJSON(os.Stdout, cfd)
			case "yaml":
				return output.EncodeYAML(os.Stdout, cfd)
			default:
				sort.Slice(cfd, func(i, j int) bool {
					// 首先按照状态排序：healthy 排在前面
					if cfd[i].Status != cfd[j].Status {
						if cfd[i].Status == "healthy" {
							return true
						}
						if cfd[j].Status == "healthy" {
							return false
						}
					}

					// 如果状态相同，根据状态类型使用不同的时间排序逻辑
					if cfd[i].Status == "healthy" {
						// healthy 状态：最近活跃的排在前面
						if cfd[i].ConnsActiveAt == nil {
							return false
						}
						if cfd[j].ConnsActiveAt == nil {
							return true
						}
						return cfd[i].ConnsActiveAt.After(*cfd[j].ConnsActiveAt)
					} else {
						// down 状态：最早失活的排在后面
						if cfd[i].ConnInactiveAt == nil {
							return false
						}
						if cfd[j].ConnInactiveAt == nil {
							return true
						}
						return cfd[i].ConnInactiveAt.After(*cfd[j].ConnInactiveAt)
					}
				})
				table := uitable.New()
				table.AddRow("ID", "Name", "TunnelType", "Status", "Connections", "CreatedAt", "ConnsActiveAt", "ConnInactiveAt")
				for _, p := range cfd {
					table.AddRow(p.ID, p.Name, p.TunnelType, p.Status, len(p.Connections), p.CreatedAt.Format("2006-01-02"), util.PtrFormatTime(p.ConnsActiveAt, "2006-01-02 15:04:05"), util.PtrFormatTime(p.ConnInactiveAt, "2006-01-02 15:04:05"))
				}
				return output.EncodeTable(os.Stdout, table)
			}
		},
	}
	cmd.Flags().StringVarP(&common.ListOutput, "output", "o", "",
		"prints the output in the specified format. Allowed values: table, json, yaml (default table)")
	return cmd
}
