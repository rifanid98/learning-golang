package handler

import (
	"encoding/json"
	error2 "golang11_dependency_injection/common/error"
	"net/http"
)

func GetRequestBody(request *http.Request, data interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(data)
	error2.PanicIfError(err)
}

func SendResponseBody(writer http.ResponseWriter, data interface{}, statusCode int) {
	writer.WriteHeader(statusCode)
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(data)
	error2.PanicIfError(err)
}
