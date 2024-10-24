package handler_utils

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/ab-dauletkhan/hot-coffee/internal/service"
)

func ErrorJSONResponse(w http.ResponseWriter, r *http.Request, code int, msg string) {
	service.CreateLog(
		r,
		slog.LevelError,
		code,
		msg,
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := make(map[string]string)
	response["error"] = msg

	json.NewEncoder(w).Encode(response)
}
