package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListHostRouteRequest struct {
	// 域名Id

	InstanceId string `json:"instance_id"`
}

func (o ListHostRouteRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListHostRouteRequest struct{}"
	}

	return strings.Join([]string{"ListHostRouteRequest", string(data)}, " ")
}