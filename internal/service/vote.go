package service

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/core"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/models"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/repository"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/utils"
)

type voteRepository interface {
	GetVotingState(ctx context.Context, votingID string) ([]*models.VotingState, error)
	GetVotingStates(ctx context.Context) ([]*models.VotingState, error)
	InsertVote(ctx context.Context, query *models.Vote) (*models.Vote, error)
}

type prevVotingStateRepository interface {
	GetPreviousVotingStates(ctx context.Context, query *models.VotingState) ([]*models.VotingState, error)
	InsertPreviousVotingStates(ctx context.Context, query []*models.VotingState) ([]*models.VotingState, error)
}

type gammaClient interface {
	SendVotingState(request *core.VotingState) error
}

type digestClient interface {
	Check(digest string) error
}

type VoteServiceDeps struct {
	Logger                    *logrus.Logger
	VoteRepository            voteRepository
	PrevVotingStateRepository prevVotingStateRepository
	GammaClient               gammaClient
	DigestClient              digestClient
}

type VoteService struct {
	logger                        *logrus.Logger
	voteRepository                voteRepository
	previousVotingStateRepository prevVotingStateRepository
	gammaClient                   gammaClient
	digestClient                  digestClient
}

func NewVoteService(deps *VoteServiceDeps) *VoteService {
	return &VoteService{
		logger:                        deps.Logger,
		voteRepository:                deps.VoteRepository,
		previousVotingStateRepository: deps.PrevVotingStateRepository,
		gammaClient:                   deps.GammaClient,
		digestClient:                  deps.DigestClient,
	}
}

func (receiver *VoteService) InsertVote(ctx context.Context, vote *core.CreateVoteRequest) error {
	if err := receiver.digestClient.Check(vote.Digest); err != nil {
		receiver.logger.Errorf("failed to check digest in VoteService method <InsertVote>: %s", err.Error())
		return err
	}

	votingState, err := receiver.voteRepository.GetVotingState(ctx, vote.VotingID)
	if err != nil && err != repository.ErrNoRecords {
		receiver.logger.Errorf("failed to get voting state in VoteService method <InsertVote>: %s", err.Error())
		return err
	}

	if _, err := receiver.voteRepository.InsertVote(ctx, &models.Vote{
		VoteID:   vote.VoteID,
		VotingID: vote.VotingID,
		OptionID: vote.OptionID,
	}); err != nil {
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

	currentVotingStatesMap := utils.TransferVotingStatesToOptionsMap(currentVotingStates)
	previousVotingStatesMap := utils.TransferVotingStatesToOptionsMap(previousVotingStates)

	changedVotingState := make([]*core.VotingState, 0)

	for _, currentVotingState := range currentVotingStatesMap {
		for _, previousVotingState := range previousVotingStatesMap {
			if currentVotingState.VotingID != previousVotingState.VotingID {
				continue
			}

			if difference := utils.GetVotingPercentageDifference(currentVotingState.Options, previousVotingState.Options); len(difference) >= 1 {
				changedVotingState = append(changedVotingState, utils.TransferVotingStateToCore(currentVotingState))
			}
		}
	}

	return changedVotingState, nil
}

func (receiver *VoteService) SendVotingStatesToGamma(ctx context.Context, votingStates []*core.VotingState) error {
	for _, votingState := range votingStates {
		if err := receiver.gammaClient.SendVotingState(votingState); err != nil {
			receiver.logger.Errorf("failed to send states in VoteService method <SendVotingStatesToGamma>: %s", err.Error())
			return err
		}
	}

	return nil
}
