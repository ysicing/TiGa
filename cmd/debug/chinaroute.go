// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package debug

import (
	"fmt"
	"time"

	"github.com/ergoapi/util/color"
	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/internal/pkg/chinaroute"
	"github.com/ysicing/tiga/internal/pkg/myip"
	"github.com/ysicing/tiga/pkg/factory"
)

func ChinaRouteCommand(f factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test-cn-route",
		Short: "测试三网回程路由",
		Long:  `测试三网回程路由,参考https://github.com/zhanghanyun/backtrace`,
		Run: func(cmd *cobra.Command, args []string) {
			var (
				s [17]string
				c = make(chan chinaroute.Result)
				t = time.After(time.Second * 10)
			)
			t1 := time.Now()
			f.GetLog().Info("正在测试三网回程路由")
			ipinfo := myip.NewIPInfoIO().IP()
			f.GetLog().Infof("国家或地区: %s 城市: %s 服务商: %s", color.SGreen(ipinfo.Country), color.SGreen(ipinfo.City), color.SGreen(ipinfo.Org))

			for i := range chinaroute.ChinaIPS {
				go chinaroute.ChinaTrace(c, i)
			}

		loop:
			for range s {
				select {
				case o := <-c:
					s[o.I] = o.S
				case <-t:
					break loop
				}
			}

			for _, r := range s {
				fmt.Println(r)
			}

			f.GetLog().Donef("测试完成, 耗时: %vs", time.Since(t1).Seconds())
		},
	}
	return cmd
}
