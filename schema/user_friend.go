package schema

import "github.com/Ptt-official-app/go-openbbsmiddleware/types"

type UserFriend struct {
	//朋友名單

	UserID   types.UUserID `bson:"user_id"`
	FriendID types.UUserID `bson:"friend_id"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}
