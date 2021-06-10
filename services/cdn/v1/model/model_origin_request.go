package model

import (
	"encoding/json"

	"strings"
)

type OriginRequest struct {
	Origin *ResourceBody `json:"origin"`
}

func (o OriginRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "OriginRequest struct{}"
	}

	return strings.Join([]string{"OriginRequest", string(data)}, " ")
}
