package service

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/core"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/models"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/repository"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/utils"
)

type VoteRepository interface {
	GetVotingState(ctx context.Context, votingID string) ([]*models.VotingState, error)
	GetVotingStates(ctx context.Context) ([]*models.VotingState, error)
	InsertVote(ctx context.Context, query *models.Vote) (*models.Vote, error)
}

type PrevVotingStateRepository interface {
	GetPreviousVotingStates(ctx context.Context, query *models.VotingState) ([]*models.VotingState, error)
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
	votingState, err := receiver.voteRepository.GetVotingState(ctx, vote.VotingID)
	if err != nil && err != repository.ErrNoRecords {
		receiver.logger.Errorf("failed to get voting state in VoteService method <InsertVote>: %s", err.Error())
		return err
	}

	if _, err := receiver.voteRepository.InsertVote(ctx, vote); err != nil {
		receiver.logger.Errorf("failed to insert vote in VoteService method <InsertVote>: %s", err.Error())
		return err
	}

	if _, err := receiver.previousVotingStateRepository.InsertPreviousVotingStates(ctx, votingState); err != nil {
		receiver.logger.Errorf("failed to insert previous voting states in VoteService method <InsertVote>: %s", err.Error())
		return err
	}

	return nil
}

func (receiver *VoteService) CheckVotingPercentageChange(ctx context.Context) ([]*core.VotingState, error) {
	currentVotingStates, err := receiver.voteRepository.GetVotingStates(ctx)
	if err != nil && err != repository.ErrNoRecords {
		receiver.logger.Errorf("failed to get voting states in VoteService method <CheckVotingPercentageChange>: %s", err.Error())
		return []*core.VotingState{}, err
	}

	previousVotingStates, err := receiver.previousVotingStateRepository.GetPreviousVotingStates(ctx, nil)
	if err != nil {
		receiver.logger.Errorf("failed to get previous voting states in VoteService method <CheckVotingPercentageChange>: %s", err.Error())
		return []*core.VotingState{}, err
	}

	currentVotingStatesCore := utils.TransferVotingStateModelsToCore(currentVotingStates)
	previousVotingStatesCore := utils.TransferVotingStateModelsToCore(previousVotingStates)
	currentVotingStatesMap := utils.TransferVotingStatesToOptionsMap(currentVotingStatesCore)
	previousVotingStatesMap := utils.TransferVotingStatesToOptionsMap(previousVotingStatesCore)

	for _, currentVotingState := range currentVotingStatesMap {
		for _, previousVotingState := range previousVotingStatesMap {
			if currentVotingState.VotingID != previousVotingState.VotingID {
				continue
			}

			difference := utils.GetVotingPercentageDifference(currentVotingState.Options, previousVotingState.Options)
		}
	}

	return nil
}
