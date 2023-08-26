// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package clash

import (
	"os"
	"strings"

	C "github.com/Dreamacro/clash/constant"
	"github.com/cockroachdb/errors"
	"github.com/ergoapi/util/color"
	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/common"
	"github.com/ysicing/tiga/internal/pkg/clash"
	"github.com/ysicing/tiga/pkg/factory"
	"gopkg.in/yaml.v2"
)

var config, filterRegexConfig, name string

func NewCmdClash(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "clash",
		Short:   "clash",
		Version: "0.3.0",
	}
	cmd.AddCommand(filterProxy(f))
	cmd.PersistentFlags().StringVar(&config, "config", "/etc/clash/config.yaml", "clash config")
	return cmd
}

func filterProxy(f factory.Factory) *cobra.Command {
	log := f.GetLog()
	t := &cobra.Command{
		Use:     "filter",
		Short:   "filter proxy",
		Version: "0.3.0",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if strings.HasPrefix(config, "http") {
				return errors.New("unsupport remote config")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var allProxies = make(map[string]clash.CProxy)
			ps, err := clash.LoadProxies(config)
			if err != nil {
				return err
			}
			log.Infof("load %s clash config success, detected %d proxies", color.SGreen(config), len(ps))
			for k, p := range ps {
				if _, ok := allProxies[k]; !ok {
					allProxies[k] = p
				}
			}
			filteredProxies := clash.FilterProxies(filterRegexConfig, allProxies)
			log.Infof("use '%s' filtered %d proxies", color.SGreen(filterRegexConfig), len(filteredProxies))
			results := make([]clash.Result, 0, len(filteredProxies))
			for _, name := range filteredProxies {
				proxy := allProxies[name]
				switch proxy.Type() {
				case C.Shadowsocks, C.ShadowsocksR, C.Snell, C.Socks5, C.Http, C.Vmess, C.Trojan:
					r := clash.Result{Name: name}
					results = append(results, r)
				}
			}
			return writeNodeConfigurationToYAML(common.SavePath, results, allProxies)
		},
	}
	t.Flags().StringVarP(&name, "name", "n", "", "rename proxy")
	t.Flags().StringVarP(&common.SavePath, "save", "s", "", "save path")
	t.Flags().StringVarP(&filterRegexConfig, "filter", "f", ".*", "filter proxies by name, use regexp")
	return t
}

func writeNodeConfigurationToYAML(filePath string, results []clash.Result, proxies map[string]clash.CProxy) error {
	fp, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer fp.Close()

	var sortedProxies []any
	for _, result := range results {
		if v, ok := proxies[result.Name]; ok {
			sortedProxies = append(sortedProxies, v.SecretConfig)
		}
	}
	bytes, err := yaml.Marshal(sortedProxies)
	if err != nil {
		return err
	}

	_, err = fp.Write(bytes)
	return err
}
