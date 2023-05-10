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
	"github.com/ergoapi/util/zos"
	"github.com/opencontainers/go-digest"
	"sigs.k8s.io/yaml"
)

type File struct {
	Generated    time.Time `json:"generated" yaml:"generated"`
	Repositories []*Entry  `json:"repositories" yaml:"repositories"`
}

type Entry struct {
	Name      string        `json:"name" yaml:"name"`
	Url       string        `json:"url" yaml:"url"`
	Desc      string        `json:"desc" yaml:"desc"`
	Digest    digest.Digest `json:"digest" yaml:"digest"`
	Platforms Platforms     `json:"platforms" yaml:"platforms"`
}

type Platforms struct {
	Windows bool `json:"windows" yaml:"windows"`
	Linux   bool `json:"linux" yaml:"linux"`
	MacOS   bool `json:"macos" yaml:"macos"`
	Amd64   bool `json:"amd64" yaml:"amd64"`
	Arm64   bool `json:"arm64" yaml:"arm64"`
}

func (e *Entry) ValidateArch() bool {
	if e.Platforms.Amd64 && runtime.GOARCH == "amd64" {
		return true
	}
	if e.Platforms.Arm64 && runtime.GOARCH == "arm64" {
		return true
	}
	return false
}

func (e *Entry) ValidateOS() bool {
	if zos.IsMacOS() {
		return e.Platforms.MacOS
	}
	if zos.IsLinux() {
		return e.Platforms.Linux
	}
	if zos.NotUnix() {
		return e.Platforms.Windows
	}
	return false
}

func NewFile() *File {
	return &File{
		Generated:    time.Now(),
		Repositories: []*Entry{},
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
