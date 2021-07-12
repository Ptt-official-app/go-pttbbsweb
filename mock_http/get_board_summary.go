package mock_http

import (
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/types"
)

func GetBoardSummary(params *api.LoadBoardSummaryParams) (ret api.LoadBoardSummaryResult) {
	ret = api.LoadBoardSummaryResult(&bbs.BoardSummary{
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
	})

	return ret
}
