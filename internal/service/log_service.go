package service

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/ab-dauletkhan/hot-coffee/internal/dal"
)

func CreateLog(r *http.Request, level slog.Level, code int, msg string) {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))
	fields := logCommonFields(r, code)
	dal.SaveJSONLog(r, level, fields, msg)

	switch level {
	case slog.LevelInfo:
		log.Info(msg, fields...)
	case slog.LevelError:
		log.Error(msg, fields...)
	case slog.LevelDebug:
		log.Debug(msg, fields...)
	case slog.LevelWarn:
		log.Warn(msg, fields...)
	default:
		log.Warn(msg, fields...)
	}
}

func logCommonFields(r *http.Request, code int) []any {
	path := getPath(r)

	return []any{
		slog.String("method", r.Method),
		slog.String("proto", r.Proto),
		slog.String("path", path),
		slog.Int("status", code),
		slog.String("user_agent", r.UserAgent()),
	}
}

func getPath(r *http.Request) string {
	path := r.URL.Path
	if len(r.URL.RawQuery) != 0 {
		path += "?" + r.URL.RawQuery
	}

	return path
}
