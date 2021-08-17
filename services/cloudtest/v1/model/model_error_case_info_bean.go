package model

import (
	"encoding/json"

	"strings"
)

type ErrorCaseInfoBean struct {
	// 失败错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 失败错误信息

	ErrorMsg *string `json:"error_msg,omitempty"`
	// 失败资源信息

	TestcaseId *string `json:"testcase_id,omitempty"`
}

func (o ErrorCaseInfoBean) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ErrorCaseInfoBean struct{}"
	}

	return strings.Join([]string{"ErrorCaseInfoBean", string(data)}, " ")
}
