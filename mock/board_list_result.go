package mock

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

type BoardListResult_t struct {
	List    []*types.BoardSummary `json:"list"`
	NextIdx string                `json:"next_idx"`
}

var (
	BoardListResult = &BoardListResult_t{
		List: []*types.BoardSummary{
			{
				BBoardID:  bbs.BBoardID("10_WhoAmI"),
				Brdname:   "WhoAmI",
				Title:     "我～是～誰？～",
				BrdAttr:   0,
				BoardType: "◎",
				Category:  "嘰哩",
				NUser:     39,
				BMs:       []string{"okcool", "teemo"},
				Reason:    "",
				Read:      true,
				Total:     134,
			},
			{
				BBoardID:  bbs.BBoardID("6_ALLPOST"),
				Brdname:   "ALLPOST",
				Title:     "所有文章都底家",
				BrdAttr:   ptttype.BRD_POSTMASK,
				BoardType: "◎",
				Category:  "嘰哩",
				NUser:     234,
				BMs:       []string{"test1"},
				Reason:    "",
				Read:      false,
				Total:     123124,
			},
		},
		NextIdx: "3",
	}
)
