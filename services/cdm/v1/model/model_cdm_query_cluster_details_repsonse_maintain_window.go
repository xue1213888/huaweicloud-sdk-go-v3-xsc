package model

import (
	"encoding/json"

	"strings"
)

// 维护窗口
type CdmQueryClusterDetailsRepsonseMaintainWindow struct {
	// 周几

	Day *string `json:"day,omitempty"`
	// 开始时间。

	StartTime *string `json:"startTime,omitempty"`
	// 结束时间。

	EndTime *string `json:"endTime,omitempty"`
}

func (o CdmQueryClusterDetailsRepsonseMaintainWindow) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CdmQueryClusterDetailsRepsonseMaintainWindow struct{}"
	}

	return strings.Join([]string{"CdmQueryClusterDetailsRepsonseMaintainWindow", string(data)}, " ")
}
