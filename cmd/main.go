package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/ab-dauletkhan/hot-coffee/internal/router"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// TODO: Implement flag parsing (e.g. --help, --port)
	// Exit on invalid flags (e.g., invalid command-line arguments, failure to bind to a port)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.Routes(),
	}

	// env = 'local' | 'dev' | 'prod'
	env := envLocal

	log := setupLogger(env)
	slog.SetDefault(log)

	log.Info(
		"starting http server",
		slog.String("env", env),
		slog.String("addr", "http://localhost:8080"),
	)
	log.Debug(fmt.Sprint(srv.ListenAndServe()))
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}),
		)
	case envProd:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
