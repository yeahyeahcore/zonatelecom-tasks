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

	repositories := initialize.NewRepositories(&initialize.RepositoriesDeps{
		Logger:   logger,
		Database: database,
	})

	clients := initialize.NewClients(&initialize.ClientsDeps{
		Logger:               logger,
		ServiceConfiguration: &config.Service,
	})

	services := initialize.NewServices(&initialize.ServicesDeps{
		Logger:       logger,
		Repositories: repositories,
		Clients:      clients,
	})

	controllers := initialize.NewControllers(&initialize.ControllersDeps{
		Logger:       logger,
		Repositories: repositories,
		Services:     services,
	})

	workers := initialize.NewWorkers(&initialize.WorkersDeps{
		Logger:   logger,
		Services: services,
	})

	httpServer := server.New(logger).Register(controllers)

	go httpServer.Start(&config.HTTP)

	workers.Run(ctx)

	gracefulShutdown(ctx, &gracefulShutdownDeps{
		httpServer: httpServer,
		database:   database,
		cancel:     cancel,
	})
}
