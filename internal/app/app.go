package app

import (
	"time"

	"log/slog"

	grpcapp "github.com/Kartochnik010/go-sso/internal/app/gprc"
)

type App struct {
	GRPCServer  *grpcapp.App
	Port        int
	StoragePath string
	TokenTTL    time.Duration
}

func New(log *slog.Logger, port int, storagePath string, TokenTTL time.Duration) *App {
	grpcApp := grpcapp.New(log, port)
	return &App{
		GRPCServer:  grpcApp,
		Port:        port,
		StoragePath: storagePath,
		TokenTTL:    TokenTTL,
	}
}
