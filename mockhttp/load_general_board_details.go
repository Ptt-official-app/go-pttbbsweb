package mockhttp

import (
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/types"
)

func LoadGeneralBoardDetails(params *api.LoadGeneralBoardDetailsParams) (ret *api.LoadGeneralBoardDetailsResult) {
	ret = &api.LoadGeneralBoardDetailsResult{
		Boards: []*bbs.BoardDetail{
			{
				BBoardID:     "1_test1",
				BrdAttr:      0,
				Brdname:      "test1",
				RealTitle:    []byte{0xb4, 0xfa, 0xb8, 0xd5, 0x31}, // 測試1
				BoardClass:   []byte{0xb4, 0xfa, 0xb8, 0xd5},       // 測試
				BoardType:    []byte{0xa1, 0xb7},                   //◎
				BM:           []bbs.UUserID{"okcool", "teemo"},
				LastPostTime: types.Time4(1234567890),
				NUser:        100,
				Total:        123,

				Gid: 3,
				Bid: 1,

				IdxByName:  "test1",
				IdxByClass: "tPq41Q@test1",
			},
			{
				BBoardID:     "2_test2",
				BrdAttr:      0,
				Brdname:      "test2",
				RealTitle:    []byte{0xb4, 0xfa, 0xb8, 0xd5, 0x32}, // 測試2
				BoardClass:   []byte{0xb4, 0xfa, 0xb8, 0xd5},       // 測試
				BoardType:    []byte{0xa1, 0xb7},                   //◎
				BM:           []bbs.UUserID{"okcool2", "teemo2"},
				LastPostTime: types.Time4(1300000000),
				NUser:        101,
				Total:        124,

				PostType: [][]byte{{0xb4, 0xfa, 0xb8, 0xd5}, {0xa1, 0xb7}, {0x00, 0x00, 0x00, 0x00}},

				Gid:        3,
				Bid:        2,
				IdxByName:  "test2",
				IdxByClass: "tPq41Q@test2",
			},
		},
		NextIdx: "test3",
	}

	return ret
}
