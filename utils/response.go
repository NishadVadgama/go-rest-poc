package utils

import (
	"encoding/json"
	"net/http"
)

// Format response as JSON
func JSONResponse(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
