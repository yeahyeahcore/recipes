package service

import "errors"

var (
	ErrEmptyArgs = errors.New("empty arguments")
	ErrInternal  = errors.New("internal error")
)
