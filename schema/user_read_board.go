package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

type UserReadBoard struct {
	//已讀板紀錄

	UserID       bbs.UUserID  `bson:"user_id"`
	BBoardID     bbs.BBoardID `bson:"bid"`
	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"`
}

const USER_READ_BOARD_USER_ID_b = "user_id"
const USER_READ_BOARD_BBOARD_ID_b = "bid"
