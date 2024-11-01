package handler

import (
	"encoding/json"
	"net/http"
)

// response is a common response structure
type response struct {
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

// writeJSON helper function
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// writeError helper function
func writeError(w http.ResponseWriter, status int, err string) {
	writeJSON(w, status, response{Error: err})
}
