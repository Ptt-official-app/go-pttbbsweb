package schema

import (
	"time"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
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

// UpdateUserVisit updates user_visit's document data
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

// CalculateAllUserVisitCounts count recent 10 mins users who are calling login_required_api
func CalculateAllUserVisitCounts() (int64, error) {
	query := bson.M{
		USER_VISIT_UPDATE_NANO_TS_b: bson.M{
			"$gte": types.TimeToNanoTS(time.Now().Add(TIME_CALC_ALL_USER_VISIT_COUNTS)),
		},
	}

	count, err := UserVisit_c.Count(query, 0)
	if err != nil {
		return 0, err
	}

	return count, nil
}
