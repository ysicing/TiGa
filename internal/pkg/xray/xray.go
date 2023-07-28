// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package xray

import (
	"context"
	"math"
	"strings"

	"github.com/ergoapi/util/ztime"
	"github.com/sirupsen/logrus"
	statsService "github.com/xtls/xray-core/app/stats/command"
	"google.golang.org/grpc"
)

type XrayController struct {
	StatsClient statsService.StatsServiceClient
	CmdConn     *grpc.ClientConn
}

func (xrayCtl *XrayController) Init(api string) (err error) {
	xrayCtl.CmdConn, err = grpc.Dial(api, grpc.WithInsecure())
	if err != nil {
		return err
	}
	xrayCtl.StatsClient = statsService.NewStatsServiceClient(xrayCtl.CmdConn)
	return
}

func int2float(k int64) float64 {
	return math.Round(float64(k)/1024.0/1024.0*100) / 100
}

func (xrayCtl *XrayController) QueryTraffic(reset bool) error {
	resp, err := xrayCtl.StatsClient.QueryStats(context.Background(), &statsService.QueryStatsRequest{
		// 是否重置流量信息(true, false)，即完成查询后是否把流量统计归零
		Reset_: reset, // reset traffic data everytime
	})
	if err != nil {
		return err
	}

	for _, st := range resp.Stat {
		s := strings.Split(st.GetName(), ">>>")
		v := int2float(st.GetValue())
		logrus.Infof("%s %s, name: %s, type: %s, value: %f, value1: %v", ztime.NowFormat(), s[3], s[0], s[1], v, st.GetValue())
	}
	return nil
}
