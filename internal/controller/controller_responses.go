package controller

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/yeahyeahcore/zonatelecom-tasks/internal/core"
)

func responseInternal(ctx echo.Context, err error) error {
	statusCode, response := handleErrorResponse(http.StatusInternalServerError, err)

	return ctx.JSON(statusCode, response)
}

func responseConflict(ctx echo.Context, err error) error {
	statusCode, response := handleErrorResponse(http.StatusInternalServerError, err)

	return ctx.JSON(statusCode, response)
}

func responseNotFound(ctx echo.Context, err error) error {
	statusCode, response := handleErrorResponse(http.StatusNotFound, err)

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
