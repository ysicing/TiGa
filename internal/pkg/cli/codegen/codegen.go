// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/BeidouCloudPlatform/go-bindata/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	os.Unsetenv("GOPATH")
	bc := &bindata.Config{
		Input: []bindata.InputConfig{
			{
				Path:      "hack/manifests",
				Recursive: true,
			},
		},
		Package:    "manifests",
		NoCompress: true,
		NoMemCopy:  true,
		NoMetadata: true,
		Output:     "internal/static/manifests/zz_generated_bindata.go.go",
	}
	if err := bindata.Translate(bc); err != nil {
		logrus.Fatal(err)
	}
}
