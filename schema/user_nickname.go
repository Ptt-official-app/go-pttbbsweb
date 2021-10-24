package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserNickname struct {
	Username string `bson:"username"`
	Nickname string `bson:"nickname"`
}

var (
	EMPTY_USER_NICKNAME = &UserNickname{}
	userNicknameFields  = getFields(EMPTY_USER, EMPTY_USER_NICKNAME)
)

func GetUserNickname(userID bbs.UUserID) (nickname string, err error) {
	query := &UserQuery{
		UserID: userID,
	}

	result := &UserNickname{}
	err = User_c.FindOne(query, &result, userNicknameFields)
	if err == mongo.ErrNoDocuments {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return result.Nickname, nil
}
