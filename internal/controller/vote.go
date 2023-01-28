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
	GetVotingState(ctx context.Context, query *models.Vote) (*models.VotingState, error)
	CreateVote(ctx context.Context, query *models.Vote) (*models.Vote, error)
}

type VoteControllerDeps struct {
	VotesRepository VoteRepository
	Logger          *logrus.Logger
}

type VotesController struct {
	debtRepository VoteRepository
	logger         *logrus.Logger
}

func NewVotesController(deps *VoteControllerDeps) *VotesController {
	return &VotesController{
		debtRepository: deps.VotesRepository,
		logger:         deps.Logger,
	}
}

func (receiver *VotesController) CreateVote(ctx echo.Context) error {
	voteBody, err := json.Parse[models.Vote](ctx.Request().Body)
	if err != nil {
		return responseBadRequest(ctx, err)
	}

	if _, err := receiver.debtRepository.CreateVote(ctx.Request().Context(), voteBody); err != nil {
		if err == repository.ErrAlreadyExist {
			return responseConflict(ctx, err)
		}

		return responseInternal(ctx, err)
	}

	return responseOK(ctx)
}
