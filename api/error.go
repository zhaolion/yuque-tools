package api

import (
	"encoding/json"
)

type YuqueError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e *YuqueError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
