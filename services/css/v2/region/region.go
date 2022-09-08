package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var CN_SOUTH_1 = region.NewRegion("cn-south-1", "https://css.cn-south-1.myhuaweicloud.com")
var CN_EAST_3 = region.NewRegion("cn-east-3", "https://css.cn-east-3.myhuaweicloud.com")
var CN_NORTH_4 = region.NewRegion("cn-north-4", "https://css.cn-north-4.myhuaweicloud.com")
var CN_NORTH_2 = region.NewRegion("cn-north-2", "https://css.cn-north-2.myhuaweicloud.com")
var CN_NORTH_1 = region.NewRegion("cn-north-1", "https://es.cn-north-1.myhuaweicloud.com")
var CN_NORTH_9 = region.NewRegion("cn-north-9", "https://css.cn-north-9.myhuaweicloud.com")
var CN_EAST_2 = region.NewRegion("cn-east-2", "https://es.cn-east-2.myhuaweicloud.com")
var CN_SOUTH_4 = region.NewRegion("cn-south-4", "https://css.cn-south-4.myhuaweicloud.com")
var CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2", "https://css.cn-southwest-2.myhuaweicloud.com")
var AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2", "https://css.ap-southeast-2.myhuaweicloud.com")
var AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3", "https://css.ap-southeast-3.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"cn-south-1":     CN_SOUTH_1,
	"cn-east-3":      CN_EAST_3,
	"cn-north-4":     CN_NORTH_4,
	"cn-north-2":     CN_NORTH_2,
	"cn-north-1":     CN_NORTH_1,
	"cn-north-9":     CN_NORTH_9,
	"cn-east-2":      CN_EAST_2,
	"cn-south-4":     CN_SOUTH_4,
	"cn-southwest-2": CN_SOUTHWEST_2,
	"ap-southeast-2": AP_SOUTHEAST_2,
	"ap-southeast-3": AP_SOUTHEAST_3,
}

var provider = region.DefaultProviderChain("CSS")

func ValueOf(regionId string) *region.Region {
	if regionId == "" {
		panic("unexpected empty parameter: regionId")
	}

	reg := provider.GetRegion(regionId)
	if reg != nil {
		return reg
	}

	if _, ok := staticFields[regionId]; ok {
		return staticFields[regionId]
	}
	panic(fmt.Sprintf("unexpected regionId: %s", regionId))
}
