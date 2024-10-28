package cmd

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/ab-dauletkhan/hot-coffee/internal"
	"github.com/ab-dauletkhan/hot-coffee/internal/core"
	"github.com/ab-dauletkhan/hot-coffee/internal/util"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func Start() {
	// env = 'local' | 'dev' | 'prod'
	env := envLocal

	log := setupLogger(env)
	slog.SetDefault(log)

	err := core.ParseFlags()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	err = util.InitDir()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	log.Info(fmt.Sprintf("working directory: %s", core.Dir))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", core.Port),
		Handler: internal.Routes(),
	}

	log.Info(
		"starting http server",
		slog.String("env", env),
		slog.String("addr", fmt.Sprintf("http://127.0.0.1:%d", core.Port)),
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
