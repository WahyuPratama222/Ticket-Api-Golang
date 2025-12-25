package utils

import (
	"encoding/json"
	"net/http"
)

type StandardResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := StandardResponse{
		Success: status >= 200 && status < 300,
		Data:    data,
	}

	b, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	return err
}

func WriteSuccessJSON(w http.ResponseWriter, status int, message string, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := StandardResponse{
		Success: true,
		Message: message,
		Data:    data,
	}

	b, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	return err
}

func WriteErrorJSON(w http.ResponseWriter, status int, message string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := StandardResponse{
		Success: false,
		Error:   message,
	}

	b, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	return err
}
