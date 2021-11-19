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
	ErrNoBoard                  = errors.New("no board")
	ErrNoArticle                = errors.New("no article")
	ErrAlreadyDeleted           = errors.New("already deleted")
	ErrFileNotFound             = errors.New("file not found")

	ErrInvalidUser   = errors.New("invalid user")
	ErrInvalidClient = errors.New("invalid client")

	ErrAlreadyExists = errors.New("already exists")

	ErrInvalidFav = errors.New("invalid fav")
)
