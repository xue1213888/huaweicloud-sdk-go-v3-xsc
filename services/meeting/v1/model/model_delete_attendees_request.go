package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteAttendeesRequest struct {

	// 会议ID
	ConferenceID string `json:"conferenceID"`

	// 会控授权令牌，通过获取会控token接口获得。
	XConferenceAuthorization string `json:"X-Conference-Authorization"`

	Body *RestBulkDelAttendReqBody `json:"body,omitempty"`
}

func (o DeleteAttendeesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAttendeesRequest struct{}"
	}

	return strings.Join([]string{"DeleteAttendeesRequest", string(data)}, " ")
}
