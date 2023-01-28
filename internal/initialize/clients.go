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
	BetaClient *client.BetaClient
}

func NewClients(deps *ClientsDeps) *Clients {
	return &Clients{
		BetaClient: client.NewBetaClient(&client.BetaClientDeps{
			Logger:        deps.Logger,
			Configuration: &deps.ServiceConfiguration.BetaServiceConfiguration,
		}),
	}
}
