package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

var (
	BoardBanuser_c *db.Collection
)

type BoardBanUser struct {
	//板的水桶名單 (不可 po 文)

	BoardID   bbs.BBoardID `bson:"bid"`
	UserID    bbs.UUserID  `bson:"user_id"`
	EndNanoTS types.NanoTS `bson:"end_nano_ts"`
	PosterID  bbs.UUserID  `bson:"poster_id"`
	Reason    string       `bson:"reason"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}
