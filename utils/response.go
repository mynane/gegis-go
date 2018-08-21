package utils

import (
	"net/http"
	"encoding/json"
)

func NotAuth(res string) []byte {
	result := make(map[string]interface{})
	result["code"] = http.StatusUnauthorized
	result["data"] = res
	str, _ := json.Marshal(result)
	return str
}
