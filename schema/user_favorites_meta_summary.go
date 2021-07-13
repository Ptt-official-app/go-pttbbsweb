package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/bson"
)

type UserFavoritesMetaSummary struct {
	UserID          bbs.UUserID  `bson:"user_id"`
	DoubleBufferIdx int          `bson:"db_idx"`
	MTime           types.NanoTS `bson:"mtime_nano_ts"`
}

var (
	EMPTY_USER_FAVORITES_META_SUMMARY      = &UserFavoritesMetaSummary{}
	userFavoritesMetaDoubleBufferIdxFields = getFields(EMPTY_USER_FAVORITES_META, EMPTY_USER_FAVORITES_META_SUMMARY)
)

func GetUserFavoritesMetaSummary(userID bbs.UUserID) (result *UserFavoritesMetaSummary, err error) {
	query := bson.M{
		USER_FAVORITES_META_USER_ID_b: userID,
	}

	result = &UserFavoritesMetaSummary{}
	err = UserFavoritesMeta_c.FindOne(query, result, userFavoritesMetaDoubleBufferIdxFields)
	if err != nil {
		return nil, err
	}

	return result, nil
}
