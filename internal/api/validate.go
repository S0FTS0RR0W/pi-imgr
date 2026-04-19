package api

import (
	"encoding/json"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]any{
		"status": "ok",
	})
}

func ValidateImageHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]any{
		"valid": true,
		"type":  "raspios",
	})
}
