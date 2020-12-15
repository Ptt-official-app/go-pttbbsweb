package mock_http

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/backend"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/types"
)

func LoadGeneralBoards(params *backend.LoadGeneralBoardsParams) (ret *backend.LoadGeneralBoardsResult) {
	ret = &backend.LoadGeneralBoardsResult{
		Boards: []*bbs.BoardSummary{
			{
				BBoardID:     "1_test1",
				BrdAttr:      0,
				StatAttr:     ptttype.NBRD_BOARD,
				Brdname:      "test1",
				RealTitle:    "測試1",
				BoardClass:   "測試",
				BoardType:    "◎",
				BM:           []string{"okcool", "teemo"},
				LastPostTime: types.Time4(1234567890),
				NUser:        100,
				Total:        123,
			},
			{
				BBoardID:     "2_test2",
				BrdAttr:      0,
				StatAttr:     ptttype.NBRD_BOARD,
				Brdname:      "test2",
				RealTitle:    "測試2",
				BoardClass:   "測試",
				BoardType:    "◎",
				BM:           []string{"okcool2", "teemo2"},
				LastPostTime: types.Time4(1300000000),
				NUser:        101,
				Total:        124,
			},
		},
		NextIdx: "testNextIdx",
	}

	return ret
}
