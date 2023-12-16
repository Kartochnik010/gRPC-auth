package suite

import (
	"context"
	"net"
	"strconv"
	"testing"

	"github.com/Kartochnik010/go-sso/internal/config"
	ssov1 "github.com/Kartochnik010/go-sso/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcHost = "localhost"
)

type Suite struct {
	*testing.T
	Cfg        *config.Config
	AuthClient ssov1.AuthClient
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := config.MustLoadByPath("../config/config.yml")
	ctx, cancelCtx := context.WithTimeout(context.TODO(), cfg.GRPC.Timeout)

	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	cc, err := grpc.DialContext(ctx, grpcAdress(cfg), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal("grpc server connection failed: ", err)
	}

	return ctx, &Suite{
		AuthClient: ssov1.NewAuthClient(cc),
		Cfg:        cfg,
		T:          t,
	}
}

func grpcAdress(cfg *config.Config) string {
	return net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.Port))
}
