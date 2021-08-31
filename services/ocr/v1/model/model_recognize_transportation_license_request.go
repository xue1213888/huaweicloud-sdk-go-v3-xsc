package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeTransportationLicenseRequest struct {
	Body *TransportationLicenseRequestBody `json:"body,omitempty"`
}

func (o RecognizeTransportationLicenseRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeTransportationLicenseRequest struct{}"
	}

	return strings.Join([]string{"RecognizeTransportationLicenseRequest", string(data)}, " ")
}