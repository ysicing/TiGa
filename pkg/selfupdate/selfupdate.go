// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package selfupdate

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"github.com/inconshreveable/go-update"
	"github.com/ysicing/tiga/pkg/log"
)

type Updater struct{}

func DefaultUpdater() *Updater {
	return &Updater{}
}

func UpdateTo(log log.Logger, assetURL, cmdPath string) error {
	up := DefaultUpdater()
	src, err := up.downloadDirectlyFromURL(assetURL)
	if err != nil {
		return err
	}
	defer src.Close()
	return uncompressAndUpdate(log, src, assetURL, cmdPath)
}

func (up *Updater) downloadDirectlyFromURL(assetURL string) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", assetURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request to %s: %s", assetURL, err)
	}

	req.Header.Add("Accept", "application/octet-stream")
	req = req.WithContext(context.Background())
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to download a release file from %s: %s", assetURL, err)
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to download a release file from %s: Not successful status %d", assetURL, res.StatusCode)
	}

	return res.Body, nil
}

func uncompressAndUpdate(log log.Logger, src io.Reader, assetURL, cmdPath string) error {
	_, cmd := filepath.Split(cmdPath)
	asset, err := UncompressCommand(log, src, assetURL, cmd)
	if err != nil {
		return err
	}

	log.Debugf("will upgrade %s to the latest downloaded from %s", cmdPath, assetURL)
	return update.Apply(asset, update.Options{
		TargetPath: cmdPath,
	})
}
