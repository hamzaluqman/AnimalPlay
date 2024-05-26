package handler_response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// type HandlerResponse struct {
// 	Data         interface{}
// 	ResponseCode int
// 	ErrorMessage string
// }

func WriteErrorResponse(writer http.ResponseWriter, errCode int, errMsg string) {
	response := make(map[string]string)
	response["message"] = errMsg

	writer.WriteHeader(errCode)
	jsonResponse, err := json.Marshal(response)

	if err != nil {
		fmt.Println("Error marshaling response")
		return
	}

	_, err = writer.Write(jsonResponse)

	if err != nil {
		// Add logging
		fmt.Println("Error writing response")
		return
	}
}

func WriteSuccessResponse(writer http.ResponseWriter, data any, errCode int) {
	writer.WriteHeader(errCode)
	jsonResponse, err := json.Marshal(data)

	if err != nil {
		fmt.Println("Error marshaling response")
		return
	}

	_, err = writer.Write(jsonResponse)

	if err != nil {
		// Add logging
		fmt.Println("Error writing response")
		return
	}
}
