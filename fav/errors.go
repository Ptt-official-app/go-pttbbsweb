package fav

import "errors"

var (
	ErrInvalidFavBoard   = errors.New("invalid fav-board")
	ErrInvalidFavLine    = errors.New("invalid fav-line")
	ErrInvalidFavFolder  = errors.New("invalid fav-folder")
	ErrInvalidFavType    = errors.New("invalid fav-type")
	ErrInvalidFavRecord  = errors.New("invalid fav-record")
	ErrInvalidFav4Record = errors.New("invalid fav4-record")
	ErrOutdatedFav       = errors.New("outdated fav")
	ErrTooManyFavs       = errors.New("too many favs")
	ErrTooManyLines      = errors.New("too many lines")
	ErrTooManyFolders    = errors.New("too many folders")
	ErrTooMuchDepth      = errors.New("too much depth")
)
