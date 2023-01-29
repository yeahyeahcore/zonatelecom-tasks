package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/core"
)

type HTTP struct {
	Logger *logrus.Logger
	server *http.Server
	echo   *echo.Echo
	tracer io.Closer
}

func New(logger *logrus.Logger) *HTTP {
	echo := echo.New()
	tracer := jaegertracing.New(echo, nil)

	echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${method}]: ${uri} ${status} ${time_rfc3339} (trace: ${latency_human})\n",
	}))
	echo.Use(middleware.Recover())

	return &HTTP{
		echo:   echo,
		Logger: logger,
		tracer: tracer,
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

func (receiver *HTTP) Start(config *core.HTTPConfiguration) {
	defer func() {
		if err := recover(); err != nil {
			receiver.Logger.Errorf("http server start recover: %s", err)
		}
	}()

	connectionString := fmt.Sprintf("%s:%s", config.Host, config.Port)
	startServerMessage := fmt.Sprintf("Starting HTTP Server on %s", connectionString)

	receiver.Logger.Infoln(startServerMessage)

	if err := receiver.Listen(connectionString); err != nil && err != http.ErrServerClosed {
		receiver.Logger.Infoln("HTTP Listen error:", err)
		panic(err)
	}
}

func (receiver *HTTP) Stop(ctx context.Context) {
	receiver.echo.Shutdown(ctx)
	receiver.server.Shutdown(ctx)
	receiver.tracer.Close()
}
