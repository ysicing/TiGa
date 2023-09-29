// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package systemutil

import (
	"fmt"
	"os"
	"strings"

	"github.com/ysicing/tiga/common"
)

func GetOSPath() string {
	p := os.Getenv("PATH")
	if !strings.Contains(p, common.GetDefaultBinDir()) {
		p = fmt.Sprintf("%s:%s", p, common.GetDefaultBinDir())
	}
	os.Setenv("PATH", p)
	return p
}
