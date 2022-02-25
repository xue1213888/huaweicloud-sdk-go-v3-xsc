package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePwdRequest struct {
	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用uuId，若不携带，则后台自动生成

	XRequestID *string `json:"X-Request-ID,omitempty"`
	// 语言参数，默认为中文zh-CN, 英文为en-US

	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	Body *ModPwdReqDto `json:"body,omitempty"`
}

func (o UpdatePwdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePwdRequest struct{}"
	}

	return strings.Join([]string{"UpdatePwdRequest", string(data)}, " ")
}
