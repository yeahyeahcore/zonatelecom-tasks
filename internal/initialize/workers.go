package initialize

import (
	"github.com/sirupsen/logrus"

	"github.com/yeahyeahcore/zonatelecom-tasks/internal/worker"
)

type WorkersDeps struct {
	Logger   *logrus.Logger
	Services *Services
}

type Workers struct {
	VoteWorker worker.VoteWorker
}

func NewWorkers(deps *WorkersDeps) *Workers {
	return &Workers{
		VoteWorker: *worker.NewVoteWorker(&worker.VoteWorkerDeps{
			Logger:      deps.Logger,
			VoteService: deps.Services.VoteService,
		}),
	}
}
