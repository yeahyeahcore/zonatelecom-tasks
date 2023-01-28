package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"github.com/yeahyeahcore/zonatelecom-tasks/internal/models"
	"github.com/yeahyeahcore/zonatelecom-tasks/pkg/json"
)

type DebtRepository interface {
	GetDebts(ctx context.Context, queryModel *models.Debt) (*[]models.Debt, error)
	GetDebt(ctx context.Context, queryModel *models.Debt) (*models.Debt, error)
}

type DebtControllerDeps struct {
	DebtRepository DebtRepository
	Logger         *logrus.Logger
}

type DebtController struct {
	debtRepository DebtRepository
	logger         *logrus.Logger
}

func NewDebtController(deps *DebtControllerDeps) *DebtController {
	return &DebtController{
		debtRepository: deps.DebtRepository,
		logger:         deps.Logger,
	}
}

func (receiver *DebtController) GetAllDebts(ctx echo.Context) error {
	debts, err := receiver.debtRepository.GetDebts(ctx.Request().Context(), nil)
	if err != nil {
		if err == repository.ErrNoRecord {
			return notFoundError(ctx, err)
		}

		return internalError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, debts)
}

func (receiver *DebtController) GetDebts(ctx echo.Context) error {
	debtBody, err := json.Parse[models.Debt](ctx.Request().Body)
	if err != nil {
		return badRequestError(ctx, err)
	}

	debt, err := receiver.debtRepository.GetDebts(ctx.Request().Context(), debtBody)
	if err != nil {
		if err == repository.ErrNoRecord {
			return notFoundError(ctx, err)
		}

		return internalError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, debt)
}

func (receiver *DebtController) GetDebtByID(ctx echo.Context) error {
	id := ctx.Param("id")
	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return badRequestError(ctx, err)
	}

	debt, err := receiver.debtRepository.GetDebt(ctx.Request().Context(), &models.Debt{ID: uint(parsedID)})
	if err != nil {
		if err == repository.ErrNoRecord {
			return notFoundError(ctx, err)
		}

		return internalError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, debt)
}
