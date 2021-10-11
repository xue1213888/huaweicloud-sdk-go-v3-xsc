package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateJobResponse struct {
	// 校验结果：如果修改失败，返回失败原因。如果修改成功，返回空列表

	ValidationResult *[]JobValidationResult `json:"validation-result,omitempty"`
	HttpStatusCode   int                    `json:"-"`
}

func (o UpdateJobResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateJobResponse struct{}"
	}

	return strings.Join([]string{"UpdateJobResponse", string(data)}, " ")
}