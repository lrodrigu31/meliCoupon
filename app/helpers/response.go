package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SenData : func used to send success response to services
func SenData(rw http.ResponseWriter, data interface{}, status int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)

	output, _ := json.Marshal(&data)
	fmt.Fprintln(rw, string(output))
}

// SendError func used to send error response to services
func SendError(rw http.ResponseWriter, status int, message string) {
	rw.WriteHeader(status)
	fmt.Fprintln(rw, message)
}
