package model

import (
	"encoding/json"

	"strings"
)

// 批量主备倒换请求体
type BatchSwitchoverReq struct {
	// 批量主备调换的任务详情ID请求列表

	Jobs []string `json:"jobs"`
}

func (o BatchSwitchoverReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchSwitchoverReq struct{}"
	}

	return strings.Join([]string{"BatchSwitchoverReq", string(data)}, " ")
}
