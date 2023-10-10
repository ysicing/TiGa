// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package nnr

import (
	"fmt"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/ergoapi/util/exstr"
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

func (o *Option) ServersMap() map[string]Server {
	s, _ := o.ListServers()
	if len(s) == 0 {
		return nil
	}
	m := make(map[string]Server)
	for _, j := range s {
		m[j.Sid] = j
	}
	return m
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

func (o *Option) ListRemoteNode() (map[string]Node, error) {
	var rulesResp RulesResp
	_, err := o.SetSuccessResult(&rulesResp).Post("https://nnr.moe/api/rules")
	if err != nil {
		return nil, err
	}
	if rulesResp.Status != 1 {
		return nil, errors.New("list rules failed")
	}
	mapNodes := make(map[string]Node, 60)
	for _, rule := range rulesResp.Data {
		key := fmt.Sprintf("%s:%v", rule.Remote, rule.RPort)
		if _, ok := mapNodes[key]; ok {
			node := mapNodes[key]
			node.Traffic = node.Traffic + rule.Traffic
			mapNodes[key] = node
		} else {
			mapNodes[key] = Node{
				Remote: key,
			}
		}
	}
	return mapNodes, nil
}

func (o *Option) SortRemoteNode() ([]Node, error) {
	rNodes, err := o.ListRemoteNode()
	if err != nil {
		return nil, err
	}
	var nodes []Node
	for _, node := range rNodes {
		nodes = append(nodes, node)
	}
	return nodes, nil
}

func (o *Option) AddRule(sid, remote, t string) (*Rule, error) {
	var ruleResp RuleResp
	if t == "" {
		t = "tcp"
	}
	s := strings.Split(remote, ":")
	_, err := o.SetBody(&Rule{
		Sid:    sid,
		Remote: s[0],
		RPort:  exstr.Str2Int(s[1]),
		Type:   t,
	}).SetSuccessResult(&ruleResp).Post("https://nnr.moe/api/rules/add")
	if err != nil {
		return nil, err
	}
	if ruleResp.Status != 1 {
		return nil, errors.New("add rule failed")
	}
	return &ruleResp.Data, nil
}
