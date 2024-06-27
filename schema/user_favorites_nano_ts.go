package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/db"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

var UserFavoritesNanoTS_c *db.Collection

type UserFavoritesNanoTS struct {
	UserID       bbs.UUserID  `bson:"user_id"`
	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}

var EMPTY_USER_FAVORITES_NANO_TS = &UserFavoritesNanoTS{}

var (
	USER_FAVORITES_NANO_TS_USER_ID_b        = getBSONName(EMPTY_USER_FAVORITES_NANO_TS, "UserID")
	USER_FAVORITES_NANO_TS_UPDATE_NANO_TS_b = getBSONName(EMPTY_USER_FAVORITES_NANO_TS, "UpdateNanoTS")
)
