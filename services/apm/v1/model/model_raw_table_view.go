package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 原始数据视图
type RawTableView struct {

	// 视图类型
	ViewType *RawTableViewViewType `json:"view_type,omitempty"`

	// 采集器名称
	CollectorName *string `json:"collector_name,omitempty"`

	// 视图对应的指标集的名称
	MetricSet *string `json:"metric_set,omitempty"`

	// 图标所需展示的标题
	Title *string `json:"title,omitempty"`

	// 表格的方向，H：默认，表头横向，V：表头纵向
	TableDirection *RawTableViewTableDirection `json:"table_direction,omitempty"`

	// 分组规则
	GroupBy *string `json:"group_by,omitempty"`

	// 过滤列表模型
	Filter *string `json:"filter,omitempty"`

	// 所需展示的字段列表模型集合
	FieldItemList *[]FieldItem `json:"field_item_list,omitempty"`

	// 跨度
	Span *bool `json:"span,omitempty"`

	// 跨度字段
	SpanField *string `json:"span_field,omitempty"`

	// 排序规则
	OrderBy *string `json:"order_by,omitempty"`

	// 是否只展示最近一笔数据
	Latest *bool `json:"latest,omitempty"`
}

func (o RawTableView) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RawTableView struct{}"
	}

	return strings.Join([]string{"RawTableView", string(data)}, " ")
}

type RawTableViewViewType struct {
	value string
}

type RawTableViewViewTypeEnum struct {
	TREND    RawTableViewViewType
	SUMTABLE RawTableViewViewType
	RAWTABLE RawTableViewViewType
}

func GetRawTableViewViewTypeEnum() RawTableViewViewTypeEnum {
	return RawTableViewViewTypeEnum{
		TREND: RawTableViewViewType{
			value: "trend",
		},
		SUMTABLE: RawTableViewViewType{
			value: "sumtable",
		},
		RAWTABLE: RawTableViewViewType{
			value: "rawtable",
		},
	}
}

func (c RawTableViewViewType) Value() string {
	return c.value
}

func (c RawTableViewViewType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RawTableViewViewType) UnmarshalJSON(b []byte) error {
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

type RawTableViewTableDirection struct {
	value string
}

type RawTableViewTableDirectionEnum struct {
	H RawTableViewTableDirection
	V RawTableViewTableDirection
}

func GetRawTableViewTableDirectionEnum() RawTableViewTableDirectionEnum {
	return RawTableViewTableDirectionEnum{
		H: RawTableViewTableDirection{
			value: "H",
		},
		V: RawTableViewTableDirection{
			value: "V",
		},
	}
}

func (c RawTableViewTableDirection) Value() string {
	return c.value
}

func (c RawTableViewTableDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RawTableViewTableDirection) UnmarshalJSON(b []byte) error {
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
