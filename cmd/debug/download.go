// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package debug

import (
	"net/url"
	"path/filepath"

	"github.com/cockroachdb/errors"
	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/common"
	"github.com/ysicing/tiga/internal/pkg/repo"
	"github.com/ysicing/tiga/pkg/factory"
	"github.com/ysicing/tiga/pkg/util/fileutil"
)

func DownloadCommand(f factory.Factory) *cobra.Command {
	e := repo.Entry{
		Name: "tiga",
		Desc: "tiga",
		Platforms: repo.Platforms{
			Windows: true,
			Linux:   true,
			MacOS:   true,
			Amd64:   true,
			Arm64:   true,
		},
	}
	dl := &cobra.Command{
		Use:     "download",
		Short:   "download",
		Aliases: []string{"dl"},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(e.Url) == 0 {
				return errors.New("url is empty")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			u, _ := url.Parse(e.Url)
			filename := filepath.Base(u.Path)
			dlPath := common.GetDefaultCacheDir() + "/" + filename
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
	dl.Flags().StringVar(&e.Url, "url", "", "download file url")
	return dl
}
