package mockhttp

import (
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/types"
	"github.com/sirupsen/logrus"
)

func LoadBoardsByBids(params *api.LoadBoardsByBidsParams) (ret *api.LoadBoardsByBidsResult) {
	logrus.Infof("mockhttp.LoadBoardsByBids: params: %v", params)

	boardsMap := map[ptttype.Bid]*bbs.BoardSummary{
		10: {
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
		9: {
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

			Gid:        3,
			Bid:        9,
			IdxByName:  "test9",
			IdxByClass: "tPq41Q@test1",
		},
		8: {
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

			Gid:        3,
			Bid:        8,
			IdxByClass: "tPq41Q@test2",
		},
		1: {
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

			Gid: 3,
			Bid: 1,

			IdxByName:  "test1",
			IdxByClass: "tPq41Q@test1",
		},
		2: {
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

			Gid:        3,
			Bid:        2,
			IdxByName:  "test2",
			IdxByClass: "tPq41Q@test2",
		},
	}
	boards := make([]*bbs.BoardSummary, len(params.Bids))
	for idx, each := range params.Bids {
		boards[idx] = boardsMap[each]
	}
	ret = &api.LoadBoardsByBidsResult{
		Boards: boards,
	}

	return ret
}
