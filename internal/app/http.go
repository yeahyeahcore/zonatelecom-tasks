package app

import (
	"fmt"
	"net/http"

	"github.com/yeahyeahcore/zonatelecom-tasks/internal/core"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/server"
)

type RunHTTPDeps struct {
	httpServer *server.HTTP
	config     *core.HTTPConfiguration
}

func runHTTP(deps *RunHTTPDeps) {
	defer func() {
		if err := recover(); err != nil {
			deps.httpServer.Logger.Errorf("accrual worker recover: %s", err)
		}
	}()

	connectionString := fmt.Sprintf("%s:%s", deps.config.Host, deps.config.Port)
	startServerMessage := fmt.Sprintf("Starting HTTP Server on %s", connectionString)

	deps.httpServer.Logger.Infoln(startServerMessage)

	if err := deps.httpServer.Listen(connectionString); err != nil && err != http.ErrServerClosed {
		deps.httpServer.Logger.Infoln("HTTP Listen error:", err)
		panic(err)
	}
}
