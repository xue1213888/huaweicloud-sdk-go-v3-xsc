package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListNewInstantMessagesResponse struct {
	// 状态列表

	Imstatus *[]ImStatusV2 `json:"imstatus,omitempty"`
	// 留言内容列表

	Immsg *[]UserInstantIncidentMsgV2 `json:"immsg,omitempty"`
	// 上次查询留言时间

	LastMessageTimeId *string `json:"last_message_time_id,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o ListNewInstantMessagesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListNewInstantMessagesResponse struct{}"
	}

	return strings.Join([]string{"ListNewInstantMessagesResponse", string(data)}, " ")
}
