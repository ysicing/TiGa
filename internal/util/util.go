// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package util

import (
	"fmt"
	"math"
	"time"
)

func Traffic(k int64, f ...float64) string {
	value := 1024.0
	if len(f) > 0 {
		value = f[0]
	}
	t := math.Round(float64(k)/value/value*100) / 100
	tunit := "MB"
	if t >= value {
		t = t / value
		tunit = "GB"
	}
	if t >= value {
		t = t / value
		tunit = "TB"
	}
	return fmt.Sprintf("%.2f%v", t, tunit)
}

func PtrFormatTime(t *time.Time, layout string) string {
	if t == nil {
		return ""
	}
	return t.Format(layout)
}
