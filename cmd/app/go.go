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
	"github.com/ysicing/tiga/pkg/log"
)

const (
	goexmapleInstall = `
  # Install Go
  tiga app install go --version v1.20.4
`
)

func InstallGo() *cobra.Command {
	var version string
	cmd := &cobra.Command{
		Use:     "go",
		Short:   "Install Go",
		Long:    `Install Go programming language and SDK.`,
		Example: goexmapleInstall,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			installV := strings.ReplaceAll(version, "v", "")
			if _, err := gv.Parse(installV); err != nil {
				return errors.Newf("invalid go version %s, recommend %s", version, color.SGreen(common.GoDefaultVersion))
			}
			defaultV := strings.ReplaceAll(common.GoDefaultVersion, "go", "")
			if gv.LT(installV, defaultV) {
				return errors.Newf("now go stable version is %s", color.SGreen(common.GoDefaultVersion))
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			logpkg := log.GetInstance()
			logpkg.Infof("install go: %v", version)
			dlURL := fmt.Sprintf("https://go.dev/dl/%s.linux-%s.tar.gz", version, runtime.GOARCH)
			if netutil.ValidChinaNetwork() {
				dlURL = fmt.Sprintf("https://golang.google.cn/dl/%s.linux-%s.tar.gz", version, runtime.GOARCH)
			}
			logpkg.Debugf("download url: %v", dlURL)
			download.Download("", dlURL, download.WithCache())
		},
	}
	cmd.Flags().StringVarP(&version, "version", "v", common.GoDefaultVersion, "install go version")
	return cmd
}
