package mockhttp

import (
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/types"
)

func LoadFullClassBoards(params *api.LoadFullClassBoardsParams) (ret *api.LoadFullClassBoardsResult) {
	var boards []*bbs.BoardSummary
	if params.StartBid == 1 {
		boards = []*bbs.BoardSummary{
			{
				BBoardID: "3_3..........",
				BrdAttr:  ptttype.BRD_GROUPBOARD,
				StatAttr: ptttype.NBRD_BOARD,
				Brdname:  "3..........",
				RealTitle: []byte{ //.... Σ市民廣場     報告  站長  ㄜ！
					0xa5, 0xab, 0xa5,
					0xc1, 0xbc, 0x73, 0xb3, 0xf5, 0x20, 0x20, 0x20, 0x20, 0x20,
					0xb3, 0xf8, 0xa7, 0x69, 0x20, 0x20, 0xaf, 0xb8, 0xaa, 0xf8,
					0x20, 0x20, 0xa3, 0xad, 0xa1, 0x49,
				}, // 測試9
				BoardClass:   []byte{0x2e, 0x2e, 0x2e, 0x2e}, // ....
				BoardType:    []byte{0xa3, 0x55},             // Σ
				BM:           []bbs.UUserID{"okcool", "teemo"},
				LastPostTime: types.Time4(1234567890),

				Gid:        1,
				Bid:        3,
				IdxByName:  "3..........",
				IdxByClass: "....@3..........",
			},
		}
	}

	return &api.LoadFullClassBoardsResult{
		Boards: boards,
	}
}
