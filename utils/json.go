package utils

import (
	"encoding/json"
	"net/http"
)

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
