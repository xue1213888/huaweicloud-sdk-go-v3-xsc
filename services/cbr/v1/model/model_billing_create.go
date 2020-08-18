/*
 * Cbr
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// 创建参数
type BillingCreate struct {
	// [云平台，云平台，公有云或者混合云](tag:hws,hws_hk,fcs_vm,ctc) [云平台，云平台，公有云](tag:dt,ocb,tlf,sbc)
	CloudType BillingCreateCloudType `json:"cloud_type,omitempty"`
	// [规格，崩溃一致性（crash_consistent）或应用一致性（app_consistent）](tag:hws,hws_hk,fcs_vm,ctc) [规格，默认为崩溃一致性（crash_consistent）](tag:dt,ocb,tlf,sbc)
	ConsistentLevel string `json:"consistent_level"`
	// 对象类型：云服务器（server），云硬盘（disk）。
	ObjectType BillingCreateObjectType `json:"object_type"`
	// 保护类型：备份（backup）、复制(replication)
	ProtectType BillingCreateProtectType `json:"protect_type"`
	// 容量，单位GB
	Size int32 `json:"size"`
	// 创建模式，按需：post_paid，包周期：pre_paid，默认为post_paid
	ChargingMode BillingCreateChargingMode `json:"charging_mode,omitempty"`
	// 创建类型，按年(year)或者按月(month)
	PeriodType BillingCreatePeriodType `json:"period_type,omitempty"`
	// 创建类型的数量
	PeriodNum *int32 `json:"period_num,omitempty"`
	// 到期后是否自动续期，默认不续期
	IsAutoRenew *bool `json:"is_auto_renew,omitempty"`
	// 是否自动付费，默认为不自动付费
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
	// 跳转URL
	ConsoleUrl *string                  `json:"console_url,omitempty"`
	ExtraInfo  *BillbingCreateExtraInfo `json:"extra_info,omitempty"`
}

func (o BillingCreate) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BillingCreate", string(data)}, " ")
}

type BillingCreateCloudType struct {
	value string
}

type BillingCreateCloudTypeEnum struct {
	PUBLIC BillingCreateCloudType
	HYBRID BillingCreateCloudType
}

func GetBillingCreateCloudTypeEnum() BillingCreateCloudTypeEnum {
	return BillingCreateCloudTypeEnum{
		PUBLIC: BillingCreateCloudType{
			value: "public",
		},
		HYBRID: BillingCreateCloudType{
			value: "hybrid",
		},
	}
}

func (c BillingCreateCloudType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BillingCreateCloudType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type BillingCreateObjectType struct {
	value string
}

type BillingCreateObjectTypeEnum struct {
	SERVER BillingCreateObjectType
	DISK   BillingCreateObjectType
}

func GetBillingCreateObjectTypeEnum() BillingCreateObjectTypeEnum {
	return BillingCreateObjectTypeEnum{
		SERVER: BillingCreateObjectType{
			value: "server",
		},
		DISK: BillingCreateObjectType{
			value: "disk",
		},
	}
}

func (c BillingCreateObjectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BillingCreateObjectType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type BillingCreateProtectType struct {
	value string
}

type BillingCreateProtectTypeEnum struct {
	BACKUP      BillingCreateProtectType
	REPLICATION BillingCreateProtectType
}

func GetBillingCreateProtectTypeEnum() BillingCreateProtectTypeEnum {
	return BillingCreateProtectTypeEnum{
		BACKUP: BillingCreateProtectType{
			value: "backup",
		},
		REPLICATION: BillingCreateProtectType{
			value: "replication",
		},
	}
}

func (c BillingCreateProtectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BillingCreateProtectType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type BillingCreateChargingMode struct {
	value string
}

type BillingCreateChargingModeEnum struct {
	POST_PAID BillingCreateChargingMode
	PRE_PAID  BillingCreateChargingMode
}

func GetBillingCreateChargingModeEnum() BillingCreateChargingModeEnum {
	return BillingCreateChargingModeEnum{
		POST_PAID: BillingCreateChargingMode{
			value: "post_paid",
		},
		PRE_PAID: BillingCreateChargingMode{
			value: "pre_paid",
		},
	}
}

func (c BillingCreateChargingMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BillingCreateChargingMode) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type BillingCreatePeriodType struct {
	value string
}

type BillingCreatePeriodTypeEnum struct {
	YEAR  BillingCreatePeriodType
	MONTH BillingCreatePeriodType
}

func GetBillingCreatePeriodTypeEnum() BillingCreatePeriodTypeEnum {
	return BillingCreatePeriodTypeEnum{
		YEAR: BillingCreatePeriodType{
			value: "year",
		},
		MONTH: BillingCreatePeriodType{
			value: "month",
		},
	}
}

func (c BillingCreatePeriodType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BillingCreatePeriodType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
