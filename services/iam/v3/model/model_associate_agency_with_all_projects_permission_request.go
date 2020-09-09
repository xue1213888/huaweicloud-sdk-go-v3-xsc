/*
 * IAM
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
type AssociateAgencyWithAllProjectsPermissionRequest struct {
	AgencyId string `json:"agency_id"`
	DomainId string `json:"domain_id"`
	RoleId   string `json:"role_id"`
}

func (o AssociateAgencyWithAllProjectsPermissionRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AssociateAgencyWithAllProjectsPermissionRequest", string(data)}, " ")
}