package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteServerRequest struct {
	// 源端服务器在主机迁移服务中的id

	SourceId string `json:"source_id"`
}

func (o DeleteServerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteServerRequest struct{}"
	}

	return strings.Join([]string{"DeleteServerRequest", string(data)}, " ")
}
