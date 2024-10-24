package dal

import (
	"bytes"
	"context"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func SaveJSONLog(r *http.Request, level slog.Level, fields []any, msg string) {
	file, err := os.OpenFile("data/logs.json", os.O_RDWR|os.O_CREATE, 0o666)
	if err != nil {
		slog.Error("Failed to open log file", "error", err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		slog.Error("Failed to get file info", "error", err)
		return
	}

	if fileInfo.Size() == 0 {
		_, err := file.Write([]byte("[\n"))
		if err != nil {
			slog.Error("Failed to write opening bracket to log file", "error", err)
			return
		}
	} else {
		content, err := io.ReadAll(file)
		if err != nil {
			log.Fatalf("Failed to read file: %v", err)
		}

		modifiedContent := string(content)
		if len(modifiedContent) >= 2 {
			modifiedContent = modifiedContent[:len(modifiedContent)-2]
		} else {
			log.Println("File content is too short to truncate")
			return
		}

		_, err = file.Seek(0, 0)
		if err != nil {
			log.Fatalf("Failed to seek to the beginning of the file: %v", err)
		}

		_, err = file.WriteString(modifiedContent)
		if err != nil {
			log.Fatalf("Failed to write to file: %v", err)
		}

		err = file.Truncate(int64(len(modifiedContent)))
		if err != nil {
			log.Printf("Failed to truncate file: %v", err)
			return
		}

		_, err = file.Write([]byte(",\n"))
		if err != nil {
			slog.Error("Failed to write comma to log file", "error", err)
			return
		}
	}

	var buf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buf, nil))

	ctx := context.Background()

	logger.Log(ctx, level, msg, fields...)

	logContent := bytes.TrimSuffix(buf.Bytes(), []byte("\n"))
	logContent = append([]byte("  "), logContent...)
	_, err = file.Write(logContent)
	if err != nil {
		slog.Error("Failed to write log content to file", "error", err)
		return
	}

	_, err = file.Write([]byte("\n]"))
	if err != nil {
		slog.Error("Failed to write closing bracket to log file", "error", err)
	}
}
