package schema

import "errors"

var (
	ErrNoLock  = errors.New("no lock")
	ErrNoMatch = errors.New("no match")
)
