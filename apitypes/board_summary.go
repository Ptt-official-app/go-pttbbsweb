package apitypes

import (
	"strconv"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	pttbbsfav "github.com/Ptt-official-app/go-pttbbs/ptt/fav"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

type BoardSummary struct {
	FBoardID     FBoardID        `json:"bid"`
	Brdname      string          `json:"brdname"`
	Title        string          `json:"title"`
	BrdAttr      ptttype.BrdAttr `json:"flag"`
	BoardType    string          `json:"type"`
	Category     string          `json:"class"`
	NUser        int             `json:"nuser"`
	BMs          []bbs.UUserID   `json:"moderators"`
	Reason       string          `json:"reason"`
	Read         bool            `json:"read"`
	Fav          bool            `json:"fav"`
	Total        int             `json:"total"`
	LastPostTime types.Time8     `json:"last_post_time"`

	StatAttr ptttype.BoardStatAttr `json:"stat_attr,omitempty"`
	LevelIdx schema.LevelIdx       `json:"level_idx,omitempty"` // sub-level-idx for folder, "" if type is line or board.

	URL string `json:"url,omitempty"`

	Gid ptttype.Bid `json:"gid"`
	Bid ptttype.Bid `json:"pttbid"`

	Idx string `json:"idx"`

	TokenUser bbs.UUserID `json:"tokenuser,omitempty"`
}

func NewBoardSummary(b_db *schema.BoardSummary, idx string, userBoardInfo *UserBoardInfo, userID bbs.UUserID) *BoardSummary {
	if b_db == nil {
		return nil
	}

	fboardID := ToFBoardID(b_db.BBoardID)

	url := ""
	if b_db.BrdAttr.HasPerm(ptttype.BRD_GROUPBOARD) {
		// XXX api.LOAD_CLASS_BOARDS_R
		bidStr := strconv.Itoa(int(b_db.Bid))
		url = "/cls/" + bidStr
	} else {
		// XXX api.LOAD_GENERAL_ARTICLES_R
		url = "/board/" + string(fboardID) + "/articles"
	}

	bms := b_db.BMs
	if bms == nil {
		bms = []bbs.UUserID{}
	}
	return &BoardSummary{
		FBoardID:     fboardID,
		Brdname:      b_db.Brdname,
		Title:        b_db.Title,
		BrdAttr:      b_db.BrdAttr,
		BoardType:    b_db.BoardType,
		Category:     b_db.Category,
		BMs:          bms,
		Total:        b_db.Total,
		LastPostTime: b_db.LastPostTime.ToTime8(),
		NUser:        b_db.NUser,

		Gid: b_db.Gid,
		Bid: b_db.Bid,

		Idx: idx,

		StatAttr: userBoardInfo.Stat,
		Read:     userBoardInfo.Read,
		Fav:      userBoardInfo.Fav,

		URL: url,

		TokenUser: userID,
	}
}

func NewBoardSummaryFromUserFavorites(userID bbs.UUserID, uf_db *schema.UserFavorites, b_db *schema.BoardSummary, userBoardInfo *UserBoardInfo) *BoardSummary {
	idxStr := strconv.Itoa(uf_db.Idx)

	switch uf_db.TheType {
	case pttbbsfav.FAVT_LINE:
		return &BoardSummary{
			StatAttr: ptttype.NBRD_LINE,
			Idx:      idxStr,
			BMs:      []bbs.UUserID{},

			TokenUser: userID,
		}
	case pttbbsfav.FAVT_FOLDER:
		// XXX api.LOAD_FAVORITE_BOARDS_R
		subLevelIdx := schema.SetLevelIdx(uf_db.LevelIdx, uf_db.Idx)
		url := "/user/" + string(userID) + "/favorites?level_idx=" + string(subLevelIdx)
		return &BoardSummary{
			Title: uf_db.FolderTitle,

			StatAttr: ptttype.NBRD_FOLDER,
			LevelIdx: subLevelIdx,
			Idx:      idxStr,
			URL:      url,
			BMs:      []bbs.UUserID{},

			TokenUser: userID,
		}

	case pttbbsfav.FAVT_BOARD:
		return NewBoardSummary(b_db, idxStr, userBoardInfo, userID)
	}

	return nil
}
