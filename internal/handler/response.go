package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ab-dauletkhan/hot-coffee/internal/service"
)

// ErrorJSONResponse sends an error response with a given status code and message.
func ErrorJSONResponse(w http.ResponseWriter, r *http.Request, code int, msg string) {
	response := map[string]string{"error": msg}
	writeResponse(w, r, slog.LevelError, code, msg, response)
}

// SuccessJSONResponse sends a success response with a given status code and message.
func SuccessJSONResponse(w http.ResponseWriter, r *http.Request, code int, msg string) {
	response := map[string]string{"success": msg}
	writeResponse(w, r, slog.LevelInfo, code, msg, response)
}

// CustomJSONResponse sends a custom response with a key-value pair or a single value.
func CustomJSONResponse(w http.ResponseWriter, r *http.Request, code int, msg string, key, value interface{}, level slog.Level) {
	var response interface{}

	if key != nil {
		response = map[interface{}]interface{}{key: value}
	} else {
		response = value
	}

	writeResponse(w, r, level, code, msg, response)
}

// writeResponse logs the response and writes it to the http.ResponseWriter.
func writeResponse(w http.ResponseWriter, r *http.Request, level slog.Level, code int, msg string, response interface{}) {
	service.CreateLog(r, level, code, msg)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Debug(fmt.Sprintf("error encoding JSON response: %v", err))
		http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
	}
}
