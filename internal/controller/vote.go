package controller

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/yeahyeahcore/zonatelecom-tasks/internal/client"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/core"
	"github.com/yeahyeahcore/zonatelecom-tasks/pkg/json"
)

type voteService interface {
	InsertVote(ctx context.Context, vote *core.CreateVoteRequest) error
}

type VoteControllerDeps struct {
	Logger      *logrus.Logger
	VoteService voteService
}

type VoteController struct {
	logger      *logrus.Logger
	voteService voteService
}

func NewVoteController(deps *VoteControllerDeps) *VoteController {
	return &VoteController{
		logger:      deps.Logger,
		voteService: deps.VoteService,
	}
}

func (receiver *VoteController) CreateVote(ctx echo.Context) error {
	voteBody, err := json.Parse[core.CreateVoteRequest](ctx.Request().Body)
	if err != nil {
		return responseBadRequest(ctx, err)
	}

	if err := receiver.voteService.InsertVote(ctx.Request().Context(), voteBody); err != nil {
		if err == client.ErrWrongDigest {
			return responseBadRequest(ctx, err)
		}

		return responseInternal(ctx, err)
	}

	return responseOK(ctx)
}
