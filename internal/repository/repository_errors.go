package repository

import "errors"

var (
	ErrNoRecord             = errors.New("no record in DB")
	ErrCreateRecord         = errors.New("failed to create record")
	ErrScanRecord           = errors.New("failed to scan record")
	ErrGetRecord            = errors.New("failed to get record")
	ErrUpdateRecord         = errors.New("failed to update record")
	ErrDeleteRecord         = errors.New("failed to delete record")
	ErrIterateRows          = errors.New("failed to iterate rows")
	ErrMethodNotImplemented = errors.New("method is not implemented")
	ErrEmptyArgs            = errors.New("arguments are empty")
)
