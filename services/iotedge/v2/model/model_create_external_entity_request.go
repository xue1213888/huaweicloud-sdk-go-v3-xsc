package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateExternalEntityRequest struct {
	// 边缘节点ID

	EdgeNodeId string `json:"edge_node_id"`

	Body *CreateExternalEntityReqDto `json:"body,omitempty"`
}

func (o CreateExternalEntityRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateExternalEntityRequest struct{}"
	}

	return strings.Join([]string{"CreateExternalEntityRequest", string(data)}, " ")
}