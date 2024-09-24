package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"go.mongodb.org/mongo-driver/bson"
)

type UserPosttime struct {
	UserID               bbs.UUserID  `bson:"user_id"`
	Posttime             int          `bson:"postitme"`
	PosttimeUpdateNanoTS types.NanoTS `bson:"posttime_update_nano_ts"`
}

var EMPTY_USER_POSTTIME = &UserPosttime{}

func UpdateUserPosttime(userID bbs.UUserID, postTime int) (err error) {
	nowNanoTS := types.NowNanoTS()
	query := bson.M{
		USER_USER_ID_b: userID,
		USER_POSTTIME_UPDATE_NANO_TS_b: bson.M{
			"$lt": nowNanoTS,
		},
	}
	update := bson.M{
		USER_POSTTIME_b:                postTime,
		USER_POSTTIME_UPDATE_NANO_TS_b: nowNanoTS,
	}

	_, err = User_c.UpdateOneOnly(query, update)

	return err
}
