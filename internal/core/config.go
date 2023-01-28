package core

import (
	"github.com/yeahyeahcore/zonatelecom-tasks/pkg/postgres"
)

type Config struct {
	HTTP     HTTPConfiguration
	Database postgres.PostgreSQLConfiguration
}

type HTTPConfiguration struct {
	Host string `env:"HTTP_HOST,default=localhost"`
	Port string `env:"HTTP_PORT,default=8080"`
}