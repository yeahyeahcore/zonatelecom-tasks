package initialize

import (
	"github.com/sirupsen/logrus"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/client"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/core"
)

type ClientsDeps struct {
	Logger               *logrus.Logger
	ServiceConfiguration *core.ServiceConfiguration
}

type Clients struct {
	GammaClient *client.GammaClient
}

func NewClients(deps *ClientsDeps) *Clients {
	return &Clients{
		GammaClient: client.NewGammaClient(&client.GammaClientDeps{
			Logger:        deps.Logger,
			Configuration: &deps.ServiceConfiguration.GammaServiceConfiguration,
		}),
	}
}
