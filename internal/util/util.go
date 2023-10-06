// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package util

import (
	"fmt"
	"math"
)

func Traffic(k int64) string {
	t := math.Round(float64(k)/1024.0/1024.0*100) / 100
	tunit := "MB"
	if t >= 1024.0 {
		t = t / 1024.0
		tunit = "GB"
	}
	if t >= 1024.0 {
		t = t / 1024.0
		tunit = "TB"
	}
	return fmt.Sprintf("%.2f%v", t, tunit)
}
