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

// Request Object
type ListPoliciesRequest struct {
	OperationType ListPoliciesRequestOperationType `json:"operation_type,omitempty"`
	VaultId       *string                          `json:"vault_id,omitempty"`
}

func (o ListPoliciesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListPoliciesRequest", string(data)}, " ")
}

type ListPoliciesRequestOperationType struct {
	value string
}

type ListPoliciesRequestOperationTypeEnum struct {
	BACKUP      ListPoliciesRequestOperationType
	REPLICATION ListPoliciesRequestOperationType
}

func GetListPoliciesRequestOperationTypeEnum() ListPoliciesRequestOperationTypeEnum {
	return ListPoliciesRequestOperationTypeEnum{
		BACKUP: ListPoliciesRequestOperationType{
			value: "backup",
		},
		REPLICATION: ListPoliciesRequestOperationType{
			value: "replication",
		},
	}
}

func (c ListPoliciesRequestOperationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListPoliciesRequestOperationType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
