package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

var BoardFriend_c *db.Collection

type BoardFriend struct {
	// 可看見板的名單

	BoardID bbs.BBoardID `bson:"bid"`
	UserID  bbs.UUserID  `bson:"user_id"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}
