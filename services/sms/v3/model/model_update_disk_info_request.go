package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateDiskInfoRequest struct {
	// 源端服务器id

	SourceId string `json:"source_id"`

	Body *PutDiskInfoReq `json:"body,omitempty"`
}

func (o UpdateDiskInfoRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDiskInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateDiskInfoRequest", string(data)}, " ")
}
