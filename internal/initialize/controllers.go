package initialize

import (
	"github.com/sirupsen/logrus"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/controller"
)

type ControllersDeps struct {
	Logger       *logrus.Logger
	Repositories *Repositories
}

type Controllers struct {
	VoteController controller.VoteController
}

func NewControllers(deps *ControllersDeps) *Controllers {
	return &Controllers{
		VoteController: *controller.NewVoteController(&controller.VoteControllerDeps{
			Logger: deps.Logger,
		}),
	}
}
