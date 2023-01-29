package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/yeahyeahcore/zonatelecom-tasks/internal/core"
)

func responseInternal(ctx echo.Context, err error) error {
	statusCode, response := handleErrorResponse(http.StatusInternalServerError, err)

	return ctx.JSON(statusCode, response)
}

func responseBadRequest(ctx echo.Context, err error) error {
	statusCode, response := handleErrorResponse(http.StatusBadRequest, err)

	return ctx.JSON(statusCode, response)
}

func responseOK(ctx echo.Context) error {
	return ctx.JSON(http.StatusBadRequest, core.HTTPErrorResponse{
		Result: "ok",
	})
}

func handleErrorResponse(statusCode int, err error) (int, *core.HTTPErrorResponse) {
	if err != nil {
		errorText := err.Error()

		return statusCode, &core.HTTPErrorResponse{
			Message: &errorText,
			Result:  "fail",
		}
	}

	return statusCode, &core.HTTPErrorResponse{
		Result: "fail",
	}
}
