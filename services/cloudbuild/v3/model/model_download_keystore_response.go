package model

import (
	"encoding/json"

	"io"

	"strings"
)

// Response Object
type DownloadKeystoreResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o DownloadKeystoreResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o DownloadKeystoreResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DownloadKeystoreResponse struct{}"
	}

	return strings.Join([]string{"DownloadKeystoreResponse", string(data)}, " ")
}