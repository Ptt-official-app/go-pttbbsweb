package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/db"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

var UserAloha_c *db.Collection

type UserAloha struct {
	// 上站通知名單
	UserID  bbs.UUserID `bson:"user_id"`
	AlohaID bbs.UUserID `bson:"aloha_id"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}
