package apitypes

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

type BoardSummary struct {
	BBoardID     bbs.BBoardID    `json:"bid" bson:"bid"`
	Brdname      string          `json:"brdname" bson:"brdname"`
	Title        string          `json:"title" bson:"title"`
	BrdAttr      ptttype.BrdAttr `json:"flag" bson:"flag"`
	BoardType    string          `json:"type" bson:"the_type"`
	Category     string          `json:"class" bson:"class"`
	NUser        int             `json:"nuser" bson:"nuser"`
	BMs          []bbs.UUserID   `json:"moderators" bson:"bms"`
	Reason       string          `json:"reason" bson:"-"`
	Read         bool            `json:"read" bson:"-"`
	Total        int             `json:"total" bson:"total"`
	LastPostTime types.Time8     `json:"last_post_time"`

	StatAttr ptttype.BoardStatAttr `json:"stat_attr"`
}

func NewBoardSummary(b_db *schema.BoardSummary) *BoardSummary {
	return &BoardSummary{
		BBoardID:     b_db.BBoardID,
		Brdname:      b_db.Brdname,
		Title:        b_db.Title,
		BrdAttr:      b_db.BrdAttr,
		BoardType:    b_db.BoardType,
		Category:     b_db.Category,
		BMs:          b_db.BMs,
		Total:        b_db.Total,
		LastPostTime: b_db.LastPostTime.ToTime8(),
	}
}
