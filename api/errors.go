package api

import "errors"

var (
	ErrInvalidRemoteAddr = errors.New("invalid remote addr")

	ErrInvalidParams            = errors.New("invalid params")
	ErrLoginFailed              = errors.New("login failed")
	ErrNotImplemented           = errors.New("not implemented")
	ErrInvalidPath              = errors.New("invalid path")
	ErrInvalidToken             = errors.New("invalid token")
	ErrInvalidOrigin            = errors.New("invalid origin")
	ErrInvalidBackendStatusCode = errors.New("invalid backend status code")
	ErrNoArticle                = errors.New("no article")
	ErrAlreadyDeleted           = errors.New("already deleted")
	ErrFileNotFound             = errors.New("file not found")
)
