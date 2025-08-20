package mock

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
)

type BoardListResult_t struct {
	List    []*apitypes.BoardSummary `json:"list"`
	NextIdx string                   `json:"next_idx"`
}

var BoardListResult = &BoardListResult_t{
	List: []*apitypes.BoardSummary{
		{
			FBoardID:  apitypes.FBoardID("WhoAmI"),
			Brdname:   "WhoAmI",
			Title:     "我～是～誰？～",
			BrdAttr:   0,
			BoardType: "◎",
			Category:  "嘰哩",
			NUser:     39,
			BMs:       []bbs.UUserID{"okcool", "teemo"},
			Reason:    "",
			Read:      true,
			Total:     134,
		},
		{
			FBoardID:  apitypes.FBoardID("ALLPOST"),
			Brdname:   "ALLPOST",
			Title:     "所有文章都底家",
			BrdAttr:   ptttype.BRD_POSTMASK,
			BoardType: "◎",
			Category:  "嘰哩",
			NUser:     234,
			BMs:       []bbs.UUserID{"test1"},
			Reason:    "",
			Read:      false,
			Total:     123124,
		},
	},
	NextIdx: "3",
}
