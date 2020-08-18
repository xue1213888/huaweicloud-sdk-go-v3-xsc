/*
 * EPS
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
type ListEpResponse struct {
	// 企业项目列表
	EnterpriseProjects []EpDetail `json:"enterprise_projects,omitempty"`
	// 企业项目总数
	TotalCount *int32 `json:"total_count,omitempty"`
}

func (o ListEpResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListEpResponse", string(data)}, " ")
}
