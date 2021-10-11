package model

import (
	"encoding/json"

	"strings"
)

type AclApiBindingCreate struct {
	// ACL策略编号

	AclId *string `json:"acl_id,omitempty"`
	// API发布记录编号

	PublishIds *[]string `json:"publish_ids,omitempty"`
}

func (o AclApiBindingCreate) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AclApiBindingCreate struct{}"
	}

	return strings.Join([]string{"AclApiBindingCreate", string(data)}, " ")
}