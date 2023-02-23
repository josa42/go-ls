package utils

import (
	"encoding/json"
)

func Unmarkshal[A interface{}](input interface{}) (A, error) {
	var params A
	jsonData, _ := json.Marshal(input)
	err := json.Unmarshal(jsonData, &params)

	return params, err
}
