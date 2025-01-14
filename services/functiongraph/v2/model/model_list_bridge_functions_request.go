package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBridgeFunctionsRequest Request Object
type ListBridgeFunctionsRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	// servicebridge类型
	Type *string `json:"type,omitempty"`
}

func (o ListBridgeFunctionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBridgeFunctionsRequest struct{}"
	}

	return strings.Join([]string{"ListBridgeFunctionsRequest", string(data)}, " ")
}
