package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

type AccessToken struct {
	AccessToken  string       `bson:"access_token"`
	UserID       bbs.UUserID  `bson:"user_id"`
	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}
