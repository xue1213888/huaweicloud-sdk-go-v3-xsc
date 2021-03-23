package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeletePublicZoneResponse struct {
	// Zone的ID

	Id *string `json:"id,omitempty"`
	// zone名称

	Name *string `json:"name,omitempty"`
	// 对zone的描述信息

	Description *string `json:"description,omitempty"`
	// 管理该zone的管理员邮箱

	Email *string `json:"email,omitempty"`
	// zone类型，公网（public）或者内网（private）

	ZoneType *string `json:"zone_type,omitempty"`
	// 该zone下SOA记录中的ttl值

	Ttl *int32 `json:"ttl,omitempty"`
	// 该zone下SOA记录中用于标识zone文件变更的序列值，用于主从节点同步

	Serial *int32 `json:"serial,omitempty"`
	// 该zone下的recordset个数

	Status *string `json:"status,omitempty"`
	// 该zone下的recordset个数

	RecordNum *int32 `json:"record_num,omitempty"`
	// 托管该zone的pool，由系统分配

	PoolId *string `json:"pool_id,omitempty"`
	// zone所属的项目ID

	ProjectId *string `json:"project_id,omitempty"`
	// 创建时间

	CreatedAt *string `json:"created_at,omitempty"`
	// 更新时间

	UpdatedAt *string `json:"updated_at,omitempty"`
	// 主从模式中，从DNS服务器用以获取DNS信息

	Masters *string `json:"masters,omitempty"`
	// 指向当前资源或者其他资源的链接。当查询需要分页时，需要包含一个next链接指向下一页

	Links          *[]PageLink `json:"links,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o DeletePublicZoneResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeletePublicZoneResponse struct{}"
	}

	return strings.Join([]string{"DeletePublicZoneResponse", string(data)}, " ")
}