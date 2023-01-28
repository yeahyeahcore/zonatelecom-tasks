package controller

import (
	"context"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"github.com/yeahyeahcore/zonatelecom-tasks/internal/models"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/repository"
	"github.com/yeahyeahcore/zonatelecom-tasks/pkg/json"
)

type VoteRepository interface {
	GetVotingStates(ctx context.Context, votingID string) ([]*models.VotingState, error)
	InsertVote(ctx context.Context, query *models.Vote) (*models.Vote, error)
}

type VoteControllerDeps struct {
	VoteRepository VoteRepository
	Logger         *logrus.Logger
}

type VoteController struct {
	debtRepository VoteRepository
	logger         *logrus.Logger
}

func NewVoteController(deps *VoteControllerDeps) *VoteController {
	return &VoteController{
		debtRepository: deps.VoteRepository,
		logger:         deps.Logger,
	}
}

func (receiver *VoteController) CreateVote(ctx echo.Context) error {
	voteBody, err := json.Parse[models.Vote](ctx.Request().Body)
	if err != nil {
		return responseBadRequest(ctx, err)
	}

	if _, err := receiver.debtRepository.InsertVote(ctx.Request().Context(), voteBody); err != nil {
		if err == repository.ErrAlreadyExist {
			return responseConflict(ctx, err)
		}

		return responseInternal(ctx, err)
	}

	return responseOK(ctx)
}
