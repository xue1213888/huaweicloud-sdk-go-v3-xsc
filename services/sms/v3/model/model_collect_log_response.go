package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CollectLogResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CollectLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectLogResponse struct{}"
	}

	return strings.Join([]string{"CollectLogResponse", string(data)}, " ")
}
