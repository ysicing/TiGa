// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package fileutil

import (
	"fmt"
	"path"

	"github.com/cockroachdb/errors"
	"github.com/ysicing/tiga/internal/pkg/download"
	"github.com/ysicing/tiga/internal/pkg/repo"
	"github.com/ysicing/tiga/pkg/log"
)

func DownloadFile(entry *repo.Entry, decompress bool, dlPath string) (string, error) {
	if entry == nil {
		return "", errors.New("entry is nil")
	}
	if !entry.ValidateOS() {
		return "", errors.New("os is not support")
	}
	if !entry.ValidateArch() {
		return "", errors.New("arch is not support")
	}
	log.GetInstance().Infof("attempting download file: %s", entry.Url)
	res, err := download.Download(dlPath, entry.Url,
		download.WithCache(),
		download.WithDecompress(decompress),
		download.WithDescription(fmt.Sprintf("%s (%s)", entry.Desc, path.Base(entry.Url))),
		download.WithExpectedDigest(entry.Digest),
	)
	if err != nil {
		return "", fmt.Errorf("failed to download %q: %w", entry.Url, err)
	}
	log.GetInstance().Debugf("res.ValidatedDigest=%v", res.ValidatedDigest)
	switch res.Status {
	case download.StatusDownloaded:
		log.GetInstance().Infof("Downloaded %s from %q", entry.Desc, entry.Url)
	case download.StatusUsedCache:
		log.GetInstance().Infof("Using cache %q", res.CachePath)
	case download.StatusSkipped:
		log.GetInstance().Infof("Skipped download from %q", entry.Url)
	default:
		return "", errors.Newf("Unexpected result from download.Download(): %+v", res)
	}
	return res.CachePath, nil
}
