package handler_utils

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/ab-dauletkhan/hot-coffee/internal/service"
)

func ErrorJSONResponse(w http.ResponseWriter, r *http.Request, code int, msg string) {
	response := make(map[string]string)
	response["error"] = msg

	writeHeader(w, r, code, msg, response)
}

func SuccessJSONResponse(w http.ResponseWriter, r *http.Request, code int, msg string) {
	response := make(map[string]string)
	response["success"] = msg

	writeHeader(w, r, code, msg, response)
}

func CustomJSONREsponse(w http.ResponseWriter, r *http.Request, code int, msg string, key, value interface{}) {
	response := make(map[interface{}]interface{})
	response[key] = value
	writeHeader(w, r, code, msg, response)
}

func writeHeader(w http.ResponseWriter, r *http.Request, code int, msg string, response interface{}) {
	service.CreateLog(
		r,
		slog.LevelError,
		code,
		msg,
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(response)
}
