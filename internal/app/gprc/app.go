package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	authgrpc "github.com/Kartochnik010/go-sso/internal/grpc/auth"
	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, port int, auth authgrpc.Auth) *App {
	srv := grpc.NewServer()
	authgrpc.RegisterGRPCServer(srv, auth)
	return &App{
		log:        log,
		port:       port,
		gRPCServer: srv,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"
	log := a.log.With(slog.String("op", op), slog.Int("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("gRPC server started", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"
	a.log.With(slog.String("op", op)).Info("Stopping gRPC server")
	a.gRPCServer.GracefulStop()
}
