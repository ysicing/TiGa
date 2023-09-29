// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package boot

import (
	"os"

	"github.com/ergoapi/util/environ"
	"github.com/ysicing/tiga/internal/static"
	"github.com/ysicing/tiga/pkg/util/ipdb"

	"github.com/cockroachdb/errors"
	"github.com/ysicing/tiga/common"
)

var rootDirs = []string{
	common.GetDefaultLogDir(),
	common.GetDefaultDataDir(),
	common.GetDefaultBinDir(),
	common.GetDefaultCfgDir(),
	common.GetDefaultCacheDir(),
}

func initRootDirectory() error {
	for _, dir := range rootDirs {
		if err := os.MkdirAll(dir, common.FileMode0755); err != nil {
			return errors.Errorf("failed to mkdir %s, err: %s", dir, err)
		}
	}
	if err := static.StageFiles(); err != nil {
		return errors.Errorf("failed to stage files, err: %s", err)
	}
	return nil
}

func OnBoot() error {
	if environ.GetEnv("TIGA_SKIP_IPDB", "false") == "false" {
		if err := ipdb.InitMMDB(); err != nil {
			return err
		}
	}
	return initRootDirectory()
}
