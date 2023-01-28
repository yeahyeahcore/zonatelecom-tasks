package main

import (
	"fmt"
	"path"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/yeahyeahcore/zonatelecom-tasks/internal/app"
	"github.com/yeahyeahcore/zonatelecom-tasks/internal/core"
	"github.com/yeahyeahcore/zonatelecom-tasks/pkg/env"
)

func main() {
	logger := logrus.New()

	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05",
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := fmt.Sprintf(" (%s:%s)", path.Base(frame.File), strconv.Itoa(frame.Line))
			return "", fileName
		},
	})

	config, err := env.Parse[core.Config]("./.env.example")
	if err != nil {
		logger.Fatalf("config read is failed: %s", err)
	}

	app.Run(config, logger)
}
