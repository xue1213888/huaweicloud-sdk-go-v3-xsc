/*
 * IAM
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

//
type AssumeroleSessionuser struct {
	// 委托方对应的企业用户名。用户名需满足如下规则：长度5~32，只能包含大写字母、小写字母、数字（0-9）、特殊字符（\"-\"与\"_\"）且只能以字母开头。
	Name *string `json:"name,omitempty"`
}

func (o AssumeroleSessionuser) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AssumeroleSessionuser", string(data)}, " ")
}
