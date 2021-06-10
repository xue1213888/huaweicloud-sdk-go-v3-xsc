package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListProblemTypesResponse struct {
	// 总数

	TotalCount *int32 `json:"total_count,omitempty"`
	// 问题类型列表

	IncidentBusinessTypeList *[]SimpleIncidentBusinessTypeV2 `json:"incident_business_type_list,omitempty"`
	HttpStatusCode           int                             `json:"-"`
}

func (o ListProblemTypesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProblemTypesResponse struct{}"
	}

	return strings.Join([]string{"ListProblemTypesResponse", string(data)}, " ")
}
