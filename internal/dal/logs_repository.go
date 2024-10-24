package dal

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func SaveJSONLog(r *http.Request, level slog.Level, fields []any, msg string) {
	// Open or create the log file with read/write and append permissions
	file, err := os.OpenFile("data/logs.json", os.O_RDWR|os.O_CREATE, 0o666)
	if err != nil {
		slog.Error("Failed to open log file", "error", err)
		return
	}
	defer file.Close()

	// Get the file's size
	fileInfo, err := file.Stat()
	if err != nil {
		slog.Error("Failed to get file info", "error", err)
		return
	}

	// If the file is empty, start it as a JSON array
	if fileInfo.Size() == 0 {
		_, err := file.Write([]byte("[\n"))
		if err != nil {
			slog.Error("Failed to write opening bracket to log file", "error", err)
			return
		}
	} else {
		content, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatalf("Failed to read file: %v", err)
		}

		// Convert content to a string and remove the last two characters
		modifiedContent := string(content)
		if len(modifiedContent) >= 2 {
			modifiedContent = modifiedContent[:len(modifiedContent)-2]
		} else {
			log.Println("File content is too short to truncate")
			return
		}

		// Move the file pointer to the beginning of the file
		_, err = file.Seek(0, 0)
		if err != nil {
			log.Fatalf("Failed to seek to the beginning of the file: %v", err)
		}

		// Write the modified content back to the file
		_, err = file.WriteString(modifiedContent)
		if err != nil {
			log.Fatalf("Failed to write to file: %v", err)
		}

		// Truncate the file to the new size
		err = file.Truncate(int64(len(modifiedContent)))
		if err != nil {
			log.Printf("Failed to truncate file: %v", err)
			return
		}

		// Write a comma to separate log entries
		_, err = file.Write([]byte(",\n"))
		if err != nil {
			slog.Error("Failed to write comma to log file", "error", err)
			return
		}
	}

	// Capture log output in a buffer to manipulate the string
	var buf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buf, nil))

	// Create a context
	ctx := context.Background()

	// Log the provided message with the given fields and level
	logger.Log(ctx, level, msg, fields...)

	// Convert buffer content to string, trimming any trailing newline
	logContent := bytes.TrimSuffix(buf.Bytes(), []byte("\n"))
	logContent = append([]byte("  "), logContent...)
	// Write the trimmed log content to the file
	_, err = file.Write(logContent)
	if err != nil {
		slog.Error("Failed to write log content to file", "error", err)
		return
	}

	// Append the closing ']' back to the file to maintain proper JSON array format
	_, err = file.Write([]byte("\n]"))
	if err != nil {
		slog.Error("Failed to write closing bracket to log file", "error", err)
	}
}
