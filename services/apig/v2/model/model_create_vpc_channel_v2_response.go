package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"strings"
)

// Response Object
type CreateVpcChannelV2Response struct {
	// VPC通道的名称。  长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、中划线、下划线组成，且只能以英文或中文开头。 > 中文字符必须为UTF-8或者unicode编码。

	Name string `json:"name"`
	// VPC通道的类型。 - 1：私网ELB通道（待废弃） - 2：API网关内置支持负载均衡功能的快速通道类型

	Type int32 `json:"type"`
	// VPC通道中主机的端口号。  取值范围1 ~ 65535，仅VPC通道类型为2时有效。  VPC通道类型为2时必选。

	Port *int32 `json:"port,omitempty"`
	// 分发算法。 - 1：加权轮询（wrr） - 2：加权最少连接（wleastconn） - 3：源地址哈希（source） - 4：URI哈希（uri）  VPC通道类型为2时必选。

	BalanceStrategy *CreateVpcChannelV2ResponseBalanceStrategy `json:"balance_strategy,omitempty"`
	// VPC通道的成员类型。 - ip - ecs  VPC通道类型为2时必选。

	MemberType *CreateVpcChannelV2ResponseMemberType `json:"member_type,omitempty"`
	// VPC通道的创建时间

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
	// VPC通道的编号

	Id *string `json:"id,omitempty"`
	// VPC通道的状态。 - 1：正常 - 2：异常

	Status *int32 `json:"status,omitempty"`
	// 私网ELB通道的编号。  仅当VPC通道类型为1时生效

	ElbId          *string `json:"elb_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateVpcChannelV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateVpcChannelV2Response struct{}"
	}

	return strings.Join([]string{"CreateVpcChannelV2Response", string(data)}, " ")
}

type CreateVpcChannelV2ResponseBalanceStrategy struct {
	value int64
}

type CreateVpcChannelV2ResponseBalanceStrategyEnum struct {
	E_1 CreateVpcChannelV2ResponseBalanceStrategy
	E_2 CreateVpcChannelV2ResponseBalanceStrategy
	E_3 CreateVpcChannelV2ResponseBalanceStrategy
	E_4 CreateVpcChannelV2ResponseBalanceStrategy
}

func GetCreateVpcChannelV2ResponseBalanceStrategyEnum() CreateVpcChannelV2ResponseBalanceStrategyEnum {
	return CreateVpcChannelV2ResponseBalanceStrategyEnum{
		E_1: CreateVpcChannelV2ResponseBalanceStrategy{
			value: 1,
		}, E_2: CreateVpcChannelV2ResponseBalanceStrategy{
			value: 2,
		}, E_3: CreateVpcChannelV2ResponseBalanceStrategy{
			value: 3,
		}, E_4: CreateVpcChannelV2ResponseBalanceStrategy{
			value: 4,
		},
	}
}

func (c CreateVpcChannelV2ResponseBalanceStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateVpcChannelV2ResponseBalanceStrategy) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int64")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int64)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int64 error")
	}
}

type CreateVpcChannelV2ResponseMemberType struct {
	value string
}

type CreateVpcChannelV2ResponseMemberTypeEnum struct {
	IP  CreateVpcChannelV2ResponseMemberType
	ECS CreateVpcChannelV2ResponseMemberType
}

func GetCreateVpcChannelV2ResponseMemberTypeEnum() CreateVpcChannelV2ResponseMemberTypeEnum {
	return CreateVpcChannelV2ResponseMemberTypeEnum{
		IP: CreateVpcChannelV2ResponseMemberType{
			value: "ip",
		},
		ECS: CreateVpcChannelV2ResponseMemberType{
			value: "ecs",
		},
	}
}

func (c CreateVpcChannelV2ResponseMemberType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateVpcChannelV2ResponseMemberType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}