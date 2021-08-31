package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeBusinessCardResponse struct {
	Result         *BusinessCardResult `json:"result,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o RecognizeBusinessCardResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeBusinessCardResponse struct{}"
	}

	return strings.Join([]string{"RecognizeBusinessCardResponse", string(data)}, " ")
}