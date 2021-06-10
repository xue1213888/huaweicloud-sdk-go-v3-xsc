package model

import (
	"encoding/json"

	"strings"
)

// 容灾任务监控数据响应体
type QueryDataGuardMonitorAndChartResp struct {
	// 任务id

	Id string `json:"id"`

	DataGuardMinitor *QueryDataGuardMonitorResponse `json:"data_guard_minitor"`
}

func (o QueryDataGuardMonitorAndChartResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QueryDataGuardMonitorAndChartResp struct{}"
	}

	return strings.Join([]string{"QueryDataGuardMonitorAndChartResp", string(data)}, " ")
}
