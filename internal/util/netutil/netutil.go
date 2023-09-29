// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package netutil

import (
	"github.com/ergoapi/util/exnet"
	"github.com/ysicing/tiga/pkg/util/ipdb"
)

// ValidChinaNetwork check if the network is in China
func ValidChinaNetwork() bool {
	ip, _ := exnet.OutboundIPv2()
	return ipdb.MatchCN(ip)
}
