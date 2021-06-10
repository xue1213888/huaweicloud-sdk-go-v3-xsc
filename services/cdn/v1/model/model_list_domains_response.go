package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListDomainsResponse struct {
	// 总条数。

	Total *int32 `json:"total,omitempty"`
	// 域名信息

	Domains        *[]Domains `json:"domains,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListDomainsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListDomainsResponse struct{}"
	}

	return strings.Join([]string{"ListDomainsResponse", string(data)}, " ")
}
