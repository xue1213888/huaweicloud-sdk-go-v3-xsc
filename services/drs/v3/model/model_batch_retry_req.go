package model

import (
	"encoding/json"

	"strings"
)

// 批量续传任务请求列表
type BatchRetryReq struct {
	// 批量续传任务请求列表

	Jobs []RetryInfo `json:"jobs"`
}

func (o BatchRetryReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchRetryReq struct{}"
	}

	return strings.Join([]string{"BatchRetryReq", string(data)}, " ")
}
