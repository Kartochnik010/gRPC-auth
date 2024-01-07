package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"log/slog"

	"github.com/Kartochnik010/go-sso/internal/app"
	"github.com/Kartochnik010/go-sso/internal/config"
	"github.com/Kartochnik010/go-sso/internal/lib/logger/sl"
	slogpretty "github.com/Kartochnik010/go-sso/internal/lib/logger/slogretty"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	log := NewLogger(cfg.Env)

	application, err := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)
	if err != nil {
		log.Error("failed to init app", sl.Err(err))
	}

	successfulStart := make(chan struct{}, 1)
	go func() {
		if err := application.GRPCServer.Run(); err != nil {
			log.Error("failed to start gRPC server", sl.Err(err))
			os.Exit(1)
		}
		successfulStart <- struct{}{}
	}()
	<-successfulStart

	log.Info("Application started",
		slog.String("env", cfg.Env),
		slog.Any("cfg", cfg),
	)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()
	log.Info("Gracefully stopped")
}

func NewLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case "local":
		logger = setupPrettySlog()
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

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
