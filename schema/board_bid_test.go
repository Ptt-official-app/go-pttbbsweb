package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/go-pttbbsweb/mockhttp"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

func TestGetBoardBidsByClsID(t *testing.T) {
	setupTest()
	defer teardownTest()

	ret := mockhttp.LoadGeneralBoards(nil)

	updateNanoTS := types.NowNanoTS()

	boardSummaries0 := make([]*BoardSummary, len(ret.Boards))
	for idx, each_b := range ret.Boards {
		boardSummaries0[idx] = NewBoardSummary(each_b, updateNanoTS)
	}

	_ = UpdateBoardSummaries(boardSummaries0, updateNanoTS)

	boardBid0 := &BoardBid{
		Bid:        1,
		IdxByName:  "test1",
		IdxByClass: "tPq41Q@test1",
	}

	BoardBid1 := &BoardBid{
		Bid:        2,
		IdxByName:  "test2",
		IdxByClass: "tPq41Q@test2",
	}

	expected0 := []*BoardBid{boardBid0, BoardBid1}
	expected1 := []*BoardBid{boardBid0}
	type args struct {
		clsID    ptttype.Bid
		startIdx string
		isAsc    bool
		limit    int
		sortBy   ptttype.BSortBy
	}
	tests := []struct {
		name              string
		args              args
		expectedBoardBids []*BoardBid
		wantErr           bool
	}{
		// TODO: Add test cases.
		{
			args:              args{clsID: 3, limit: 100, sortBy: ptttype.BSORT_BY_CLASS, isAsc: true},
			expectedBoardBids: expected0,
		},
		{
			args:              args{clsID: 3, limit: 100, startIdx: "test1", sortBy: ptttype.BSORT_BY_NAME, isAsc: true},
			expectedBoardBids: expected0,
		},
		{
			args:              args{clsID: 3, limit: 100, startIdx: "test1", sortBy: ptttype.BSORT_BY_NAME, isAsc: false},
			expectedBoardBids: expected1,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotBoardBids, err := GetBoardBidsByClsID(tt.args.clsID, tt.args.startIdx, tt.args.isAsc, tt.args.limit, tt.args.sortBy)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBoardBidsByClsID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotBoardBids, tt.expectedBoardBids)
		})
		wg.Wait()
	}
}
