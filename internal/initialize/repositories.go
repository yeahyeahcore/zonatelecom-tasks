package initialize

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"

	"github.com/yeahyeahcore/zonatelecom-tasks/internal/repository"
)

type RepositoriesDeps struct {
	Database *pgxpool.Pool
	Logger   *logrus.Logger
}

type Repositories struct {
	HealthRepository       *repository.HealthRepository
	VoteRepository         *repository.VoteRepository
	PreviousVoteRepository *repository.PreviousVoteRepository
}

func NewRepositories(deps *RepositoriesDeps) *Repositories {
	return &Repositories{
		HealthRepository: repository.NewHealthRepository(&repository.RepositoryDeps{
			Logger:   deps.Logger,
			Database: deps.Database,
		}),
		VoteRepository: repository.NewVoteRepository(&repository.RepositoryDeps{
			Logger:   deps.Logger,
			Database: deps.Database,
		}),
		PreviousVoteRepository: repository.NewPreviousVoteRepository(&repository.RepositoryDeps{
			Logger:   deps.Logger,
			Database: deps.Database,
		}),
	}
}
