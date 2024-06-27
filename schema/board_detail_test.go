package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/go-pttbbsweb/mockhttp"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"github.com/sirupsen/logrus"
)

func TestUpdateBoardDetails(t *testing.T) {
	setupTest()
	defer teardownTest()

	ret := mockhttp.LoadGeneralBoardDetails(nil)

	updateNanoTS := types.NowNanoTS()

	boardDetails0 := make([]*BoardDetail, len(ret.Boards))
	for idx, each_b := range ret.Boards {
		boardDetails0[idx] = NewBoardDetail(each_b, updateNanoTS)
		logrus.Infof("boardDetails0: %v", boardDetails0[idx])
	}

	query0 := &BoardQuery{BBoardID: "1_test1"}
	boardDetail0 := &BoardDetail{
		BBoardID:  "1_test1",
		Brdname:   "test1",
		Title:     "測試1",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "測試",
		BMs:       []bbs.UUserID{"okcool", "teemo"},
		Total:     123,
		NUser:     100,

		LastPostTime: types.NanoTS(1234567890000000000),

		PostType: DEFAULT_POST_TYPE,

		UpdateNanoTS: updateNanoTS,

		Gid: 3,
		Bid: 1,

		IdxByName:  "test1",
		IdxByClass: "tPq41Q@test1",
	}

	type args struct {
		boardDetails []*BoardDetail
		updateNanoTS types.NanoTS
	}
	tests := []struct {
		name    string
		args    args
		query   *BoardQuery
		want    *BoardDetail
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args:  args{boardDetails: boardDetails0, updateNanoTS: updateNanoTS},
			query: query0,
			want:  boardDetail0,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if err := UpdateBoardDetails(tt.args.boardDetails, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("UpdateBoardDetails() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := GetBoardDetail(tt.query.BBoardID, nil)
			if err != nil {
				t.Errorf("UpdateBoardDetails: unable to GetBoardDetail: BBoardID: %v e: %v", tt.query.BBoardID, err)
			}
			testutil.TDeepEqual(t, "got", got, tt.want)
		})
		wg.Wait()
	}
}
