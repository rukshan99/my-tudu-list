package utils

import (
	"encoding/json"
	"net/http"
)

// WriteJSONResponse marshals the given data into JSON and writes it to the HTTP response.
func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

// ParseJSONRequest unmarshals the JSON payload from the HTTP request into the given struct.
func ParseJSONRequest(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
