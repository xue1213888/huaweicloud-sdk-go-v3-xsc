package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteClusterRequest struct {
	// 集群ID。获取方法，请参见[获取集群ID](mrs_02_9001.xml)。

	ClusterId string `json:"cluster_id"`
}

func (o DeleteClusterRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteClusterRequest struct{}"
	}

	return strings.Join([]string{"DeleteClusterRequest", string(data)}, " ")
}
