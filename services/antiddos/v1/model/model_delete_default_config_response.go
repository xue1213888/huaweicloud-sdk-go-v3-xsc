package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteDefaultConfigResponse struct {
	// 内部错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 内部错误描述

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDefaultConfigResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDefaultConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteDefaultConfigResponse", string(data)}, " ")
}
