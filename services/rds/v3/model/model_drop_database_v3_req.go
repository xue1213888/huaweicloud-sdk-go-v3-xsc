package model

import (
	"encoding/json"

	"strings"
)

type DropDatabaseV3Req struct {
	// 是否强制删除数据库，默认是false。

	IsForceDelete *bool `json:"is_force_delete,omitempty"`
}

func (o DropDatabaseV3Req) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DropDatabaseV3Req struct{}"
	}

	return strings.Join([]string{"DropDatabaseV3Req", string(data)}, " ")
}
