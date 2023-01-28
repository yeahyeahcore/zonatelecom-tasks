package postgres

import (
	"context"
	"errors"

	"github.com/georgysavva/scany/v2/pgxscan"
)

func Select[T interface{}](ctx context.Context, wrapper *QueryWrapper[T]) ([]*T, error) {
	if wrapper == nil {
		return nil, errors.New("wrapper is nil")
	}

	models := make([]*T, 0)

	sql, err := wrapper.selectQuery()
	if err != nil {
		return nil, err
	}

	if err := pgxscan.Select(ctx, wrapper.DB, &models, sql); err != nil {
		return nil, err
	}

	return models, nil
}

func Insert[T interface{}](ctx context.Context, wrapper *QueryWrapper[T]) ([]*T, error) {
	if wrapper == nil {
		return []*T{}, errors.New("wrapper is nil")
	}

	sql, err := wrapper.insertQuery()
	if err != nil {
		return []*T{}, err
	}

	rows, err := wrapper.DB.Query(ctx, sql)
	if err != nil {
		return []*T{}, err
	}

	if wrapper.Models != nil && len(wrapper.Models) > 0 {
		models := make([]*T, len(wrapper.Models))

		if err := pgxscan.ScanAll(models, rows); err != nil {
			return []*T{}, err
		}

		return models, nil
	}

	model := new(T)

	if err := pgxscan.ScanOne(model, rows); err != nil {
		return []*T{}, err
	}

	return []*T{model}, nil
}
