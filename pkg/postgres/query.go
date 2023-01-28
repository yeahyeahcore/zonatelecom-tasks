package postgres

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/yeahyeahcore/zonatelecom-tasks/pkg/convert"
)

type QueryWrapper[T interface{}] struct {
	DB        *pgxpool.Pool
	TableName string
	Model     *T
	Models    []*T
	SQL       string
}

func (receiver *QueryWrapper[T]) selectQuery() (string, error) {
	if receiver.SQL != "" {
		return receiver.SQL, nil
	}

	sqlBuilder := goqu.From(receiver.TableName)

	if receiver.Model != nil {
		expression, err := convert.ObjectToMap[goqu.Ex](receiver.Model, "db")
		if err != nil {
			return "", err
		}

		sql, _, err := sqlBuilder.Where(expression).ToSQL()
		if err != nil {
			return "", err
		}

		return sql, nil
	}

	sql, _, err := sqlBuilder.ToSQL()
	if err != nil {
		return "", err
	}

	return sql, nil
}

func (receiver *QueryWrapper[T]) insertQuery() (string, error) {
	if receiver.SQL != "" {
		return receiver.SQL, nil
	}

	sqlBuilder := goqu.Insert(receiver.TableName)

	if receiver.Models != nil && len(receiver.Models) > 0 {
		sql, _, err := sqlBuilder.Rows(receiver.Models).ToSQL()
		if err != nil {
			return "", err
		}

		return sql, err
	}

	if receiver.Model != nil {
		sql, _, err := sqlBuilder.Rows([]*T{receiver.Model}).ToSQL()
		if err != nil {
			return "", err
		}

		return sql, err
	}

	sql, _, err := sqlBuilder.ToSQL()
	if err != nil {
		return "", err
	}

	return sql, nil
}
