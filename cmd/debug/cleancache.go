// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package debug

import (
	"fmt"
	"os"

	"github.com/ergoapi/util/ztime"
	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/cmd/boot"
	"github.com/ysicing/tiga/common"
	"github.com/ysicing/tiga/pkg/factory"
)

func CleanCacheCommand(f factory.Factory) *cobra.Command {
	return &cobra.Command{
		Use:     "clean-cache",
		Aliases: []string{"cc"},
		Short:   "clean tiga cache",
		RunE: func(cmd *cobra.Command, args []string) error {
			cacheDir := common.GetDefaultCacheDir()
			if err := os.Rename(cacheDir, fmt.Sprintf("/tmp/tigacache.%s", ztime.GetTodayMin())); err != nil {
				return err
			}
			_ = boot.OnBoot()
			f.GetLog().Debugf("rebuilt cache dir")
			f.GetLog().Donef("clean cache %s success", cacheDir)
			if err := os.Remove(common.GetDefaultMMDB()); err == nil || os.IsNotExist(err) {
				f.GetLog().Donef("clean mmdb success")
			}
			f.GetLog().Done("clean cache success")
			return nil
		},
	}
}
