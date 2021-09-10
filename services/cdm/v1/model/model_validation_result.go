package model

import (
	"encoding/json"

	"strings"
)

type ValidationResult struct {
	// 创建或更新连接校验结果，请参见linkConfig参数说明

	LinkConfig *[]ValidationLinkConfig `json:"linkConfig,omitempty"`
}

func (o ValidationResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ValidationResult struct{}"
	}

	return strings.Join([]string{"ValidationResult", string(data)}, " ")
}
