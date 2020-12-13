package schema

import "github.com/Ptt-official-app/go-openbbsmiddleware/types"

type UserReadBoard struct {
	UserID       string       `bson:"user_id"`
	BoardID      string       `bson:"board_id"`
	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}

const USER_READ_BOARD_USER_ID_b = "user_id"
const USER_READ_BOARD_BOARD_ID_b = "board_id"
