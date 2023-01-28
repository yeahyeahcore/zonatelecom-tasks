package service

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/models"
)

type VoteRepository interface {
	GetVotingStates(ctx context.Context, votingID string) ([]*models.VotingState, error)
	InsertVote(ctx context.Context, query *models.Vote) (*models.Vote, error)
}

type PrevVotingStateRepository interface {
	GetPreviousVotingStates(ctx context.Context, query models.VotingState) ([]*models.VotingState, error)
	InsertPreviousVotingState(ctx context.Context, query *models.VotingState) (*models.VotingState, error)
}

type VoteServiceDeps struct {
	Logger                    *logrus.Logger
	VoteRepository            VoteRepository
	PrevVotingStateRepository PrevVotingStateRepository
}

type VoteService struct {
	logger                        *logrus.Logger
	voteRepository                VoteRepository
	previousVotingStateRepository PrevVotingStateRepository
}

func NewVoteService(deps *VoteServiceDeps) *VoteService {
	return &VoteService{
		logger:                        deps.Logger,
		voteRepository:                deps.VoteRepository,
		previousVotingStateRepository: deps.PrevVotingStateRepository,
	}
}
