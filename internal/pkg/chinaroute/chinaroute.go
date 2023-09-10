// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package chinaroute

type Result struct {
	I int
	S string
}

var (
	ChinaIPS = []string{"219.141.136.12", "202.106.50.1", "221.179.155.161", "202.96.209.133", "210.22.97.1",
		"211.136.112.200", "58.60.188.222", "210.21.196.6", "120.196.165.24", "61.139.2.69", "119.6.6.6",
		"211.137.96.205", "182.96.201.97", "61.132.163.68", "211.138.180.2", "218.104.78.2", "202.98.0.68"}
	names = []string{"北京电信", "北京联通", "北京移动", "上海电信", "上海联通", "上海移动", "广州电信", "广州联通", "广州移动",
		"成都电信", "成都联通", "成都移动", "南昌电信", "合肥电信", "合肥移动", "合肥联通", "长春联通"}
	m = map[string]string{"AS4134": "电信163   [普通线路]", "AS4809": "电信CN2  [优质线路]", "AS4837": "联通4837  [普通线路]", "AS9929": "联通9929  [优质线路]", "AS9808": "移动CMI   [普通线路]", "AS58453": "移动CMI   [普通线路]", "AS58807": "移动CMIN2 [精品线路]"}
)
