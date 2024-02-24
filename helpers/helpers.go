package helpers

import (
	"encoding/json"
	"net/http"
)

// RespondWithError sends an HTTP error response with the given status code and message
func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	http.Error(w, message, statusCode)
}

// RespondWithJSON sends an HTTP response with the given status code and JSON body
func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
