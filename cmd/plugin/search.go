// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package plugin

import (
	"fmt"
	"os"

	"github.com/cockroachdb/errors"
	"github.com/ergoapi/util/file"
	"github.com/sirupsen/logrus"
	"github.com/ysicing/tiga/common"
	"github.com/ysicing/tiga/internal/pkg/repo"
	"github.com/ysicing/tiga/pkg/exec"
	"github.com/ysicing/tiga/pkg/log"
)

type SearchOptions struct {
	RepoFile string
	Log      log.Logger
}

type SearchResult struct {
	Index   string         `json:"index" yaml:"index"`
	Plugins []*repo.Plugin `json:"plugins" yaml:"plugins"`
}

func (s *SearchOptions) Search(name string) ([]SearchResult, error) {
	// load repo file
	s.Log.Debugf("start load repo file")
	if err := s.buildIndex(); err != nil {
		return nil, err
	}
	indexs, err := repo.LoadIndex()
	if err != nil {
		return nil, errors.Newf("failed to load index file: %v", err)
	}
	if len(indexs.Index) == 0 {
		return nil, errors.Newf("not found any plugin in index file")
	}
	if name == "" || name == "*" || name == "all" {
		name = ""
	}

	var res []SearchResult

	for _, v := range indexs.Index {
		cachefile := fmt.Sprintf("%s/.%s.index", common.GetDefaultCacheDir(), v.Name)
		s.Log.Debugf("start search %s plugin form %s (cache fiel: %s)", name, v.Name, cachefile)
		ps, err := repo.LoadPlugin(cachefile)
		if err != nil {
			s.Log.Warnf("failed to load %s plugin file: %v", v.Name, err)
			continue
		}
		if len(ps.Plugins) > 0 {
			if name != "" {
				if ps.Has(name) {
					res = append(res, SearchResult{
						Index:   v.Name,
						Plugins: []*repo.Plugin{ps.Get(name)},
					})
				}
			} else {
				res = append(res, SearchResult{
					Index:   v.Name,
					Plugins: ps.Plugins,
				})
			}
		}
	}
	return res, nil
}

func (s *SearchOptions) buildIndex() error {
	pluginCfg := common.GetDefaultTiGAPluginConfig()
	if !file.CheckFileExists(pluginCfg) {
		// 下载插件配置, 重新buildIndex
		s.Log.Debugf("not found index file, start download default plugin index")
		indexData, err := common.RepoIndex.ReadFile("index.yaml")
		if err != nil {
			return errors.Newf("failed to read default plugin index file: %v", err)
		}
		if err := file.WriteToFile(pluginCfg, indexData); err != nil {
			return errors.Newf("failed to write default plugin index file: %v", err)
		}
		return s.buildIndex()
	}
	// load index
	args := []string{"repo", "update"}
	if s.Log.GetLevel() == logrus.DebugLevel {
		args = append(args, "--debug")
	}
	return exec.CommandRun(os.Args[0], args...)
}
