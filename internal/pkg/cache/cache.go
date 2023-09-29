// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package cache

import (
	"os"
	"sync"
	"time"

	"github.com/ysicing/tiga/common"
	"github.com/ysicing/tiga/pkg/log"
	"go.etcd.io/bbolt"
)

var (
	initOnce     sync.Once
	defaultCache *CacheFile
)

// CacheFile store and update cache file
type CacheFile struct {
	DB *bbolt.DB
}

func initCache() {
	logpkg := log.GetInstance()
	options := bbolt.Options{Timeout: time.Second}
	db, err := bbolt.Open(common.GetDefaultTiGACache(), common.FileMode0600, &options)
	switch err {
	case bbolt.ErrInvalid, bbolt.ErrChecksum, bbolt.ErrVersionMismatch:
		if err = os.Remove(common.GetDefaultTiGACache()); err != nil {
			logpkg.Warnf("remove cache file failed: %v", err)
			break
		}
		logpkg.Infof("remove invalid cache file and create new one")
		db, err = bbolt.Open(common.GetDefaultTiGACache(), common.FileMode0600, &options)
	}
	if err != nil {
		logpkg.Warnf("open cache file failed: %v", err)
	}

	defaultCache = &CacheFile{DB: db}
}

func Cache() *CacheFile {
	initOnce.Do(initCache)
	return defaultCache
}
