// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
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
	"github.com/ysicing/tiga/pkg/log"
)

func DownloadFile(dlURL, dlPath string) (string, error) {
	log.GetInstance().Infof("attempting download file: %s", dlURL)
	res, err := download.Download(dlPath, dlURL,
		download.WithCache(),
		download.WithDecompress(false),
		download.WithDescription(path.Base(dlURL)),
	)
	if err != nil {
		return "", fmt.Errorf("failed to download %q: %w", dlURL, err)
	}
	// log.GetInstance().Debugf("res.ValidatedDigest=%v", res.ValidatedDigest)
	switch res.Status {
	case download.StatusDownloaded:
		log.GetInstance().Debugf("downloaded %s from %s", path.Base(dlURL), dlURL)
	case download.StatusUsedCache:
		log.GetInstance().Debugf("using cache %s", res.CachePath)
	case download.StatusSkipped:
		log.GetInstance().Debugf("skipped download from %s", dlURL)
	default:
		return "", errors.Newf("Unexpected result from download.Download(): %+v", res)
	}
	return res.CachePath, nil
}
