// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package myip

import (
	"strings"

	"github.com/imroc/req/v3"
	"github.com/ysicing/tiga/common"
)

type ipinfoio struct{}

type ipinfoioResp struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

func NewIPInfoIO() MyIP {
	return &ipinfoio{}
}

func (i *ipinfoio) IP() *IPMeta {
	var resp ipinfoioResp
	if _, err := req.C().
		SetUserAgent(common.GetUG()).
		// EnableDebugLog().
		// EnableDumpAll().
		R().
		SetSuccessResult(&resp).
		SetHeader("Accept", "application/json").
		Get("http://ipinfo.io"); err != nil {
		return &IPMeta{}
	}
	ipMeta := &IPMeta{
		IP:      resp.IP,
		RDns:    resp.Hostname,
		City:    resp.City,
		Region:  resp.Region,
		Country: resp.Country,
		Org:     resp.Org,
	}

	asn := strings.Split(resp.Org, " ")[0]
	if strings.Contains(asn, "AS") {
		ipMeta.Asn = asn
	}

	return ipMeta
}
