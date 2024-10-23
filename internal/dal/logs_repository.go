package dal

import (
	"log/slog"
	"net/http"
)

func SaveJSONLog(r *http.Request, level slog.Level, fields []any, msg string) {
	// TODO: implement saving json logs
}
