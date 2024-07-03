package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserNewInfo struct {
	Avatar       []byte       `bson:"avatar"`
	AvatarNanoTS types.NanoTS `bson:"avatar_nano_ts"`

	IsDeleted bool `bson:"deleted,omitempty"`
}

var (
	EMPTY_USER_NEW_INFO = &UserNewInfo{}
	userNewInfoFields   = getFields(EMPTY_USER, EMPTY_USER_NEW_INFO)
)

func GetUserNewInfo(userID bbs.UUserID) (result *UserNewInfo, err error) {
	query := &UserQuery{
		UserID: userID,
	}

	err = User_c.FindOne(query, &result, userNewInfoFields)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}
