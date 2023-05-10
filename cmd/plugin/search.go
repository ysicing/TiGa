// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package plugin

import (
	"github.com/ergoapi/util/file"
	"github.com/ysicing/tiga/common"
)

type SearchOptions struct {
	RepoFile string
}

type SearchResult struct {
	Name    string
	Version string
	Url     string
	Desc    string
}

func (s *SearchOptions) Run() error {
	// load repo file
	if _, err := s.buildIndex(); err != nil {
		return err
	}
	// search
	return nil
}

func (s *SearchOptions) buildIndex() (map[string]string, error) {
	pluginCfg := common.GetDefaultTiGAPluginConfig()
	if !file.CheckFileExists(pluginCfg) {
		// 下载插件配置, 重新buildIndex
		return s.buildIndex()
	}
	return nil, nil
}
