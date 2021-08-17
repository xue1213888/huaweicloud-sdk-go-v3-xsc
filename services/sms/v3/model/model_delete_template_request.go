package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteTemplateRequest struct {
	// 需要删除的模板的ID

	Id string `json:"id"`
}

func (o DeleteTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteTemplateRequest", string(data)}, " ")
}
