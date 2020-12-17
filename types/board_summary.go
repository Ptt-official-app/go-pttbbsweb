package types

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

type BoardSummary struct {
	BBoardID  bbs.BBoardID    `json:"bid" bson:"bid"`
	Brdname   string          `json:"brdname" bson:"brdname"`
	Title     string          `json:"title" bson:"title"`
	BrdAttr   ptttype.BrdAttr `json:"flag" bson:"flag"`
	BoardType string          `json:"type" bson:"the_type"`
	Category  string          `json:"class" bson:"class"`
	NUser     int             `json:"nuser" bson:"nuser"`
	BMs       []bbs.UUserID   `json:"moderators" bson:"bms"`
	Reason    string          `json:"reason" bson:"-"`
	Read      bool            `json:"read" bson:"-"`
	Total     int             `json:"total" bson:"total"`

	LastPostTimeTS_d Time8                 `json:"-" bson:"last_post_time"`
	StatAttr_d       ptttype.BoardStatAttr `json:"-" bson:"-"`
}

func (b *BoardSummary) Deserialize(b_b *bbs.BoardSummary) {
	b.BBoardID = b_b.BBoardID
	b.Brdname = b_b.Brdname
	b.Title = b_b.RealTitle
	b.BrdAttr = b_b.BrdAttr
	b.BoardType = b_b.BoardType
	b.Category = b_b.BoardClass
	b.NUser = int(b_b.NUser)
	b.BMs = b_b.BM
	b.Reason = b_b.Reason
	b.Total = int(b_b.Total)
	b.Read = b_b.Read

	b.LastPostTimeTS_d = Time8(b_b.LastPostTime)
	b.StatAttr_d = b_b.StatAttr
}
