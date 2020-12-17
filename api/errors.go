package api

import "errors"

var (
	ErrInvalidParams  = errors.New("invalid params")
	ErrLoginFailed    = errors.New("login failed")
	ErrNotImplemented = errors.New("not implemented")
	ErrInvalidPath    = errors.New("invalid path")
	ErrInvalidToken   = errors.New("invalid token")
)
