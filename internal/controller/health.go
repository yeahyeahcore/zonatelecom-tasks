package controller

import (
	"context"

	"github.com/labstack/echo/v4"
)

type HealthCheckRepository interface {
	Check(ctx context.Context) error
}

type HealthCheckControllerDeps struct {
	HealthCheckRepository HealthCheckRepository
}

type HealthCheckController struct {
	healthCheckRepository HealthCheckRepository
}

func NewHealthCheckController(deps *HealthCheckControllerDeps) *HealthCheckController {
	return &HealthCheckController{
		healthCheckRepository: deps.HealthCheckRepository,
	}
}

func (receiver *HealthCheckController) Health(ctx echo.Context) error {
	if err := receiver.healthCheckRepository.Check(ctx.Request().Context()); err != nil {
		return responseServiceUnavailable(ctx)
	}

	return responseOK(ctx)
}

func (receiver *HealthCheckController) Ready(ctx echo.Context) error {
	return responseOK(ctx)
}
