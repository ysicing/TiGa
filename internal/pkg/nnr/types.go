// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package nnr

type Server struct {
	Sid    string   `json:"sid,omitempty"`
	Name   string   `json:"name,omitempty"`
	Host   string   `json:"host,omitempty"`
	Mf     int      `json:"mf,omitempty"`
	Level  int      `json:"level,omitempty"`
	Top    int      `json:"top,omitempty"`
	Status int      `json:"status,omitempty"`
	Detail string   `json:"detail,omitempty"`
	Types  []string `json:"types,omitempty"`
}

type ServersResp struct {
	Status int      `json:"status,omitempty"`
	Data   []Server `json:"data,omitempty"`
}

type Rule struct {
	Rid     string `json:"rid,omitempty"`
	Uid     string `json:"uid,omitempty"`
	Sid     string `json:"sid,omitempty"`
	Host    string `json:"host,omitempty"`
	Port    int    `json:"port,omitempty"`
	Remote  string `json:"remote,omitempty"`
	RPort   int    `json:"rport,omitempty"`
	Type    string `json:"type,omitempty"`
	Status  int    `json:"status,omitempty"`
	Name    string `json:"name,omitempty"`
	Traffic int64  `json:"traffic,omitempty"`
	Date    int64  `json:"date,omitempty"`
}

type RulesResp struct {
	Status int    `json:"status,omitempty"`
	Data   []Rule `json:"data,omitempty"`
}

type Node struct {
	Traffic int64  `json:"traffic,omitempty"`
	Remote  string `json:"remote,omitempty"`
}

type AddRule struct {
	Sid    string `json:"sid"`
	Remote string `json:"remote"`
	RPort  int    `json:"rport"`
	Type   string `json:"type"`
	Name   string `json:"name,omitempty"`
}

type UpdateRule struct {
	Rid    string `json:"rid"`
	Remote string `json:"remote"`
	RPort  int    `json:"rport"`
	Type   string `json:"type"`
	Name   string `json:"name,omitempty"`
}

type GetOrDeleteRule struct {
	Rid string `json:"rid"`
}
