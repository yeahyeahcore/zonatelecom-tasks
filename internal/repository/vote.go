package repository

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"

	"github.com/yeahyeahcore/zonatelecom-tasks/internal/models"
	"github.com/yeahyeahcore/zonatelecom-tasks/pkg/postgres"
)

type VoteRepository struct {
	db     *pgxpool.Pool
	logger *logrus.Logger
}

func NewVoteRepository(deps *RepositoryDeps) *VoteRepository {
	return &VoteRepository{
		db:     deps.Database,
		logger: deps.Logger,
	}
}

func (receiver *VoteRepository) InsertVote(ctx context.Context, model *models.Vote) (*models.Vote, error) {
	votes, err := postgres.Insert(ctx, &postgres.QueryWrapper[models.Vote]{
		DB:        receiver.db,
		TableName: votesTableName,
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

func (receiver *VoteRepository) GetVotingState(ctx context.Context, votingID string) ([]*models.VotingState, error) {
	sql, err := receiver.constituteVotingStateSQL(votingID)
	if err != nil {
		return []*models.VotingState{}, err
	}

	votingStates, err := postgres.Select(ctx, &postgres.QueryWrapper[models.VotingState]{
		DB:  receiver.db,
		SQL: sql,
	})
	if err != nil {
		return []*models.VotingState{}, err
	}
	if len(votingStates) == 0 {
		return []*models.VotingState{}, ErrNoRecords
	}

	return votingStates, nil
}

func (receiver *VoteRepository) GetVotingStates(ctx context.Context) ([]*models.VotingState, error) {
	sql, err := receiver.constituteVotingStateSQL("")
	if err != nil {
		return []*models.VotingState{}, err
	}

	votingStates, err := postgres.Select(ctx, &postgres.QueryWrapper[models.VotingState]{
		DB:  receiver.db,
		SQL: sql,
	})
	if err != nil {
		return []*models.VotingState{}, err
	}
	if len(votingStates) == 0 {
		return []*models.VotingState{}, ErrNoRecords
	}

	return votingStates, nil
}

func (receiver *VoteRepository) constituteVotingStateSQL(votindID string) (string, error) {
	dataset := goqu.Select(goqu.L("voting_id, option_id, COUNT(option_id")).From(votesTableName)

	if votindID == "" {
		sql, _, err := dataset.
			GroupBy(goqu.L("voting_id, option_id")).
			ToSQL()

		if err != nil {
			return "", err
		}

		return sql, nil
	}

	sql, _, err := dataset.
		Where(goqu.Ex{"voting_id": votindID}).
		GroupBy(goqu.L("voting_id, option_id")).
		ToSQL()

	if err != nil {
		return "", err
	}

	return sql, nil
}
