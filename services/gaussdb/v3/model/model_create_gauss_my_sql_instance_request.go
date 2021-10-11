package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateGaussMySqlInstanceRequest struct {
	// 语言。

	XLanguage *string `json:"X-Language,omitempty"`

	Body *MysqlInstanceRequest `json:"body,omitempty"`
}

func (o CreateGaussMySqlInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateGaussMySqlInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateGaussMySqlInstanceRequest", string(data)}, " ")
}