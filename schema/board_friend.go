package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

type BoardFriend struct {
	//可看見板的名單

	BoardID bbs.BBoardID  `bson:"bid"`
	UserID  types.UUserID `bson:"user_id"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}
