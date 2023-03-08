package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Msg struct {
	Message string `json:"message"`
}

func TestRun(t *testing.T) {
	str := []byte("{\n    \"error\": {\n        \"message\": \"Incorrect API key provided: sk-3XSO2***************************************DBtc. You can find your API key at https://platform.openai.com/account/api-keys.\",\n        \"type\": \"invalid_request_error\",\n        \"param\": null,\n        \"code\": \"invalid_api_key\"\n    }\n}")

	respMap := make(map[string]any)
	_ = json.Unmarshal(str, &respMap)
	fmt.Println(respMap)
	errorStr := respMap["error"]
	m := errorStr.(Msg)
	fmt.Println(m)
}
