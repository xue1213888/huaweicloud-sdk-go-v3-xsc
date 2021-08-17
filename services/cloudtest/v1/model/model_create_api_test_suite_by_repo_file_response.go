package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateApiTestSuiteByRepoFileResponse struct {
	// 生成的测试套的id

	TestsuiteId *string `json:"testsuite_id,omitempty"`
	// 生成的测试用例id列表

	TestcaseIds    *[]string `json:"testcase_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateApiTestSuiteByRepoFileResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateApiTestSuiteByRepoFileResponse struct{}"
	}

	return strings.Join([]string{"CreateApiTestSuiteByRepoFileResponse", string(data)}, " ")
}
