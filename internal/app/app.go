package app

import (
	"time"

	"log/slog"

	grpcapp "github.com/Kartochnik010/go-sso/internal/app/gprc"
	"github.com/Kartochnik010/go-sso/internal/service/auth"
	"github.com/Kartochnik010/go-sso/internal/storage/sqlite"
)

type App struct {
	GRPCServer  *grpcapp.App
	Port        int
	StoragePath string
	TokenTTL    time.Duration
}

func New(log *slog.Logger, port int, storagePath string, TokenTTL time.Duration) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, TokenTTL)

	grpcApp := grpcapp.New(log, port, authService)

	return &App{
		GRPCServer:  grpcApp,
		Port:        port,
		StoragePath: storagePath,
		TokenTTL:    TokenTTL,
	}
}
