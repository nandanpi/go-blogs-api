package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WritePlainText(w http.ResponseWriter, status int, v string) error {
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
