package worker

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/yeahyeahcore/zonatelecom-tasks/internal/core"
)

type voteService interface {
	CheckVotingPercentageChange(ctx context.Context) ([]*core.PreviousVotingState, error)
	SendVotingStatesToGamma(ctx context.Context, votingStates []*core.PreviousVotingState) error
}

type VoteWorkerDeps struct {
	Logger      *logrus.Logger
	VoteService voteService
}

type VoteWorker struct {
	logger      *logrus.Logger
	voteService voteService
}

func NewVoteWorker(deps *VoteWorkerDeps) *VoteWorker {
	return &VoteWorker{
		logger:      deps.Logger,
		voteService: deps.VoteService,
	}
}

func (receiver *VoteWorker) Run(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			receiver.logger.Errorf("vote worker recover: %s", err)
		}
	}()

	timer := time.NewTimer(1 * time.Second)

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			voteStates, err := receiver.voteService.CheckVotingPercentageChange(ctx)
			if err != nil {
				receiver.logger.Errorln(err)
				continue
			}
			if len(voteStates) < 1 {
				receiver.logger.Errorln("empty vote states for sending to beta service")
				continue
			}

			if err := receiver.voteService.SendVotingStatesToGamma(ctx, voteStates); err != nil {
				receiver.logger.Errorln(err)
			}
		}
	}

}
