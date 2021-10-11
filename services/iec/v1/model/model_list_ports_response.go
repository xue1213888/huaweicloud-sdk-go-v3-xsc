package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListPortsResponse struct {
	// 端口对象。

	Ports *[]Port `json:"ports,omitempty"`
	// 端口数目。

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPortsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPortsResponse struct{}"
	}

	return strings.Join([]string{"ListPortsResponse", string(data)}, " ")
}