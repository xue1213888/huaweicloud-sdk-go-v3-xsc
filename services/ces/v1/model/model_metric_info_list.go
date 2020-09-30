/*
 * CES
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 指标信息
type MetricInfoList struct {
	// 指标维度
	Dimensions []MetricsDimension `json:"dimensions"`
	// 指标名称，必须以字母开头，只能包含0-9/a-z/A-Z/_，长度最短为1，最大为64。  具体指标名请参见查询指标列表中查询出的指标名。
	MetricName string `json:"metric_name"`
	// 指标命名空间，，例如弹性云服务器命名空间。格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_，总长度最短为3，最大为32。说明： 当alarm_type为（EVENT.SYS| EVENT.CUSTOM）时允许为空。
	Namespace string `json:"namespace"`
	// 指标单位。
	Unit string `json:"unit"`
}

func (o MetricInfoList) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"MetricInfoList", string(data)}, " ")
}
