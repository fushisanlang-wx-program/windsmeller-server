package service

import (
	"encoding/json"
	"io"
	"net/http"
)

func init() {}
func transformation(response *http.Response) map[string]interface{} {
	var result map[string]interface{}
	body, err := io.ReadAll(response.Body)
	if err == nil {
		json.Unmarshal([]byte(string(body)), &result)
	}

	return result
}
