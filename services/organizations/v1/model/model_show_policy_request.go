package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowPolicyRequest struct {

	// 根、组织单元或帐号的唯一标识符（ID）。
	PolicyId string `json:"policy_id"`

	// 选择接口返回的信息的语言，默认为\"zh-cn\"中文，zh-cn中文，en-us英文
	XLanguage *ShowPolicyRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowPolicyRequest", string(data)}, " ")
}

type ShowPolicyRequestXLanguage struct {
	value string
}

type ShowPolicyRequestXLanguageEnum struct {
	ZH_CN ShowPolicyRequestXLanguage
	EN_US ShowPolicyRequestXLanguage
}

func GetShowPolicyRequestXLanguageEnum() ShowPolicyRequestXLanguageEnum {
	return ShowPolicyRequestXLanguageEnum{
		ZH_CN: ShowPolicyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowPolicyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowPolicyRequestXLanguage) Value() string {
	return c.value
}

func (c ShowPolicyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPolicyRequestXLanguage) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}