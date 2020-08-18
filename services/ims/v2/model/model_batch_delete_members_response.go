/*
 * IMS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Response Object
type BatchDeleteMembersResponse struct {
	// 异步任务ID。
	JobId *string `json:"job_id,omitempty"`
}

func (o BatchDeleteMembersResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchDeleteMembersResponse", string(data)}, " ")
}
