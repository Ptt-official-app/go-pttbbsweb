package db

import "errors"

var (
	ErrInvalidOp      = errors.New("invalid op")
	ErrInvalidSlice   = errors.New("invalid slice")
	ErrNoUpdate       = errors.New("no update")
	ErrEmptyInRemove  = errors.New("nil in remove")
	ErrNotSliceInFind = errors.New("not slice in find")
)
