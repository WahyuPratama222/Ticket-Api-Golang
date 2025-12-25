
package utils

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse is the standard error response structure
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// SuccessResponse is the standard success response structure
type SuccessResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

// WriteJSON menulis response JSON dengan format rapi
func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	return err
}

// WriteErrorJSON menulis error response dengan format konsisten
func WriteErrorJSON(w http.ResponseWriter, status int, message string) error {
	return WriteJSON(w, status, ErrorResponse{
		Error:   http.StatusText(status),
		Message: message,
	})
}