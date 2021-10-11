package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowDocWatermarkResponse struct {
	// 暗水印内容，长度不超过32个字节

	Watermark      *string `json:"watermark,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDocWatermarkResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDocWatermarkResponse struct{}"
	}

	return strings.Join([]string{"ShowDocWatermarkResponse", string(data)}, " ")
}