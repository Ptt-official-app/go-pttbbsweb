package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/db"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"go.mongodb.org/mongo-driver/bson"
)

var AccessToken_c *db.Collection

type AccessToken struct {
	AccessToken  string       `bson:"access_token"`
	UserID       bbs.UUserID  `bson:"user_id"`
	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}

var EMPTY_ACCESS_TOKEN = &AccessToken{}

var (
	ACCESS_TOKEN_ACCESS_TOKEN_b   = getBSONName(EMPTY_ACCESS_TOKEN, "AccessToken")
	ACCESS_TOKEN_USER_ID_b        = getBSONName(EMPTY_ACCESS_TOKEN, "UserID")
	ACCESS_TOKEN_UPDATE_NANO_TS_b = getBSONName(EMPTY_ACCESS_TOKEN, "UpdateNanoTS")
)

func NewAccessToken(userID bbs.UUserID, jwt string, nowNanoTS types.NanoTS) *AccessToken {
	return &AccessToken{
		AccessToken:  jwt,
		UserID:       userID,
		UpdateNanoTS: nowNanoTS,
	}
}

func UpdateAccessToken(accessToken *AccessToken) (err error) {
	query := bson.M{
		ACCESS_TOKEN_USER_ID_b:      accessToken.UserID,
		ACCESS_TOKEN_ACCESS_TOKEN_b: accessToken.AccessToken,
	}

	_, err = AccessToken_c.Update(query, accessToken)

	return err
}
