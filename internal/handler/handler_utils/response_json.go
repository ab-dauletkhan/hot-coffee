package handler_utils

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/ab-dauletkhan/hot-coffee/internal/service"
)

func JSONResponse(w http.ResponseWriter, r *http.Request, code int, msg, status string) {
	service.CreateLog(
		r,
		slog.LevelError,
		code,
		"invalid request payload",
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := make(map[string]string)
	response[status] = msg

	json.NewEncoder(w).Encode(response)
}
