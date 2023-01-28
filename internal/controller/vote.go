package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"github.com/yeahyeahcore/zonatelecom-tasks/internal/models"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/repository"
	"github.com/yeahyeahcore/zonatelecom-tasks/pkg/json"
)

type VoteRepository interface {
	GetVotes(ctx context.Context, queryModel *models.Vote) (*[]models.Vote, error)
	GetVote(ctx context.Context, queryModel *models.Vote) (*models.Vote, error)
	CreateVote(ctx context.Context, queryModel *models.Vote) (*models.Vote, error)
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

func (receiver *VotesController) GetAllVotess(ctx echo.Context) error {
	debts, err := receiver.debtRepository.GetVotes(ctx.Request().Context(), nil)
	if err != nil {
		if err == repository.ErrNoRecords {
			return notFoundError(ctx, err)
		}

		return internalError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, debts)
}

func (receiver *VotesController) GetVotes(ctx echo.Context) error {
	debtBody, err := json.Parse[models.Vote](ctx.Request().Body)
	if err != nil {
		return badRequestError(ctx, err)
	}

	debt, err := receiver.debtRepository.GetVotes(ctx.Request().Context(), debtBody)
	if err != nil {
		if err == repository.ErrNoRecords {
			return notFoundError(ctx, err)
		}

		return internalError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, debt)
}

func (receiver *VotesController) GetVoteByID(ctx echo.Context) error {
	id := ctx.Param("id")
	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return badRequestError(ctx, err)
	}

	debt, err := receiver.debtRepository.GetVotes(ctx.Request().Context(), &models.Vote{ID: uint(parsedID)})
	if err != nil {
		if err == repository.ErrNoRecords {
			return notFoundError(ctx, err)
		}

		return internalError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, debt)
}
