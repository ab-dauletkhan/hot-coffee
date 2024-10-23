package handler_utils

import (
	"encoding/json"
	"net/http"
)

func ErrorResponseJSON(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := make(map[string]string)
	response["error"] = msg

	json.NewEncoder(w).Encode(response)
}
