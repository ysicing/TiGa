// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package nnr

import (
	"github.com/cockroachdb/errors"
	"github.com/imroc/req/v3"
	"github.com/ysicing/tiga/common"
)

type Option struct {
	*req.Request
}

func New(token string) *Option {
	reqClient := req.C().
		SetUserAgent(common.GetUG()).
		R().
		SetHeader("accept", "application/json").
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"Token":        token,
		})
	return &Option{reqClient}
}

func (o *Option) ListServers() ([]Server, error) {
	var serversResp ServersResp
	_, err := o.SetSuccessResult(&serversResp).Post("https://nnr.moe/api/servers")
	if err != nil {
		return nil, err
	}
	if serversResp.Status != 1 {
		return nil, errors.New("list servers failed")
	}
	return serversResp.Data, nil
}

func (o *Option) ListRules() ([]Rule, error) {
	var rulesResp RulesResp
	_, err := o.SetSuccessResult(&rulesResp).Post("https://nnr.moe/api/rules")
	if err != nil {
		return nil, err
	}
	if rulesResp.Status != 1 {
		return nil, errors.New("list rules failed")
	}
	return rulesResp.Data, nil
}
