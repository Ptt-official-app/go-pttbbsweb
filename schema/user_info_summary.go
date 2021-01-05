package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserInfoSummary struct {
	UserID       bbs.UUserID  `bson:"user_id"`
	IsDeleted    bool         `bson:"deleted,omitempty"`
	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}

var (
	EMPTY_USER_INFO_SUMMARY = &UserInfoSummary{}
	userInfoSummaryFields   = getFields(EMPTY_USER, EMPTY_USER_INFO_SUMMARY)
)

func GetUserInfoSummary(userID bbs.UUserID) (result *UserInfoSummary, err error) {
	query := &UserQuery{
		UserID: userID,
	}

	result = &UserInfoSummary{}
	err = User_c.FindOne(query, &result, userInfoSummaryFields)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}
