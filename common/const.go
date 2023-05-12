// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package common

const (
	FileMode0777    = 0o777
	FileMode0755    = 0o755
	FileMode0644    = 0o644
	FileMode0600    = 0o600
	DefaultLogDir   = "tiga/log"
	DefaultDataDir  = "tiga/data"
	DefaultBinDir   = "tiga/bin"
	DefaultCfgDir   = "tiga/config"
	DefaultCacheDir = "tiga/cache"

	DefaultUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.0.0"
)

const Logo = `Open Source Cli Tools For Senior Restart Engineer`

type AppType string

func (a AppType) String() string {
	return string(a)
}

const (
	AppTypeDefault AppType = "bin"
	AppTypeScript  AppType = "script"
	AppTypeHelm    AppType = "helm"
	AppTypeKube    AppType = "kube"
	AppTypeDocker  AppType = "docker"
	AppTypeSystem  AppType = "system"
)
