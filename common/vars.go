// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package common

import "runtime"

var (
	Version       string
	BuildDate     string
	GitCommitHash string
)

var (
	ValidPrefixes = []string{"tiga", "ergo", "ysicing"}
	ListOutput    string
	ListSort      string
	SavePath      string
)

var (
	GoDefaultVersion = runtime.Version()
)
