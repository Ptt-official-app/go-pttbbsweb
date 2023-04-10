package mockhttp

import (
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/types"
)

func LoadGeneralBoardsByClass(params *api.LoadGeneralBoardsParams) (ret *api.LoadGeneralBoardsResult) {
	switch params.StartIdx {
	case "vFSt-Q@WhoAmI":
		ret = &api.LoadGeneralBoardsResult{
			Boards: []*bbs.BoardSummary{
				{
					Gid:      5,
					Bid:      10,
					BBoardID: bbs.BBoardID("10_WhoAmI"),
					StatAttr: ptttype.NBRD_FAV,
					Brdname:  "WhoAmI",
					BoardClass: []byte{
						0xbc, 0x54, 0xad, 0xf9,
					},
					RealTitle: []byte{
						0xa8, 0xfe, 0xa8, 0xfe, 0xa1, 0x41, 0xb2, 0x71, 0xb2, 0x71,
						0xa7, 0xda, 0xac, 0x4f, 0xbd, 0xd6, 0xa1, 0x49,
					},
					BoardType:  []byte{0xa1, 0xb7},
					BM:         []bbs.UUserID{},
					IdxByName:  "WhoAmI",
					IdxByClass: "vFSt-Q@WhoAmI",
				},
			},
		}
	default:
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
			NextIdx: "vFSt-Q@test3",
		}
	}

	return ret
}
