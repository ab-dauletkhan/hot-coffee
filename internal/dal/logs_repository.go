package dal

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/ab-dauletkhan/hot-coffee/models"
)

func SaveJSONLog(r *http.Request, level slog.Level, msg, path string, code int) {
	log := models.Log{
		Time:       time.Now(),
		Level:      level,
		Msg:        msg,
		Method:     r.Method,
		Proto:      r.Proto,
		Path:       path,
		Status:     code,
		User_agent: r.UserAgent(),
	}

	file, err := os.ReadFile("data/logs.json")
	if err != nil {
		slog.Debug("couldn't read from logs.json")
		return
	}

	allLogs := []models.Log{}
	if err := json.Unmarshal(file, &allLogs); err != nil && err != errors.New("unexpected end of JSON input") {
		slog.Debug(fmt.Sprintf("error: %v", err))
	}

	allLogs = append(allLogs, log)

	data, err := json.MarshalIndent(allLogs, "  ", "  ")
	if err != nil {
		slog.Debug("couldn't marshall logs")
		return
	}

	if err := os.WriteFile("data/logs.json", data, 0o666); err != nil {
		slog.Debug("couldn't write to logs.json")
	}
}
