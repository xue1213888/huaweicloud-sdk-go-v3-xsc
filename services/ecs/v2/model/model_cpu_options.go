package model

import (
	"encoding/json"

	"strings"
)

// 自定义CPU选项。
type CpuOptions struct {
	// CPU超线程数， 决定CPU是否开启超线程

	HwcpuThreads *int32 `json:"hw:cpu_threads,omitempty"`
}

func (o CpuOptions) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CpuOptions struct{}"
	}

	return strings.Join([]string{"CpuOptions", string(data)}, " ")
}