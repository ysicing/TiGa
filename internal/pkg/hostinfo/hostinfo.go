// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package hostinfo

import (
	"fmt"

	"github.com/ysicing/tiga/common"
	hinfo "tailscale.com/hostinfo"
	"tailscale.com/tailcfg"
)

// New returns a partially populated Hostinfo for the current host.
func New() *tailcfg.Hostinfo {
	t := hinfo.New()
	t.IPNVersion = fmt.Sprintf("%s-%s", common.Version, common.GitCommitHash)
	return t
}
