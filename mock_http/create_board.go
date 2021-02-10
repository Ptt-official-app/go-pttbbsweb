package mock_http

import (
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

func CreateBoard(params *api.CreateBoardParams) (ret api.CreateBoardResult) {
	ret = api.CreateBoardResult(&bbs.BoardSummary{
		Bid:        13,
		Gid:        2,
		BBoardID:   "13_test3",
		BrdAttr:    0,
		StatAttr:   ptttype.NBRD_BOARD,
		Brdname:    "test3",
		RealTitle:  []byte{0xb4, 0xfa, 0xb8, 0xd5, 0x31}, //測試1
		BoardClass: []byte{0xb4, 0xfa, 0xb8, 0xd5},       //測試
		BoardType:  []byte{0xa1, 0xb7},                   //◎
		BM:         []bbs.UUserID{"okcool", "teemo"},
	})

	return ret
}
