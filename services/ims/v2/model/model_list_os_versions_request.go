/*
 * IMS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Request Object
type ListOsVersionsRequest struct {
	Tag *string `json:"tag,omitempty"`
}

func (o ListOsVersionsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListOsVersionsRequest", string(data)}, " ")
}
