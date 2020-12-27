package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

var (
	UserFriend_c *db.Collection
)

type UserFriend struct {
	//朋友名單

	UserID   bbs.UUserID `bson:"user_id"`
	FriendID bbs.UUserID `bson:"friend_id"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}
