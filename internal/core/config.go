package core

import (
	"github.com/yeahyeahcore/zonatelecom-tasks/pkg/postgres"
)

type Config struct {
	HTTP     HTTPConfiguration
	Database postgres.PostgreSQLConfiguration
	Service  ServiceConfiguration
}

type HTTPConfiguration struct {
	Host string `env:"HTTP_HOST,default=localhost"`
	Port string `env:"HTTP_PORT,default=8080"`
}

type ServiceConfiguration struct {
	GammaServiceConfiguration  GammaServiceConfiguration
	DigestServiceConfiguration DigestServiceConfiguration
}

type GammaServiceConfiguration struct {
	BaseURL string `env:"GAMMA_SERVICE_URL,default=localhost"`
}

type DigestServiceConfiguration struct {
	BaseURL string `env:"DIGEST_SERVICE_URL,default=localhost"`
}
