// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package repo

import (
	"os"
	"runtime"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/opencontainers/go-digest"
	"sigs.k8s.io/yaml"
)

type File struct {
	Generated time.Time `json:"generated" yaml:"generated"`
	Plugins   []*Plugin `json:"plugins" yaml:"plugins"`
}

type Plugin struct {
	Name      string     `json:"name" yaml:"name"`
	Home      string     `json:"home,omitempty" yaml:"home,omitempty"`
	Desc      string     `json:"desc,omitempty" yaml:"desc,omitempty"`
	Version   string     `json:"version" yaml:"version"`
	Type      string     `json:"type" yaml:"type"`
	Platforms []Platform `json:"platforms" yaml:"platforms"`
}

type Platform struct {
	OS     string        `json:"os" yaml:"os"`
	Arch   string        `json:"arch" yaml:"arch"`
	URL    string        `json:"url" yaml:"url"`
	Digest digest.Digest `json:"digest,omitempty" yaml:"digest,omitempty"`
}

func (e *Plugin) ValidateArch() bool {
	for _, a := range e.Platforms {
		if a.Arch == runtime.GOARCH {
			return true
		}
	}
	return false
}

func (e *Plugin) ValidateOS() bool {
	for _, a := range e.Platforms {
		if a.OS == runtime.GOOS {
			return true
		}
	}
	return false
}

func (e *Plugin) GetCurrentURL() string {
	for _, a := range e.Platforms {
		if a.Arch == runtime.GOARCH && a.OS == runtime.GOOS {
			return a.URL
		}
	}
	return ""
}

func NewFile() *File {
	return &File{
		Generated: time.Now(),
		Plugins:   []*Plugin{},
	}
}

func LoadFile(path string) (*File, error) {
	f := new(File)
	b, err := os.ReadFile(path)
	if err != nil {
		return f, errors.Newf("failed to read file(%s):%v", path, err)
	}
	err = yaml.Unmarshal(b, f)
	return f, err
}
