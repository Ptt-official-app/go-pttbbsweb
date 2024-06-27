package schema

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/db"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

var UserReject_c *db.Collection

type UserReject struct {
	// 壞人名單

	UUserID  bbs.UUserID `bson:"user_id"`
	RejectID bbs.UUserID `bson:"reject_id"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}
