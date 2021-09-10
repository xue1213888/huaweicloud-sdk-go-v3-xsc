package model

import (
	"encoding/json"

	"strings"
)

type CdmDeleteClusterReq struct {
	// 日志备份数，填写为默认填0即可。

	KeepLastManualBackup int32 `json:"keep_last_manual_backup"`
}

func (o CdmDeleteClusterReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CdmDeleteClusterReq struct{}"
	}

	return strings.Join([]string{"CdmDeleteClusterReq", string(data)}, " ")
}
