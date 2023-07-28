// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package xray

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/ergoapi/util/exstr"
	handlerService "github.com/xtls/xray-core/app/proxyman/command"
	statsService "github.com/xtls/xray-core/app/stats/command"
	"github.com/ysicing/tiga/pkg/log"
	"google.golang.org/grpc"
)

type DataType string

var (
	UserDataType           DataType = "user"
	GlobalInBoundDataType  DataType = "inbound"
	GlobalOutBoundDataType DataType = "outbound"
)

type XrayController struct {
	StatsClient   statsService.StatsServiceClient
	HandlerClient handlerService.HandlerServiceClient
	CmdConn       *grpc.ClientConn
	Log           log.Logger
}

type Traffic struct {
	DataType DataType `json:"data_type" yaml:"data_type"`
	Type     string   `json:"type" yaml:"type"`
	Up       string   `json:"up" yaml:"up"`
	Down     string   `json:"down" yaml:"down"`
	Name     string   `json:"name" yaml:"name"`
}

func (xrayCtl *XrayController) Init(api string) (err error) {
	xrayCtl.Log = log.GetInstance()
	xrayCtl.CmdConn, err = grpc.Dial(api, grpc.WithInsecure())
	if err != nil {
		return err
	}
	xrayCtl.StatsClient = statsService.NewStatsServiceClient(xrayCtl.CmdConn)
	xrayCtl.HandlerClient = handlerService.NewHandlerServiceClient(xrayCtl.CmdConn)
	return
}

func int2float(k int64) float64 {
	return math.Round(float64(k)/1024.0/1024.0*100) / 100
}

func (xrayCtl *XrayController) listTraffic(reset bool) ([]string, error) {
	var s []string
	resp, err := xrayCtl.StatsClient.QueryStats(context.Background(), &statsService.QueryStatsRequest{
		// 是否重置流量信息(true, false)，即完成查询后是否把流量统计归零
		Reset_: reset, // reset traffic data everytime
	})
	if err != nil {
		return nil, err
	}
	for _, st := range resp.Stat {
		k := strings.ReplaceAll(st.GetName(), ">>>traffic>>>uplink", "")
		k = strings.ReplaceAll(k, ">>>traffic>>>downlink", "")
		s = append(s, k)
		xrayCtl.Log.Debugf("sql: %v", s)
	}
	return exstr.DuplicateStrElement(s), nil
}

func (xrayCtl *XrayController) querySimpleTraffic(reset bool, sql string) int64 {
	resp, err := xrayCtl.StatsClient.QueryStats(context.Background(), &statsService.QueryStatsRequest{
		// 是否重置流量信息(true, false)，即完成查询后是否把流量统计归零
		Reset_:  reset, // reset traffic data everytime
		Pattern: sql,
	})
	if err != nil {
		return 0
	}
	xrayCtl.Log.Debugf("sql: %v", sql)
	stat := resp.GetStat()
	if len(stat) == 0 {
		return 0
	}
	xrayCtl.Log.Debugf("sql: %v, value: %v", sql, stat[0].GetValue())
	return stat[0].GetValue()
}

func (xrayCtl *XrayController) QueryTraffic(reset bool) ([]Traffic, error) {
	list, err := xrayCtl.listTraffic(reset)
	if err != nil {
		return nil, err
	}
	var ts []Traffic
	for _, l := range list {
		l1 := strings.Split(l, ">>>")
		t := Traffic{
			Type: "global",
			Name: l1[1],
		}
		if l1[0] == string(GlobalInBoundDataType) {
			t.DataType = GlobalInBoundDataType
		} else if l1[0] == string(GlobalOutBoundDataType) {
			t.DataType = GlobalOutBoundDataType
		} else {
			t.DataType = UserDataType
			t.Type = "user"
		}
		up := int2float(xrayCtl.querySimpleTraffic(reset, l+">>>traffic>>>uplink"))
		down := int2float(xrayCtl.querySimpleTraffic(reset, l+">>>traffic>>>downlink"))
		if up < 0.001 && down < 0.001 && t.DataType == GlobalOutBoundDataType {
			continue
		}
		t.Up = fmt.Sprintf("%vMB", up)
		t.Down = fmt.Sprintf("%vMB", down)
		ts = append(ts, t)
	}
	return ts, nil
}
