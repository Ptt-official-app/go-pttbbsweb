package schema

import "github.com/Ptt-official-app/go-openbbsmiddleware/types"

type UserReject struct {
	//壞人名單

	UserID   types.UUserID `bson:"user_id"`
	RejectID types.UUserID `bson:"reject_id"`

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}
