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
type ListVolumesByTagsResponse struct {
	// 符合查询条件的云硬盘资源个数
	TotalCount *string `json:"total_count,omitempty"`
	// 符合查询条件的资源列表
	Resources []Resource `json:"resources,omitempty"`
}

func (o ListVolumesByTagsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListVolumesByTagsResponse", string(data)}, " ")
}
