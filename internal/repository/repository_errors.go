package repository

import "errors"

var (
	ErrNoRecords    = errors.New("records not found")
	ErrInsertRecord = errors.New("no records inseted")
	ErrAlreadyExist = errors.New("can't insert exists record")
	ErrReadRecord   = errors.New("can't read record")
)
