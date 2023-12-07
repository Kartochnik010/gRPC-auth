package main

import (
	"log"
	"os"

	"log/slog"

	"github.com/Kartochnik010/go-sso/internal/app"
	"github.com/Kartochnik010/go-sso/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	log := NewLogger(cfg.Env)
	_ = cfg
	log.Info("Application started",
		slog.String("env", cfg.Env),
		slog.Any("cfg", cfg),
	)

	app := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)
	
}

func NewLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case "local":
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "dev":
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "prod":
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return logger
}
