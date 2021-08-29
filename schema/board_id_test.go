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
			gotBoardID, err := GetBoardIDByBid(tt.args.bid)
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
