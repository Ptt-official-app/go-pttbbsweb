package apitypes

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	pttbbsfav "github.com/Ptt-official-app/go-pttbbs/ptt/fav"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

type BoardSummary struct {
	BBoardID     bbs.BBoardID    `json:"bid"`
	Brdname      string          `json:"brdname"`
	Title        string          `json:"title"`
	BrdAttr      ptttype.BrdAttr `json:"flag"`
	BoardType    string          `json:"type"`
	Category     string          `json:"class"`
	NUser        int             `json:"nuser"`
	BMs          []bbs.UUserID   `json:"moderators"`
	Reason       string          `json:"reason"`
	Read         bool            `json:"read"`
	Total        int             `json:"total"`
	LastPostTime types.Time8     `json:"last_post_time"`

	StatAttr ptttype.BoardStatAttr `json:"stat_attr,omitempty"`
	LevelIdx schema.LevelIdx       `json:"level_idx,omitempty"`
}

func NewBoardSummary(b_db *schema.BoardSummary) *BoardSummary {
	if b_db == nil {
		return nil
	}
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
		NUser:        b_db.NUser,

		StatAttr: ptttype.NBRD_BOARD,
	}
}

func NewBoardSummaryFromUserFavorites(uf_db *schema.UserFavorites, b_db *schema.BoardSummary) *BoardSummary {
	switch uf_db.TheType {
	case pttbbsfav.FAVT_LINE:
		return &BoardSummary{
			StatAttr: ptttype.NBRD_LINE,
		}

	case pttbbsfav.FAVT_FOLDER:
		return &BoardSummary{
			Title: uf_db.FolderTitle,

			StatAttr: ptttype.NBRD_FOLDER,
			LevelIdx: schema.SetLevelIdx(uf_db.LevelIdx, uf_db.Idx),
		}

	case pttbbsfav.FAVT_BOARD:
		return NewBoardSummary(b_db)
	}

	return nil
}
