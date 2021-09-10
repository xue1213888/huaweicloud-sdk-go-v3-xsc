package model

import (
	"encoding/json"

	"strings"
)

// 请求消息的数据部分。
type IvsStandardByNameAndIdRequestBodyData struct {
	// 请求列表，用于支持批量调用。目前暂时只支持单个数据查询。

	ReqData *[]StandardReqDataByNameAndId `json:"req_data,omitempty"`
}

func (o IvsStandardByNameAndIdRequestBodyData) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "IvsStandardByNameAndIdRequestBodyData struct{}"
	}

	return strings.Join([]string{"IvsStandardByNameAndIdRequestBodyData", string(data)}, " ")
}
