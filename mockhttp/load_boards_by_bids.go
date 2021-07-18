package mockhttp

import (
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/types"
)

func LoadBoardsByBids(params *api.LoadBoardsByBidsParams) (ret *api.LoadBoardsByBidsResult) {
	ret = &api.LoadBoardsByBidsResult{
		Boards: []*bbs.BoardSummary{
			{
				BBoardID:     "9_test9",
				BrdAttr:      0,
				StatAttr:     ptttype.NBRD_BOARD,
				Brdname:      "test9",
				RealTitle:    []byte{0xb4, 0xfa, 0xb8, 0xd5, 0x39}, // 測試9
				BoardClass:   []byte{0xb4, 0xfa, 0xb8, 0xd5},       // 測試
				BoardType:    []byte{0xa1, 0xb7},                   //◎
				BM:           []bbs.UUserID{"okcool", "teemo"},
				LastPostTime: types.Time4(1234567890),
				NUser:        100,
				Total:        123,
			},
			{
				BBoardID:     "8_test8",
				BrdAttr:      0,
				StatAttr:     ptttype.NBRD_BOARD,
				Brdname:      "test8",
				RealTitle:    []byte{0xb4, 0xfa, 0xb8, 0xd5, 0x38}, // 測試2
				BoardClass:   []byte{0xb4, 0xfa, 0xb8, 0xd5},       // 測試
				BoardType:    []byte{0xa1, 0xb7},                   //◎
				BM:           []bbs.UUserID{"okcool2", "teemo2"},
				LastPostTime: types.Time4(1300000000),
				NUser:        101,
				Total:        124,
			},
		},
	}

	return ret
}
