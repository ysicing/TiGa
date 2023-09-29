// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package repo

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/ysicing/tiga/internal/pkg/download"

	"github.com/gofrs/flock"

	"github.com/ergoapi/util/file"
	"github.com/ysicing/tiga/pkg/log"

	"github.com/cockroachdb/errors"
	"github.com/ysicing/tiga/common"
	"sigs.k8s.io/yaml"
)

type Indexs struct {
	Generated time.Time `json:"generated" yaml:"generated"`
	Index     []*Index  `json:"index" yaml:"index"`
}

type Index struct {
	Name string `json:"name" yaml:"name"`
	Url  string `json:"url" yaml:"url"`
	Desc string `json:"desc,omitempty" yaml:"desc,omitempty"`
}

// Add adds the given index to the index file.
func (r *Indexs) Add(i ...*Index) {
	r.Index = append(r.Index, i...)
}

// Update attempts to replace one or more repo entries in a repo file. If an
// entry with the same name doesn't exist in the repo file it will add it.
func (r *Indexs) Update(re ...*Index) {
	r.Generated = time.Now()
	for _, target := range re {
		r.update(target)
	}
}

func (r *Indexs) update(e *Index) {
	for j, index := range r.Index {
		if index.Name == e.Name {
			r.Index[j] = e
			return
		}
	}
	r.Add(e)
}

// Has returns true if the given name is already a repository name.
func (r *Indexs) Has(name string) bool {
	entry := r.Get(name)
	return entry != nil
}

// Get returns an entry with the given name if it exists, otherwise returns nil
func (r *Indexs) Get(name string) *Index {
	for _, index := range r.Index {
		if index.Name == name {
			return index
		}
	}
	return nil
}

// Remove removes the entry from the list of repositories.
func (r *Indexs) Remove(name string) bool {
	var cp []*Index
	found := false
	for _, rf := range r.Index {
		if rf.Name == name {
			found = true
			continue
		}
		cp = append(cp, rf)
	}
	r.Index = cp
	return found
}

// WriteFile writes a repositories file to the given path.
func (r *Indexs) WriteFile() error {
	data, err := yaml.Marshal(r)
	if err != nil {
		return err
	}
	path := common.GetDefaultTiGAIndex()
	if err := os.MkdirAll(filepath.Dir(path), common.FileMode0600); err != nil {
		return err
	}
	return os.WriteFile(path, data, common.FileMode0600)
}

// LoadIndex loads the index file from the given path.
func LoadIndex() (*Indexs, error) {
	f := new(Indexs)
	path := common.GetDefaultTiGAIndex()
	if !file.CheckFileExists(path) {
		log.GetInstance().Debugf("index file not exists: %s, will load default index", path)
		indexData, err := common.RepoIndex.ReadFile("index.yaml")
		if err != nil {
			return f, errors.Newf("failed to read default index file: %v", err)
		}
		if err := file.WriteToFile(path, indexData); err != nil {
			return f, errors.Newf("failed to write default index file: %v", err)
		}
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return f, errors.Newf("failed to read index file: %v", err)
	}
	err = yaml.Unmarshal(b, f)
	return f, err
}

// IsValidIndexName checks if the given index name is valid.
func IsValidIndexName(name string) (bool, error) {
	indexs, err := LoadIndex()
	if err != nil {
		return false, err
	}
	for _, index := range indexs.Index {
		if index.Name == name {
			return true, nil
		}
	}
	return false, nil
}

func AddIndex(name, path string) error {
	logpkg := log.GetInstance()
	fileLock := flock.New(common.GetLockCacheFile("index-" + name))
	lockCtx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	locked, err := fileLock.TryLockContext(lockCtx, time.Second*3)
	if err == nil && locked {
		defer fileLock.Unlock()
	}
	if err != nil {
		logpkg.Warnf("another process is currently operating on the index file. Please try again later.: %v", err)
		return nil
	}
	indexs, err := LoadIndex()
	if err != nil {
		return err
	}
	if indexs.Has(name) {
		logpkg.Warnf("index already exists: %s, will be updated.", name)
		indexs.Update(&Index{
			Name: name,
			Url:  path,
			Desc: "自定义索引",
		})
	} else {
		logpkg.Debugf("index does not exist: %s, will be added.", name)
		indexs.Add(&Index{
			Name: name,
			Url:  path,
			Desc: "自定义索引",
		})
	}
	indexs.Generated = time.Now()
	if err := UpdateIndex(name, path, true); err != nil {
		logpkg.Debugf("failed to update index: %s, %v", name, err)
	}
	return indexs.WriteFile()
}

func DeleteIndex(name string) error {
	logpkg := log.GetInstance()
	fileLock := flock.New(common.GetLockCacheFile("index-" + name))
	lockCtx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	locked, err := fileLock.TryLockContext(lockCtx, time.Second*3)
	if err == nil && locked {
		defer fileLock.Unlock()
	}
	if err != nil {
		logpkg.Warnf("another process is currently operating on the index file. Please try again later.: %v", err)
		return nil
	}
	indexs, err := LoadIndex()
	if err != nil {
		return err
	}
	if indexs.Has(name) {
		logpkg.Debugf("found index: %s, will be removed.", name)
		indexs.Remove(name)
	} else {
		return nil
	}
	indexs.Generated = time.Now()
	return indexs.WriteFile()
}

// UpdateIndexs updates the index file.
func UpdateIndexs(force bool) error {
	logpkg := log.GetInstance()
	t1 := time.Now()
	indexs, err := LoadIndex()
	if err != nil {
		return err
	}
	logpkg.Infof("Hang tight while we grab the latest from your indexs...")
	for _, index := range indexs.Index {
		t2 := time.Now()
		logpkg.Debugf("updating index: %s", index.Name)
		if err := UpdateIndex(index.Name, index.Url, force); err != nil {
			logpkg.Warnf("failed to update index: %s, error: %v", index.Name, err)
			continue
		}
		logpkg.Donef("updated index: %s, cost: %.2fs", index.Name, time.Since(t2).Seconds())
	}
	logpkg.Donef("Update Complete Cost %.2fs. ⎈Happy TiGA!⎈", time.Since(t1).Seconds())
	return nil
}

func UpdateIndex(name, url string, force bool) error {
	if name == "ysicing" && url == "" {
		url = fmt.Sprintf("%s/hack/metadata/plugin.%s.yaml", common.GetDefaultDataDir(), runtime.GOARCH)
	}
	logpkg := log.GetInstance()
	path := common.GetDefaultCustomIndex(name)
	if force {
		os.Remove(path)
	}
	if file.CheckFileExists(path) {
		fileInfo, _ := os.Stat(path)
		threshold := time.Now().Add(-5 * time.Minute)
		if fileInfo.ModTime().Before(threshold) {
			logpkg.Debugf("cache file exists and expired 5 minute, will download: %s", path)
			os.Remove(path)
		} else {
			logpkg.Debugf("cache file exists and not expired 5 minute, skip")
			return nil
		}
	}
	_, err := download.Download(path, url, download.WithCache())
	if err != nil {
		return err
	}
	return nil
}
