package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVideoCoverAnalysisTaskRequestBody struct {
	Input *VideoCoverAnalysisTaskInput `json:"input"`

	Output *TaskOutput `json:"output"`

	Callback *TaskCallback `json:"callback,omitempty"`

	Config *VideoCoverAnalysisConfig `json:"config,omitempty"`
}

func (o CreateVideoCoverAnalysisTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoCoverAnalysisTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVideoCoverAnalysisTaskRequestBody", string(data)}, " ")
}