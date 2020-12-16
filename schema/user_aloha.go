package schema

import "github.com/Ptt-official-app/go-openbbsmiddleware/types"

type UserAloha struct {
	//上站通知名單

	UserID  types.UUserID `bson:"user_id"`
	AlohaID types.UUserID `bson:"aloha_id"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}
