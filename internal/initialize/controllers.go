package initialize

import (
	"github.com/sirupsen/logrus"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/controller"
)

type ControllersDeps struct {
	Logger       *logrus.Logger
	Repositories *Repositories
	Services     *Services
}

type Controllers struct {
	VoteController        *controller.VoteController
	HealthCheckController *controller.HealthCheckController
}

func NewControllers(deps *ControllersDeps) *Controllers {
	return &Controllers{
		VoteController: controller.NewVoteController(&controller.VoteControllerDeps{
			Logger:      deps.Logger,
			VoteService: deps.Services.VoteService,
		}),
		HealthCheckController: controller.NewHealthCheckController(&controller.HealthCheckControllerDeps{
			HealthCheckRepository: deps.Repositories.HealthRepository,
		}),
	}
}
