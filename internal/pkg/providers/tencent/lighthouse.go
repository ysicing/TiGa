package tencent

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
	"github.com/ysicing/tiga/internal/types"
)

func (p *Tencent) LighthouseRegion() ([]types.Region, error) {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "lighthouse.tencentcloudapi.com"
	client, _ := lighthouse.NewClient(p.getCredential(), "", cpf)
	request := lighthouse.NewDescribeRegionsRequest()
	response, err := client.DescribeRegions(request)
	if err != nil {
		return nil, err
	}
	var regions []types.Region
	for _, i := range response.Response.RegionSet {
		regions = append(regions, types.Region{
			ID:   *i.Region,
			Name: *i.RegionName,
		})
	}
	return regions, nil
}
