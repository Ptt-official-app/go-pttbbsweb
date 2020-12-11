package main

import "errors"

var (
	ErrInvalidHost       = errors.New("invalid host")
	ErrInvalidRemoteAddr = errors.New("invalid remote addr")
	ErrInvalidToken      = errors.New("invalid token")
	ErrInvalidIni        = errors.New("invalid ini")
)

type errResult struct {
	Msg string
}
