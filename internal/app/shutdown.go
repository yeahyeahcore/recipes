package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	"penahub.gitlab.yandexcloud.net/pena-services/accruals-service/internal/server"
)

type gracefulShutdownDeps struct {
	server *server.HTTP
	pool   *pgxpool.Pool
	cancel context.CancelFunc
}

func gracefulShutdown(ctx context.Context, deps *gracefulShutdownDeps) {
	defer deps.server.Stop(ctx)
	defer deps.pool.Close()
	defer deps.cancel()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
}
