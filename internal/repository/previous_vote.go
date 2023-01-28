package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"

	"github.com/yeahyeahcore/zonatelecom-tasks/internal/models"
	"github.com/yeahyeahcore/zonatelecom-tasks/pkg/postgres"
)

type PreviousVoteRepository struct {
	db     *pgxpool.Pool
	logger *logrus.Logger
}

func NewPreviousVoteRepository(deps *RepositoryDeps) *PreviousVoteRepository {
	return &PreviousVoteRepository{
		db:     deps.Database,
		logger: deps.Logger,
	}
}

func (receiver *PreviousVoteRepository) InsertPreviousVotingState(ctx context.Context, model *models.VotingState) (*models.VotingState, error) {
	votes, err := postgres.Insert(ctx, &postgres.QueryWrapper[models.VotingState]{
		DB:        receiver.db,
		TableName: previousVotesTableName,
		Model:     model,
	})
	if err != nil {
		return nil, err
	}
	if len(votes) == 0 {
		return nil, ErrInsertRecord
	}

	return votes[0], nil
}

func (receiver *PreviousVoteRepository) GetPreviousVotingStates(ctx context.Context, model *models.VotingState) ([]*models.VotingState, error) {
	votingStates, err := postgres.Select(ctx, &postgres.QueryWrapper[models.VotingState]{
		DB:        receiver.db,
		TableName: previousVotesTableName,
		Model:     model,
	})
	if err != nil {
		return []*models.VotingState{}, err
	}
	if len(votingStates) == 0 {
		return []*models.VotingState{}, ErrNoRecords
	}

	return votingStates, nil
}
