package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserNewInfo struct {
	Avatar       []byte       `bson:"avatar"`
	AvatarNanoTS types.NanoTS `bson:"avatar_nano_ts"`

	Email               string       `bson:"email"`
	EmailNanoTS         types.NanoTS `bson:"email_naon_ts"`
	EmailVerified       string       `bson:"email_verified"`
	EmailVerifiedNanoTS types.NanoTS `bson:"email_verified_nano_ts"`
	//Phone            string `bson:"phone"` /* 真的要有電話資訊嗎？～ */
	//PhoneVerified    string `bson:"phone_verified"`
	TwoFactorEnabled       bool         `bson:"twofactor_enabled"`
	TwoFactorEnabledNanoTS types.NanoTS `bson:"twofactor_enabled_nano_ts"`

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
