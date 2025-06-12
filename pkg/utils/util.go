package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
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

// ValidateAndConvertID validates and converts the ID parameter from a string to an integer.
// Returns the integer ID and an error if the ID is invalid.
func ValidateAndConvertID(idParam string) (int, error) {
    if idParam == "" {
        return 0, errors.New("task ID is required")
    }

    // Convert the ID to an integer
    id, err := strconv.Atoi(idParam)
    if err != nil {
        return 0, errors.New("invalid task ID")
    }

    return id, nil
}