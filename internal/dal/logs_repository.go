package dal

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

const (
	filePerm = 0o644
)

func SaveJSONLog(r *http.Request, level slog.Level, fields []any, msg string) {
	file, err := os.OpenFile("data/logs.txt", os.O_WRONLY|os.O_APPEND, filePerm)
	if err != nil {
		slog.Debug(fmt.Sprintf("error opening logs.txt: %v", err))
	}

	logger := slog.New(slog.NewJSONHandler(file, nil))
	logger.Log(nil, level, msg, fields...)
}
