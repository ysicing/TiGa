// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package common

import (
	"fmt"

	"github.com/morikuni/aec"

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

// ShowLogo show logo
func ShowLogo() {
	tigaLogo := aec.BlueF.Apply(Logo)
	fmt.Println(tigaLogo)
}

// GetDefaultTiGAIndex 获取默认index配置
func GetDefaultTiGAIndex() string {
	return GetDefaultCfgDir() + "/index.yaml"
}

// GetLockCacheFile 获取lock文件
func GetLockCacheFile(name string) string {
	return fmt.Sprintf("%s/%s.lock", GetDefaultCacheDir(), name)
}

// GetDefaultCustomIndex 获取index
func GetDefaultCustomIndex(name string) string {
	return fmt.Sprintf("%s/.%s.index", GetDefaultCacheDir(), name)
}

// GetDefaultMMDB 获取mmdb
func GetDefaultMMDB() string {
	return fmt.Sprintf("%s/ipdb.mmdb", GetDefaultDataDir())
}

// GetDefaultLogFile 获取log
func GetDefaultLogFile(log string) string {
	return fmt.Sprintf("%s/%s", GetDefaultLogDir(), log)
}

// GetDefaultScriptFile 获取脚本路径
func GetCustomScriptFile(path string) string {
	return fmt.Sprintf("%s/%s", GetDefaultDataDir(), path)
}
