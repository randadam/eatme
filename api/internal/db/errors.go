package db

import "errors"

var (
	ErrEmailExists = errors.New("email already exists")
	ErrNotFound    = errors.New("not found")
)
