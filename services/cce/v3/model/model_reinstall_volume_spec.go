package model

import (
	"encoding/json"

	"strings"
)

// 服务器重装云硬盘配置
type ReinstallVolumeSpec struct {
	// 用户自定义镜像ID

	ImageID *string `json:"imageID,omitempty"`
	// 用户主密钥ID。默认为空时，表示云硬盘不加密。

	CmkID *string `json:"cmkID,omitempty"`
}

func (o ReinstallVolumeSpec) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ReinstallVolumeSpec struct{}"
	}

	return strings.Join([]string{"ReinstallVolumeSpec", string(data)}, " ")
}
