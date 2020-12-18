package exception

import "encoding/json"

type Exception struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

func GetJsonBody(statusCode int, message string) string {

	exception := Exception{}
	exception.Status = statusCode
	exception.Message = message

	data , _ := json.Marshal(exception)

	return string(data)
}
