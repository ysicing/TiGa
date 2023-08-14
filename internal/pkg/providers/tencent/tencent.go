// Copyright (c) 2023 ysicing(ysicing.me, ysicing@ysicing.cloud) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package tencent

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	tag "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tag/v20180813"
	vpc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc/v20170312"
	"github.com/ysicing/tiga/internal/pkg/providers"
)

// providerName is the name of this provider.
const providerName = "tencent"

// Tencent provider tencent struct.
type Tencent struct {
	SecretId  string
	SecretKey string
	Region    string
	cvm       *cvm.Client
	vpc       *vpc.Client
	tag       *tag.Client
}

func init() {
	providers.RegisterProvider(providerName, func() (providers.Provider, error) {
		return newProvider(), nil
	})
}

func newProvider() *Tencent {
	tencentProvider := &Tencent{}
	return tencentProvider
}

// GetProviderName returns provider name.
func (p *Tencent) GetProviderName() string {
	return providerName
}

func (p *Tencent) getCredential() *common.Credential {
	return common.NewCredential(
		p.SecretId,
		p.SecretKey,
	)
}
