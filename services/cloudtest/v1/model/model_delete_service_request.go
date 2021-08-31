package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteServiceRequest struct {
	// 注册服务唯一标识，该值由注册接口返回

	ServiceId int32 `json:"service_id"`
}

func (o DeleteServiceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteServiceRequest struct{}"
	}

	return strings.Join([]string{"DeleteServiceRequest", string(data)}, " ")
}