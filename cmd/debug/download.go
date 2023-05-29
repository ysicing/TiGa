// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package debug

import (
	"net/url"
	"path/filepath"
	"runtime"

	"github.com/cockroachdb/errors"
	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/common"
	"github.com/ysicing/tiga/internal/pkg/repo"
	"github.com/ysicing/tiga/pkg/factory"
	"github.com/ysicing/tiga/pkg/util/fileutil"
)

func DownloadCommand(f factory.Factory) *cobra.Command {
	var dlUrl string
	dl := &cobra.Command{
		Use:     "download",
		Short:   "download",
		Aliases: []string{"dl"},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(dlUrl) == 0 {
				return errors.New("url is empty")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			u, _ := url.Parse(dlUrl)
			filename := filepath.Base(u.Path)
			dlPath := common.GetDefaultCacheDir() + "/" + filename
			e := repo.Plugin{
				Name: "tiga",
				Desc: "test download",
				Platforms: []repo.Platform{
					{
						OS:   runtime.GOOS,
						Arch: runtime.GOARCH,
						URL:  dlUrl,
					},
				},
			}
			cacheFile, err := fileutil.DownloadFile(&e, dlPath)
			if err != nil {
				return err
			}
			if len(cacheFile) == 0 {
				f.GetLog().Donef("skip downloaded, found %s", dlPath)
			} else {
				f.GetLog().Donef("downloaded success to %s(%s)", dlPath, cacheFile)
			}
			return nil
		},
	}
	dl.Flags().StringVar(&dlUrl, "url", "", "download file url")
	return dl
}
