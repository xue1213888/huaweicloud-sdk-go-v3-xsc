/*
 * EVS
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
type DeleteVolumeResponse struct {
	// 正常返回时返回的任务ID。
	JobId *string `json:"job_id,omitempty"`
}

func (o DeleteVolumeResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteVolumeResponse", string(data)}, " ")
}
