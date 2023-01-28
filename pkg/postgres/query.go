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
