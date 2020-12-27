package schema

import (
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/mock_http"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func TestUpdateBoardSummaries(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer Board_c.Drop()

	ret := mock_http.LoadGeneralBoards(nil)

	updateNanoTS := types.NowNanoTS()

	boardSummaries0 := make([]*BoardSummary, len(ret.Boards))
	for idx, each_b := range ret.Boards {
		boardSummaries0[idx] = NewBoardSummary(each_b, updateNanoTS)
	}

	query0 := &BoardQuery{BBoardID: "1_test1"}
	boardSummary0 := &BoardSummary{
		BBoardID:  "1_test1",
		Brdname:   "test1",
		Title:     "測試1",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "測試",
		BMs:       []bbs.UUserID{"okcool", "teemo"},
		Total:     123,

		LastPostTime: types.NanoTS(1234567890000000000),

		UpdateNanoTS: updateNanoTS,
	}

	query1 := &BoardQuery{BBoardID: "2_test2"}
	boardSummary1 := &BoardSummary{
		BBoardID:  "2_test2",
		Brdname:   "test2",
		Title:     "測試2",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "測試",
		BMs:       []bbs.UUserID{"okcool2", "teemo2"},
		Total:     124,

		LastPostTime: types.NanoTS(1300000000000000000),

		UpdateNanoTS: updateNanoTS,
	}

	boardSummary2 := &BoardSummary{
		BBoardID: bbs.BBoardID("3_test3"),

		Brdname:   "test3",
		Title:     "測試3",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "CPBL",
		BMs:       []bbs.UUserID{"okcool3", "teemo3"},
		Total:     125,

		LastPostTime: types.NanoTS(1300000000000000000),

		UpdateNanoTS: updateNanoTS,
	}

	query2 := &BoardQuery{BBoardID: "3_test3"}
	boardSummaries1 := []*BoardSummary{boardSummary1, boardSummary2}

	boardSummary3 := &BoardSummary{
		BBoardID: bbs.BBoardID("2_test2"),

		Brdname:   "test2",
		Title:     "測試4",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "CPBL",
		BMs:       []bbs.UUserID{"okcool4", "teemo4"},
		Total:     125,

		LastPostTime: types.NanoTS(1300000000000000000),

		UpdateNanoTS: updateNanoTS,
	}
	boardSummaries2 := []*BoardSummary{boardSummary2, boardSummary3}

	updateNanoTS1 := types.NowNanoTS()
	boardSummary4 := &BoardSummary{
		BBoardID: bbs.BBoardID("2_test2"),

		Brdname:   "test2",
		Title:     "測試4",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "CPBL",
		BMs:       []bbs.UUserID{"okcool4", "teemo4"},
		Total:     125,

		LastPostTime: types.NanoTS(1300000000000000000),

		UpdateNanoTS: updateNanoTS1,
	}
	boardSummaries3 := []*BoardSummary{boardSummary2, boardSummary4}

	type args struct {
		boardSummaries []*BoardSummary
		updateNanoTS   types.NanoTS
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		query    *BoardQuery
		expected *BoardSummary
	}{
		// TODO: Add test cases.
		{
			args:     args{boardSummaries: boardSummaries0, updateNanoTS: updateNanoTS},
			query:    query0,
			expected: boardSummary0,
		},
		{
			args:     args{boardSummaries: boardSummaries0, updateNanoTS: updateNanoTS},
			query:    query1,
			expected: boardSummary1,
		},
		{
			args:     args{boardSummaries: boardSummaries1, updateNanoTS: updateNanoTS},
			query:    query2,
			expected: boardSummary2,
		},
		{
			args:     args{boardSummaries: boardSummaries2, updateNanoTS: updateNanoTS},
			query:    query1,
			expected: boardSummary1,
		},
		{
			args:     args{boardSummaries: boardSummaries3, updateNanoTS: updateNanoTS1},
			query:    query1,
			expected: boardSummary4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateBoardSummaries(tt.args.boardSummaries, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("UpdateBoardSummaries() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, _ := GetBoardSummary(tt.query.BBoardID)
			testutil.TDeepEqual(t, "got", got, tt.expected)

		})
	}
}
