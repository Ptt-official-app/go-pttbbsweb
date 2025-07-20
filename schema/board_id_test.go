package schema

import (
	"reflect"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

func TestGetBoardIDByBid(t *testing.T) {
	setupTest()
	defer teardownTest()

	boardSummary0 := &BoardSummary{
		BBoardID: "1_test1",
		Brdname:  "test1",
		Bid:      1,
	}

	_ = UpdateBoardSummaries([]*BoardSummary{boardSummary0}, 1234567890000000000)

	type args struct {
		bid ptttype.Bid
	}
	tests := []struct {
		name            string
		args            args
		expectedBoardID bbs.BBoardID
		wantErr         bool
	}{
		// TODO: Add test cases.
		{
			args:            args{bid: 1},
			expectedBoardID: "1_test1",
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotBoardID, err := GetBoardIDByPttbid(tt.args.bid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBoardIDByBid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBoardID, tt.expectedBoardID) {
				t.Errorf("GetBoardIDByBid() = %v, want %v", gotBoardID, tt.expectedBoardID)
			}
		})
		wg.Wait()
	}
}

func TestGetBoardID(t *testing.T) {
	setupTest()
	defer teardownTest()

	boardSummary0 := &BoardSummary{
		BBoardID: "1_test1",
		Brdname:  "test1",
		Bid:      1,
	}

	_ = UpdateBoardSummaries([]*BoardSummary{boardSummary0}, 1234567890000000000)

	type args struct {
		brdname string
	}
	tests := []struct {
		name            string
		args            args
		expectedBoardID bbs.BBoardID
		wantErr         bool
	}{
		// TODO: Add test cases.
		{
			args:            args{brdname: "test1"},
			expectedBoardID: "1_test1",
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotBoardID, err := GetBoardID(tt.args.brdname)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBoardID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBoardID, tt.expectedBoardID) {
				t.Errorf("GetBoardID() = %v, want %v", gotBoardID, tt.expectedBoardID)
			}
		})
		wg.Wait()
	}
}

func TestGetBoardIDs(t *testing.T) {
	setupTest()
	defer teardownTest()

	boardSummary0 := &BoardSummary{
		BBoardID: "1_test1",
		Brdname:  "test1",
		Bid:      1,
	}
	boardID0 := &BoardID{BBoardID: "1_test1", Brdname: "test1"}

	boardSummary1 := &BoardSummary{
		BBoardID: "2_test2",
		Brdname:  "test2",
		Bid:      2,
	}
	boardID1 := &BoardID{BBoardID: "2_test2", Brdname: "test2"}

	boardSummary2 := &BoardSummary{
		BBoardID: "3_test3",
		Brdname:  "test3",
		Bid:      3,
	}
	boardID2 := &BoardID{BBoardID: "3_test3", Brdname: "test3"}

	boardSummary3 := &BoardSummary{
		BBoardID: "4_test4",
		Brdname:  "test4",
		Bid:      4,
	}
	boardID3 := &BoardID{BBoardID: "4_test4", Brdname: "test4"}

	boardSummary4 := &BoardSummary{
		BBoardID: "5_test5",
		Brdname:  "test5",
		Bid:      5,
	}
	boardID4 := &BoardID{BBoardID: "5_test5", Brdname: "test5"}

	_ = UpdateBoardSummaries([]*BoardSummary{boardSummary0, boardSummary1, boardSummary2, boardSummary3, boardSummary4}, 1234567890000000000)

	type args struct {
		startBrdname string
		descending   bool
		limit        int
		withDeleted  bool
	}
	tests := []struct {
		name           string
		args           args
		expectedResult []*BoardID
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			args:           args{limit: 100},
			expectedResult: []*BoardID{boardID0, boardID1, boardID2, boardID3, boardID4},
		},
		{
			args:           args{limit: 3},
			expectedResult: []*BoardID{boardID0, boardID1, boardID2},
		},
		{
			args:           args{startBrdname: "test3", limit: 2},
			expectedResult: []*BoardID{boardID2, boardID3},
		},
		{
			args:           args{startBrdname: "test4", limit: 3},
			expectedResult: []*BoardID{boardID3, boardID4},
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotResult, err := GetBoardIDs(tt.args.startBrdname, tt.args.descending, tt.args.limit, tt.args.withDeleted)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBoardIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.expectedResult) {
				t.Errorf("GetBoardIDs() = %v, want %v", gotResult, tt.expectedResult)
			}
		})
		wg.Wait()
	}
}
