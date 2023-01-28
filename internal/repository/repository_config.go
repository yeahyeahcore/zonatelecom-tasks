package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type RepositoryDeps struct {
	Logger   *logrus.Logger
	Database *pgxpool.Pool
}
