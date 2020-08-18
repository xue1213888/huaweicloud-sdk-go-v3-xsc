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
type ImportImageQuickResponse struct {
	// 异步任务ID。
	JobId *string `json:"job_id,omitempty"`
}

func (o ImportImageQuickResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ImportImageQuickResponse", string(data)}, " ")
}
