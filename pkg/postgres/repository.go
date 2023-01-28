package postgres

import (
	"context"
	"errors"

	"github.com/georgysavva/scany/v2/pgxscan"
)

func Select[T interface{}](ctx context.Context, wrapper *QueryWrapper[T]) (*[]T, error) {
	if wrapper == nil {
		return nil, errors.New("wrapper is nil")
	}

	models := make([]T, 0)

	sql, err := wrapper.selectQuery()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(ctx, wrapper.DB, &models, sql); err != nil {
		return nil, err
	}

	return &models, nil
}
