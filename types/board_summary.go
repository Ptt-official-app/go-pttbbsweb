package types

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

type BoardSummary struct {
	BoardID   string          `json:"boardID" bson:"board_id"`
	Title     string          `json:"title" bson:"title"`
	BrdAttr   ptttype.BrdAttr `json:"flag" bson:"brd_attr"`
	BoardType string          `json:"boardType" bson:"the_type"`
	Category  string          `json:"cat" bson:"category"`
	NUser     int             `json:"onlineCount" bson:"nuser"`
	BMs       []string        `json:"moderators" bson:"bms"`
	Reason    string          `json:"reason" bson:"-"`
	Read      bool            `json:"read" bson:"-"`
	Total     int             `json:"total" bson:"total"`

	LastPostTime_d int                   `json:"-" bson:"last_post_time"`
	StatAttr_d     ptttype.BoardStatAttr `json:"-" bson:"-"`
}

func (b *BoardSummary) Deserialize(b_b *bbs.BoardSummary) {
	b.BoardID = b_b.Brdname
	b.Title = b_b.RealTitle
	b.BrdAttr = b_b.BrdAttr
	b.BoardType = b_b.BoardType
	b.Category = b_b.BoardClass
	b.NUser = int(b_b.NUser)
	b.BMs = b_b.BM
	b.Reason = b_b.Reason
	b.Total = int(b_b.Total)
	b.LastPostTime_d = int(b_b.LastPostTime)
	b.StatAttr_d = b_b.StatAttr
}
