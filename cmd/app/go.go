// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package app

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/ysicing/tiga/internal/pkg/download"

	"github.com/ysicing/tiga/internal/util/netutil"

	"github.com/ergoapi/util/color"
	"github.com/ysicing/tiga/common"

	"github.com/cockroachdb/errors"

	gv "github.com/ergoapi/util/version"
	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/pkg/exec"
	"github.com/ysicing/tiga/pkg/factory"
)

const (
	goexmapleInstall = `
  # Install Go
  tiga app install go --version v1.20.4
`
)

func installGo(f factory.Factory) *cobra.Command {
	var version string
	logpkg := f.GetLog()
	cmd := &cobra.Command{
		Use:     "go",
		Short:   "Install Go",
		Long:    `Install Go programming language and SDK.`,
		Example: goexmapleInstall,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			installV := strings.ReplaceAll(version, "v", "")
			installV = strings.ReplaceAll(installV, "go", "")
			if _, err := gv.Parse(installV); err != nil {
				logpkg.Debugf("installV: %v, err: %v", installV, err)
				return errors.Newf("invalid go version %s, recommend %s", version, color.SGreen(common.GoDefaultVersion))
			}
			defaultV := strings.ReplaceAll(common.GoDefaultVersion, "go", "")
			if gv.LT(installV, defaultV) {
				return errors.Newf("now go stable version is %s", color.SGreen(common.GoDefaultVersion))
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			logpkg.Infof("install go: %v", version)
			dlURL := fmt.Sprintf("https://go.dev/dl/%s.linux-%s.tar.gz", version, runtime.GOARCH)
			if netutil.ValidChinaNetwork() {
				dlURL = fmt.Sprintf("https://golang.google.cn/dl/%s.linux-%s.tar.gz", version, runtime.GOARCH)
			}
			logpkg.Debugf("download url: %v", dlURL)
			cachefile := fmt.Sprintf("%s/%s.linux-%s.tar.gz", common.GetDefaultCacheDir(), version, runtime.GOARCH)
			download.Download(cachefile, dlURL, download.WithCache())
			return exec.CommandRun(common.GetCustomScriptFile("hack/manifests/system/go.install.sh"), cachefile)
		},
	}
	cmd.Flags().StringVarP(&version, "version", "v", common.GoDefaultVersion, "install go version")
	return cmd
}
