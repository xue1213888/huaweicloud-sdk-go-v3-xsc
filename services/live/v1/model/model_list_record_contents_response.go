package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListRecordContentsResponse struct {
	// 查询结果的总元素数量

	Total *int32 `json:"total,omitempty"`
	// 录制内容数组

	RecordContents *[]RecordContentInfoV2 `json:"record_contents,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRecordContentsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListRecordContentsResponse struct{}"
	}

	return strings.Join([]string{"ListRecordContentsResponse", string(data)}, " ")
}