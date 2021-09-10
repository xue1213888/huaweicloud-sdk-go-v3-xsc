package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type StartJobRequest struct {
	// 集群ID

	ClusterId string `json:"cluster_id"`
	// 作业名称

	JobName string `json:"job_name"`
}

func (o StartJobRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartJobRequest struct{}"
	}

	return strings.Join([]string{"StartJobRequest", string(data)}, " ")
}
