/*
 * Cbr
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// 备份共享更新信息
type UpdateMember struct {
	// 备份共享状态
	Status UpdateMemberStatus `json:"status"`
	// 共享的备份将存入的存储库，仅支持uuid 更新member状态的时候，如果是接受，必须传入vault_id，如果是拒绝，则无需
	VaultId *string `json:"vault_id,omitempty"`
}

func (o UpdateMember) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateMember", string(data)}, " ")
}

type UpdateMemberStatus struct {
	value string
}

type UpdateMemberStatusEnum struct {
	ACCEPTED UpdateMemberStatus
	PENDING  UpdateMemberStatus
	REJECTED UpdateMemberStatus
}

func GetUpdateMemberStatusEnum() UpdateMemberStatusEnum {
	return UpdateMemberStatusEnum{
		ACCEPTED: UpdateMemberStatus{
			value: "accepted",
		},
		PENDING: UpdateMemberStatus{
			value: "pending",
		},
		REJECTED: UpdateMemberStatus{
			value: "rejected",
		},
	}
}

func (c UpdateMemberStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateMemberStatus) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
