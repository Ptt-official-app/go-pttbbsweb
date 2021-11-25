package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/bson"
)

var UserVisit_c *db.Collection

type UserVisit struct {
	UserID       bbs.UUserID  `bson:"user_id"`
	Action       string       `bson:"action"`
	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}

var EMPTY_USER_VISIT = &UserVisit{}

var (
	USER_VISIT_USER_ID_b        = getBSONName(EMPTY_USER_VISIT, "UserID")
	USER_VISIT_ACTION_b         = getBSONName(EMPTY_USER_VISIT, "Action")
	USER_VISIT_UPDATE_NANO_TS_b = getBSONName(EMPTY_USER_VISIT, "UpdateNanoTS")
)

func UpdateUserVisit(userVisit *UserVisit) (err error) {
	query := bson.M{
		USER_VISIT_USER_ID_b: userVisit.UserID,
	}

	_, err = UserVisit_c.Update(query, userVisit)
	if err != nil {
		return err
	}
	return
}
