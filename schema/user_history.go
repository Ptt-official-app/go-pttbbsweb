package schema

import "github.com/Ptt-official-app/go-openbbsmiddleware/types"

type UserHistory struct {
	UserID       types.UUserID `bson:"user_id"`
	Action       string        `bson:"action"`
	RefID        string        `bson:"ref_id"`
	Meta         interface{}   `bson:"meta"`
	UpdateNanoTS types.NanoTS  `bson:"update_nano_ts"`
}
