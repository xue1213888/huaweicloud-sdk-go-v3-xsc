package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateMigprojectResponse struct {
	// 创建迁移项目返回的新添加的迁移项目的id

	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMigprojectResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMigprojectResponse struct{}"
	}

	return strings.Join([]string{"CreateMigprojectResponse", string(data)}, " ")
}
