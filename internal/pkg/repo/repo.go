// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package repo

import (
	"os"
	"time"

	"github.com/cockroachdb/errors"
	"sigs.k8s.io/yaml"
)

type PluginFile struct {
	Generated time.Time `json:"generated" yaml:"generated"`
	Plugins   []*Plugin `json:"plugins" yaml:"plugins"`
}

type Plugin struct {
	Name    string `json:"name" yaml:"name"`
	Version string `json:"version" yaml:"version"`
	Type    string `json:"type" yaml:"type"`
	Url     string `json:"url,omitempty" yaml:"url,omitempty"`
}

func NewPlugin() *PluginFile {
	return &PluginFile{
		Generated: time.Now(),
		Plugins:   []*Plugin{},
	}
}

func LoadPlugin(path string) (*PluginFile, error) {
	f := new(PluginFile)
	b, err := os.ReadFile(path)
	if err != nil {
		return f, errors.Newf("failed to read file(%s):%v", path, err)
	}
	err = yaml.Unmarshal(b, f)
	return f, err
}

// Has returns true if the given name is already a repository name.
func (r *PluginFile) Has(name string) bool {
	entry := r.Get(name)
	return entry != nil
}

// Get returns an entry with the given name if it exists, otherwise returns nil
func (r *PluginFile) Get(name string) *Plugin {
	for _, plugin := range r.Plugins {
		if plugin.Name == name {
			return plugin
		}
	}
	return nil
}
