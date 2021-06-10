package model

import (
	"encoding/json"

	"strings"
)

// 用户的角色信息
type UserRoleVo struct {
	// 角色

	Role string `json:"role"`
	// 说明

	Comment *string `json:"comment,omitempty"`
	// 是否支持迁移。

	IsTransfer string `json:"is_transfer"`
	// 权限

	Privileges string `json:"privileges"`
	// 继承角色列表

	InheritsRoles *[]string `json:"inherits_roles,omitempty"`
	// 是否选择。

	Selected *string `json:"selected,omitempty"`
}

func (o UserRoleVo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UserRoleVo struct{}"
	}

	return strings.Join([]string{"UserRoleVo", string(data)}, " ")
}
