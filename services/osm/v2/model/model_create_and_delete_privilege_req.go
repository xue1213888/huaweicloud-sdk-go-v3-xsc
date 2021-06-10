package model

import (
	"encoding/json"

	"strings"
)

type CreateAndDeletePrivilegeReq struct {
	// 执行的操作(create|delete)

	Operation string `json:"operation"`
	// 权限标识

	Privilege *string `json:"privilege,omitempty"`
}

func (o CreateAndDeletePrivilegeReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAndDeletePrivilegeReq struct{}"
	}

	return strings.Join([]string{"CreateAndDeletePrivilegeReq", string(data)}, " ")
}
