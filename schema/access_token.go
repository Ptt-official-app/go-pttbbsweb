package schema

import "github.com/Ptt-official-app/go-openbbsmiddleware/types"

type AccessToken struct {
	AccessToken  string       `bson:"access_token"`
	UserID       string       `bson:"user_id"`
	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}
