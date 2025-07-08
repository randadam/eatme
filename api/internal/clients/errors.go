package clients

import "errors"

var (
	ErrMLCallFailed = errors.New("ML gateway call failed")
	ErrMLBadRequest = errors.New("ML gateway bad request")
)
