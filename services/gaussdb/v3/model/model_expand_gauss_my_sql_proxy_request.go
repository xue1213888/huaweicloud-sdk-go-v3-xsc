package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ExpandGaussMySqlProxyRequest struct {
	// 语言。

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID，严格匹配UUID规则。

	InstanceId string `json:"instance_id"`

	Body *EnlargeProxyRequest `json:"body,omitempty"`
}

func (o ExpandGaussMySqlProxyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExpandGaussMySqlProxyRequest struct{}"
	}

	return strings.Join([]string{"ExpandGaussMySqlProxyRequest", string(data)}, " ")
}