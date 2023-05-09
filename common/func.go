// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package common

import (
	"fmt"

	"github.com/ergoapi/util/zos"
)

// GetUG 获取user-agent
func GetUG() string {
	return fmt.Sprintf("%v TiGA/%v", DefaultUserAgent, Version)
}

func getOSPath() string {
	home := zos.GetHomeDir()
	if zos.IsMacOS() {
		return home + "/.config/"
	}
	return home + "/."
}

func GetDefaultLogDir() string {
	return getOSPath() + DefaultLogDir
}

func GetDefaultBinDir() string {
	return getOSPath() + DefaultBinDir
}

func GetDefaultDataDir() string {
	return getOSPath() + DefaultDataDir
}

func GetDefaultCfgDir() string {
	return getOSPath() + DefaultCfgDir
}

func GetDefaultCacheDir() string {
	return getOSPath() + DefaultCacheDir
}

// GetDefaultTiGAConfig 获取默认配置
func GetDefaultTiGAConfig() string {
	return GetDefaultCfgDir() + "/default.yaml"
}

// GetDefaultTiGACache 获取默认cache
func GetDefaultTiGACache() string {
	return GetDefaultCacheDir() + "/cache.tiga"
}

// GetDefaultTiGAPluginConfig 获取插件配置
func GetDefaultTiGAPluginConfig() string {
	return GetDefaultCfgDir() + "/plugin.yaml"
}
