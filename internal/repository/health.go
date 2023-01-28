package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type HealthRepository struct {
	db     *pgxpool.Pool
	logger *logrus.Logger
}

func NewHealthRepository(deps *RepositoryDeps) *HealthRepository {
	return &HealthRepository{
		db:     deps.Database,
		logger: deps.Logger,
	}
}

func (receiver *HealthRepository) Check(ctx context.Context) error {
	if err := receiver.db.Ping(ctx); err != nil {
		receiver.logger.Errorln("health check error:", err.Error())
		return err
	}

	return nil
}
