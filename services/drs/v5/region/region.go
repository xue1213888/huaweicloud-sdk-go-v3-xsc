package region

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

var CN_NORTH_4 = region.NewRegion("cn-north-4", "https://drs.cn-north-4.myhuaweicloud.com")
var CN_NORTH_1 = region.NewRegion("cn-north-1", "https://drs.cn-north-1.myhuaweicloud.com")
var CN_SOUTH_1 = region.NewRegion("cn-south-1", "https://drs.cn-south-1.myhuaweicloud.com")
var CN_EAST_3 = region.NewRegion("cn-east-3", "https://drs.cn-east-3.myhuaweicloud.com")
var CN_EAST_2 = region.NewRegion("cn-east-2", "https://drs.cn-east-2.myhuaweicloud.com")
var CN_NORTH_2 = region.NewRegion("cn-north-2", "https://drs.cn-north-2.myhuaweicloud.com")
var AP_SOUTHEAST_3 = region.NewRegion("ap-southeast-3", "https://drs.ap-southeast-3.myhuaweicloud.com")
var AP_SOUTHEAST_1 = region.NewRegion("ap-southeast-1", "https://drs.ap-southeast-1.myhuaweicloud.com")
var AP_SOUTHEAST_2 = region.NewRegion("ap-southeast-2", "https://drs.ap-southeast-2.myhuaweicloud.com")
var SA_BRAZIL_1 = region.NewRegion("sa-brazil-1", "https://drs.sa-brazil-1.myhuaweicloud.com")
var LA_SOUTH_2 = region.NewRegion("la-south-2", "https://drs.la-south-2.myhuaweicloud.com")
var LA_NORTH_2 = region.NewRegion("la-north-2", "https://drs.la-north-2.myhuaweicloud.com")
var NA_MEXICO_1 = region.NewRegion("na-mexico-1", "https://drs.na-mexico-1.myhuaweicloud.com")
var AF_SOUTH_1 = region.NewRegion("af-south-1", "https://drs.af-south-1.myhuaweicloud.com")
var CN_NORTH_9 = region.NewRegion("cn-north-9", "https://drs.cn-north-9.myhuaweicloud.com")
var CN_SOUTHWEST_2 = region.NewRegion("cn-southwest-2", "https://drs.cn-southwest-2.myhuaweicloud.com")
var CN_SOUTH_2 = region.NewRegion("cn-south-2", "https://drs.cn-south-2.myhuaweicloud.com")
var CN_SOUTH_4 = region.NewRegion("cn-south-4", "https://drs.cn-south-4.myhuaweicloud.com")

var staticFields = map[string]*region.Region{
	"cn-north-4":     CN_NORTH_4,
	"cn-north-1":     CN_NORTH_1,
	"cn-south-1":     CN_SOUTH_1,
	"cn-east-3":      CN_EAST_3,
	"cn-east-2":      CN_EAST_2,
	"cn-north-2":     CN_NORTH_2,
	"ap-southeast-3": AP_SOUTHEAST_3,
	"ap-southeast-1": AP_SOUTHEAST_1,
	"ap-southeast-2": AP_SOUTHEAST_2,
	"sa-brazil-1":    SA_BRAZIL_1,
	"la-south-2":     LA_SOUTH_2,
	"la-north-2":     LA_NORTH_2,
	"na-mexico-1":    NA_MEXICO_1,
	"af-south-1":     AF_SOUTH_1,
	"cn-north-9":     CN_NORTH_9,
	"cn-southwest-2": CN_SOUTHWEST_2,
	"cn-south-2":     CN_SOUTH_2,
	"cn-south-4":     CN_SOUTH_4,
}

var provider = region.DefaultProviderChain("DRS")

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