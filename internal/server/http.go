package server

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

type HTTP struct {
	Logger *logrus.Logger
	server *http.Server
	echo   *echo.Echo
}

func New(logger *logrus.Logger) *HTTP {
	echo := echo.New()

	echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${method}]: ${uri} ${status} ${time_rfc3339} (trace: ${latency_human})\n",
	}))
	echo.Use(middleware.Recover())

	return &HTTP{
		echo:   echo,
		Logger: logger,
		server: &http.Server{
			Handler:        echo,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}
}

func (receiver *HTTP) Listen(address string) error {
	receiver.server.Addr = address

	return receiver.server.ListenAndServe()
}

func (receiver *HTTP) Stop(ctx context.Context) {
	receiver.echo.Shutdown(ctx)
	receiver.server.Shutdown(ctx)
}
