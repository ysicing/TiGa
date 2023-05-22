// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package ipdb

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/oschwald/geoip2-golang"
	"github.com/ysicing/tiga/common"
	"github.com/ysicing/tiga/internal/pkg/download"
	"github.com/ysicing/tiga/pkg/log"
)

var (
	mmdb *geoip2.Reader
	once sync.Once
)

// downLoadMMDB download mmdb
func downLoadMMDB(path string) error {
	_, err := download.Download(path, common.DefaultMMDB)
	if err != nil {
		return err
	}
	return nil
}

func InitMMDB() error {
	logpkg := log.GetInstance()
	dbfile := common.GetDefaultMMDB()
	if file, err := os.Stat(dbfile); os.IsNotExist(err) {
		logpkg.Debugf("mmdb file not exist, downloading...")
		if err := downLoadMMDB(dbfile); err != nil {
			return err
		}
		logpkg.Info("mmdb file download success")
	} else {
		if time.Since(file.ModTime()) > 7*24*time.Hour {
			logpkg.Debugf("mmdb file expired, downloading...")
			_ = os.Remove(dbfile)
			return InitMMDB()
		}
	}
	if !Verify() {
		logpkg.Warn("mmdb invalid, remove and download")
		if err := os.Remove(dbfile); err != nil {
			return fmt.Errorf("remove invalid mmdb failed: %s", err.Error())
		}

		if err := downLoadMMDB(dbfile); err != nil {
			return fmt.Errorf("download mmdb failed: %s", err.Error())
		}
	}
	logpkg.Debugf("verify country.mmdb success")
	return nil
}

func LoadFromBytes(buffer []byte) {
	once.Do(func() {
		var err error
		mmdb, err = geoip2.FromBytes(buffer)
		if err != nil {
			log.GetInstance().Warnf("load mmdb failed: %s", err.Error())
		}
	})
}

// Verify 验证
func Verify() bool {
	instance, err := geoip2.Open(common.GetDefaultMMDB())
	if err == nil {
		instance.Close()
	}
	return err == nil
}

func Instance() *geoip2.Reader {
	once.Do(func() {
		var err error
		mmdb, err = geoip2.Open(common.GetDefaultMMDB())
		if err != nil {
			log.GetInstance().Warnf("load mmdb failed: %s", err.Error())
		}
	})
	return mmdb
}

// MatchCN 匹配中国IP
func MatchCN(src string) bool {
	ip := net.ParseIP(src)
	if ip == nil {
		return false
	}
	if ip.IsPrivate() || ip.IsLoopback() {
		return true
	}
	record, _ := Instance().Country(ip)
	return strings.EqualFold(record.Country.IsoCode, "CN")
}

// ValidateIP 验证IP
func ValidateIP(ip string) bool {
	return net.ParseIP(ip) != nil
}
