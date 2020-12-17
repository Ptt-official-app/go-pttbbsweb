package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

type UserAloha struct {
	//上站通知名單
	UserID  bbs.UUserID `bson:"user_id"`
	AlohaID bbs.UUserID `bson:"aloha_id"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}
