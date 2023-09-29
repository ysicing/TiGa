// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package static

import (
	"github.com/ysicing/tiga/common"
	"github.com/ysicing/tiga/internal/static/manifests"
	"github.com/ysicing/tiga/internal/static/metadata"
)

func StageFiles() error {
	dataDir := common.GetDefaultDataDir()
	if err := manifests.Stage(dataDir); err != nil {
		return err
	}
	if err := metadata.Stage(dataDir); err != nil {
		return err
	}
	return nil
}
