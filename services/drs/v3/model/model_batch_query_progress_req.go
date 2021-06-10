package model

import (
	"encoding/json"

	"strings"
)

// 批量查询同步进度的请求体
type BatchQueryProgressReq struct {
	// 批量查询进度任务ID请求列表

	Jobs []string `json:"jobs"`
}

func (o BatchQueryProgressReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchQueryProgressReq struct{}"
	}

	return strings.Join([]string{"BatchQueryProgressReq", string(data)}, " ")
}
