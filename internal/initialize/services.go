package initialize

import (
	"github.com/sirupsen/logrus"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/service"
)

type ServicesDeps struct {
	Logger       *logrus.Logger
	Repositories *Repositories
	Clients      *Clients
}

type Services struct {
	VoteService *service.VoteService
}

func NewServices(deps *ServicesDeps) *Services {
	return &Services{
		VoteService: service.NewVoteService(&service.VoteServiceDeps{
			Logger:                    deps.Logger,
			VoteRepository:            deps.Repositories.VoteRepository,
			PrevVotingStateRepository: deps.Repositories.PreviousVotingStateRepository,
			BetaClient:                deps.Clients.BetaClient,
		}),
	}
}
