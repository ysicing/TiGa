// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package repo

import (
	"github.com/cockroachdb/errors"
	"os"
	"sigs.k8s.io/yaml"
	"time"
)

type File struct {
	Generated    time.Time `json:"generated"`
	Repositories []*Entry  `json:"repositories"`
}

type Entry struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Desc string `json:"desc"`
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
	return f, nil
}
