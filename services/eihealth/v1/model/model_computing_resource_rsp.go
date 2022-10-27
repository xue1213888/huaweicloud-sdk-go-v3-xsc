package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComputingResourceRsp struct {

	// 实例ID
	Id string `json:"id"`

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 实例名称
	Name string `json:"name"`

	Spec *ComputingSpecDto `json:"spec,omitempty"`

	// IP地址
	Ip string `json:"ip"`

	SystemDisk *DiskDto `json:"system_disk,omitempty"`

	// 数据盘
	DataDisks *[]DiskDto `json:"data_disks,omitempty"`

	Image *ImageDto `json:"image,omitempty"`

	// 计费模式
	ChargeMode string `json:"charge_mode"`

	// 购买周期
	PeriodNum *string `json:"period_num,omitempty"`

	// 购买时间
	CreateTime string `json:"create_time"`

	// 状态
	Status string `json:"status"`

	// 可用区
	AvailabilityZoneId string `json:"availability_zone_id"`

	// 资源是否可调度
	Schedulable *bool `json:"schedulable,omitempty"`
}

func (o ComputingResourceRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComputingResourceRsp struct{}"
	}

	return strings.Join([]string{"ComputingResourceRsp", string(data)}, " ")
}