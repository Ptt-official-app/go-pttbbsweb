package api

import "errors"

var (
	ErrInvalidParams = errors.New("invalid params")
	ErrLoginFailed   = errors.New("login failed")
)
