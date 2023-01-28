package controller

import (
	"context"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"github.com/yeahyeahcore/zonatelecom-tasks/internal/models"
	"github.com/yeahyeahcore/zonatelecom-tasks/pkg/json"
)

type VoteRepository interface {
	GetVotingStates(ctx context.Context, votingID string) ([]*models.VotingState, error)
	InsertVote(ctx context.Context, query *models.Vote) (*models.Vote, error)
}

type VoteService interface {
	InsertVote(ctx context.Context, vote *models.Vote) error
}

type VoteControllerDeps struct {
	Logger         *logrus.Logger
	VoteRepository VoteRepository
	VoteService    VoteService
}

type VoteController struct {
	logger         *logrus.Logger
	voteRepository VoteRepository
	voteService    VoteService
}

func NewVoteController(deps *VoteControllerDeps) *VoteController {
	return &VoteController{
		logger:         deps.Logger,
		voteRepository: deps.VoteRepository,
		voteService:    deps.VoteService,
	}
}

func (receiver *VoteController) CreateVote(ctx echo.Context) error {
	voteBody, err := json.Parse[models.Vote](ctx.Request().Body)
	if err != nil {
		return responseBadRequest(ctx, err)
	}

	if err := receiver.voteService.InsertVote(ctx.Request().Context(), voteBody); err != nil {
		return responseInternal(ctx, err)
	}

	return responseOK(ctx)
}
