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

// 更新镜像成员状态请求体
type BatchUpdateMembersRequestBody struct {
	// 镜像ID列表。
	Images []string `json:"images"`
	// 项目ID。
	ProjectId string `json:"project_id"`
	// 镜像成员的状态。 取值如下： accepted：表示接受共享镜像。接受后，该镜像在用户镜像列表中可见，用户可以使用该镜像创建云服务器。 rejected：表示拒绝共享镜像。拒绝后，该镜像在用户镜像列表中不可见，但是，用户仍然可以使用该镜像创建云服务器。
	Status BatchUpdateMembersRequestBodyStatus `json:"status"`
	// 存储库ID。 如果是通过CBR创建的整机镜像，则在接受该共享镜像时，为必选参数，需传入该值。
	VaultId *string `json:"vault_id,omitempty"`
}

func (o BatchUpdateMembersRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchUpdateMembersRequestBody", string(data)}, " ")
}

type BatchUpdateMembersRequestBodyStatus struct {
	value string
}

type BatchUpdateMembersRequestBodyStatusEnum struct {
	ACCEPTED BatchUpdateMembersRequestBodyStatus
	REJECTED BatchUpdateMembersRequestBodyStatus
}

func GetBatchUpdateMembersRequestBodyStatusEnum() BatchUpdateMembersRequestBodyStatusEnum {
	return BatchUpdateMembersRequestBodyStatusEnum{
		ACCEPTED: BatchUpdateMembersRequestBodyStatus{
			value: "accepted",
		},
		REJECTED: BatchUpdateMembersRequestBodyStatus{
			value: "rejected",
		},
	}
}

func (c BatchUpdateMembersRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BatchUpdateMembersRequestBodyStatus) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
