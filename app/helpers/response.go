package helpers

import (
	"encoding/json"
	"fmt"

	models2 "github.com/arfan21/getprint-media/app/models"
)

func Response(status string, message interface{}, data interface{}) models2.Response {
	dataType := fmt.Sprintf("%T", message)

	if dataType == "error" {
		errorJSON, _ := json.Marshal(message)
		return models2.Response{Status: status, Message: errorJSON, Data: nil}
	}

	return models2.Response{Status: status, Message: message, Data: data}
}
