package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	pttbbsfav "github.com/Ptt-official-app/go-pttbbs/ptt/fav"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"go.mongodb.org/mongo-driver/bson"
)

type UserFavoriteID struct {
	UserID          bbs.UUserID  `bson:"user_id"`
	DoubleBufferIdx int          `bson:"dbuffer_idx"`
	MTime           types.NanoTS `bson:"mtime_nano_ts"`

	LevelIdx LevelIdx `bson:"level_idx"`

	TheType pttbbsfav.FavT `bson:"the_type"`
	TheID   int            `bson:"the_id"`
}

var EMPTY_USER_FAVORITE_ID = &UserFavoriteID{}

func GetUserFavoriteIDsByPttbids(userID bbs.UUserID, doubleBufferIdx int, pttbids []ptttype.Bid, mTime types.NanoTS) (userFavoriteIDs []*UserFavoriteID, err error) {
	query := bson.M{
		USER_FAVORITES_USER_ID_b:           userID,
		USER_FAVORITES_DOUBLE_BUFFER_IDX_b: doubleBufferIdx,
		USER_FAVORITES_MTIME_b:             mTime,
		USER_FAVORITES_THE_TYPE_b:          pttbbsfav.FAVT_BOARD,
		USER_FAVORITES_THE_ID_b: bson.M{
			"$in": pttbids,
		},
	}

	err = UserFavorites_c.Find(query, 0, &userFavoriteIDs, nil, nil)
	if err != nil {
		return nil, err
	}

	return userFavoriteIDs, nil
}
