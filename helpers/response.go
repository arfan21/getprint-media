package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/arfan21/getprint-media/models"
)

func Response(status string, message interface{}, data interface{}) models.Response {
	dataType := fmt.Sprintf("%T", message)

	if dataType == "error" {
		errorJSON, _ := json.Marshal(message)
		return models.Response{Status: status, Message: errorJSON, Data: nil}
	}

	return models.Response{Status: status, Message: message, Data: data}
}
