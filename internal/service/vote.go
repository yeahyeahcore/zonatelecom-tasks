package service

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/models"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/repository"
)

type VoteRepository interface {
	GetVotingStates(ctx context.Context, votingID string) ([]*models.VotingState, error)
	InsertVote(ctx context.Context, query *models.Vote) (*models.Vote, error)
}

type PrevVotingStateRepository interface {
	GetPreviousVotingStates(ctx context.Context, query models.VotingState) ([]*models.VotingState, error)
	InsertPreviousVotingState(ctx context.Context, query *models.VotingState) (*models.VotingState, error)
	InsertPreviousVotingStates(ctx context.Context, query []*models.VotingState) ([]*models.VotingState, error)
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

func (receiver *VoteService) InsertVote(ctx context.Context, vote *models.Vote) error {
	votingStates, err := receiver.voteRepository.GetVotingStates(ctx, vote.VotingID)
	if err != nil && err != repository.ErrNoRecords {
		receiver.logger.Errorf("failed to get voting states in VoteService method <InsertVote>: %s", err.Error())
		return err
	}

	if _, err := receiver.voteRepository.InsertVote(ctx, vote); err != nil {
		receiver.logger.Errorf("failed to insert vote in VoteService method <InsertVote>: %s", err.Error())
		return err
	}

	if _, err := receiver.previousVotingStateRepository.InsertPreviousVotingStates(ctx, votingStates); err != nil {
		receiver.logger.Errorf("failed to insert previous voting states in VoteService method <InsertVote>: %s", err.Error())
		return err
	}

	return nil
}
