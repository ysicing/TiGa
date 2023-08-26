// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package clash

import (
	"os"
	"regexp"
	"sort"

	"github.com/Dreamacro/clash/adapter"
	C "github.com/Dreamacro/clash/constant"
	"github.com/cockroachdb/errors"
	"gopkg.in/yaml.v2"
)

type CProxy struct {
	C.Proxy
	SecretConfig any
}

type Result struct {
	Name string
}

type RawConfig struct {
	// Providers map[string]map[string]any `yaml:"proxy-providers"`
	Proxies []map[string]any `yaml:"proxies"`
}

// LoadProxies load proxies from config
func LoadProxies(filepath string) (map[string]CProxy, error) {
	body, err := os.ReadFile(filepath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read config")
	}
	rawCfg := &RawConfig{
		Proxies: []map[string]any{},
	}
	if err := yaml.Unmarshal(body, rawCfg); err != nil {
		return nil, err
	}
	proxies := make(map[string]CProxy)
	proxiesConfig := rawCfg.Proxies
	// providersConfig := rawCfg.Providers
	for i, config := range proxiesConfig {
		proxy, err := adapter.ParseProxy(config)
		if err != nil {
			return nil, errors.Errorf("proxy %d: %w", i, err)
		}

		if _, exist := proxies[proxy.Name()]; exist {
			return nil, errors.Errorf("proxy %s is the duplicate name", proxy.Name())
		}
		proxies[proxy.Name()] = CProxy{Proxy: proxy, SecretConfig: config}
	}
	// for name, config := range providersConfig {
	// 	if name == provider.ReservedName {
	// 		return nil, errors.Errorf("can not defined a provider called `%s`", provider.ReservedName)
	// 	}
	// 	pd, err := provider.ParseProxyProvider(name, config)
	// 	if err != nil {
	// 		return nil, errors.Errorf("parse proxy provider %s error: %w", name, err)
	// 	}
	// 	if err := pd.Initial(); err != nil {
	// 		return nil, errors.Errorf("initial proxy provider %s error: %w", pd.Name(), err)
	// 	}
	// 	for _, proxy := range pd.Proxies() {
	// 		spew.Dump(proxy)
	// 		proxies[fmt.Sprintf("[%s] %s", name, proxy.Name())] = CProxy{Proxy: proxy, SecretConfig: config}
	// 	}
	// }
	return proxies, nil
}

func FilterProxies(filter string, proxies map[string]CProxy) []string {
	filterRegexp := regexp.MustCompile(filter)
	filteredProxies := make([]string, 0, len(proxies))
	for name := range proxies {
		if filterRegexp.MatchString(name) {
			filteredProxies = append(filteredProxies, name)
		}
	}
	sort.Strings(filteredProxies)
	return filteredProxies
}
