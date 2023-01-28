package controller

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/yeahyeahcore/zonatelecom-tasks/internal/core"
)

func internalError(ctx echo.Context, err error) error {
	return ctx.JSON(http.StatusInternalServerError, core.HTTPErrorResponse{
		Message:    err.Error(),
		StatusCode: http.StatusInternalServerError,
	})
}

func notFoundError(ctx echo.Context, err error) error {
	return ctx.JSON(http.StatusNotFound, core.HTTPErrorResponse{
		Message:    err.Error(),
		StatusCode: http.StatusNotFound,
	})
}

func badRequestError(ctx echo.Context, err error) error {
	return ctx.JSON(http.StatusBadRequest, core.HTTPErrorResponse{
		Message:    err.Error(),
		StatusCode: http.StatusBadRequest,
	})
}

func authError(ctx echo.Context, err error) error {
	return ctx.JSON(http.StatusUnauthorized, core.HTTPErrorResponse{
		Message:    err.Error(),
		StatusCode: http.StatusUnauthorized,
	})
}
