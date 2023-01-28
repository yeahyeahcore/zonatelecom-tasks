package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/yeahyeahcore/zonatelecom-tasks/internal/server"
)

type gracefulShutdownDeps struct {
	httpServer *server.HTTP
	database   *pgxpool.Pool
	cancel     context.CancelFunc
}

func gracefulShutdown(ctx context.Context, deps *gracefulShutdownDeps) {
	defer deps.httpServer.Stop(ctx)
	defer deps.database.Close()
	defer deps.cancel()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	deps.httpServer.Logger.Infoln("Shutting down server ...")
}
