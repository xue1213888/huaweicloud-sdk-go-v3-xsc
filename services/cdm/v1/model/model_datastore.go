package model

import (
	"encoding/json"

	"strings"
)

// cdm信息
type Datastore struct {
	// 类型，一般为cdm。

	Type *string `json:"type,omitempty"`
	// 集群版本。

	Version *string `json:"version,omitempty"`
}

func (o Datastore) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Datastore struct{}"
	}

	return strings.Join([]string{"Datastore", string(data)}, " ")
}
