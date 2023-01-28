package repository

import "errors"

var (
	ErrNoRecords    = errors.New("records not found")
	ErrInsertRecord = errors.New("can't insert new record")
	ErrReadRecord   = errors.New("can't read record")
)
