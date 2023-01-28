package app

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/yeahyeahcore/zonatelecom-tasks/internal/core"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/initialize"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/server"
	"github.com/yeahyeahcore/zonatelecom-tasks/pkg/postgres"
)

func Run(config *core.Config, logger *logrus.Logger) {
	ctx, cancel := context.WithCancel(context.Background())

	database, err := postgres.Connect(ctx, &postgres.PostgresConnectDeps{
		Configuration: &config.Database,
		Timeout:       10 * time.Second,
	})
	if err != nil {
		logger.Fatalln("failed connection to db: ", err)
		panic(err)
	}

	controllers := initialize.NewControllers(&initialize.ControllersDeps{
		Logger: logger,
	})

	httpServer := server.New(logger).Register(controllers)

	go runHTTP(&RunHTTPDeps{
		httpServer: httpServer,
		config:     &config.HTTP,
	})

	// gracefulShutdown(ctx, &gracefulShutdownDeps{
	// 	httpServer: httpServer,
	// 	database:   database,
	// 	cancel:     cancel,
	// })
}
