package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateJobRequest struct {
	// 集群ID

	ClusterId string `json:"cluster_id"`
	// 作业名称。

	JobName string `json:"job_name"`

	Body *CdmUpdateJobJsonReq `json:"body,omitempty"`
}

func (o UpdateJobRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateJobRequest struct{}"
	}

	return strings.Join([]string{"UpdateJobRequest", string(data)}, " ")
}