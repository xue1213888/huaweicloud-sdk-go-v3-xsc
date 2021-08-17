package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowAlertConfigResponse struct {
	// 告警群组的唯一标识

	TopicUrn *string `json:"topic_urn,omitempty"`
	// 告警群组描述

	DisplayName *string `json:"display_name,omitempty"`

	WarnConfig     *AlertConfigRespWarnConfig `json:"warn_config,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowAlertConfigResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowAlertConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowAlertConfigResponse", string(data)}, " ")
}
