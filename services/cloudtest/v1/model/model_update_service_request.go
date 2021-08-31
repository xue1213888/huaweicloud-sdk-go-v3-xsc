package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateServiceRequest struct {
	// 注册服务唯一标识，该值由注册接口返回

	ServiceId int32 `json:"service_id"`

	Body *ServiceRequestBody `json:"body,omitempty"`
}

func (o UpdateServiceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateServiceRequest struct{}"
	}

	return strings.Join([]string{"UpdateServiceRequest", string(data)}, " ")
}