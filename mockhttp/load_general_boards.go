package mockhttp

import (
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/types"
)

func LoadGeneralBoards(params *api.LoadGeneralBoardsParams) (ret *api.LoadGeneralBoardsResult) {
	ret = &api.LoadGeneralBoardsResult{
		Boards: []*bbs.BoardSummary{
			{
				BBoardID:     "1_test1",
				BrdAttr:      0,
				StatAttr:     ptttype.NBRD_BOARD,
				Brdname:      "test1",
				RealTitle:    []byte{0xb4, 0xfa, 0xb8, 0xd5, 0x31}, // 測試1
				BoardClass:   []byte{0xb4, 0xfa, 0xb8, 0xd5},       // 測試
				BoardType:    []byte{0xa1, 0xb7},                   //◎
				BM:           []bbs.UUserID{"okcool", "teemo"},
				LastPostTime: types.Time4(1234567890),
				NUser:        100,
				Total:        123,
			},
			{
				BBoardID:     "2_test2",
				BrdAttr:      0,
				StatAttr:     ptttype.NBRD_BOARD,
				Brdname:      "test2",
				RealTitle:    []byte{0xb4, 0xfa, 0xb8, 0xd5, 0x32}, // 測試2
				BoardClass:   []byte{0xb4, 0xfa, 0xb8, 0xd5},       // 測試
				BoardType:    []byte{0xa1, 0xb7},                   //◎
				BM:           []bbs.UUserID{"okcool2", "teemo2"},
				LastPostTime: types.Time4(1300000000),
				NUser:        101,
				Total:        124,
			},
		},
		NextIdx: "test3",
	}

	return ret
}
