package common

import (
	"encoding/json"
	"net/http"
)

func GetRequestBody(request *http.Request, data interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(data)
	PanicIfError(err)
}

func SendResponseBody(writer http.ResponseWriter, data interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(data)
	PanicIfError(err)
}
