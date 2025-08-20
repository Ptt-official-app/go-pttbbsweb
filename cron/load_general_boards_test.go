package cron

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
)

func Test_loadGeneralBoards(t *testing.T) {
	setupTest()
	defer teardownTest()

	expected0 := []*schema.BoardDetail{
		{
			BBoardID:  "1_test1",
			Brdname:   "test1",
			Title:     "測試1",
			BrdAttr:   0,
			BoardType: "◎",
			Category:  "測試",
			NUser:     100,
			BMs:       []bbs.UUserID{"okcool", "teemo"},
			Total:     123,

			PostType: schema.DEFAULT_POST_TYPE,

			LastPostTime: 1234567890000000000,
			Gid:          3,
			Bid:          1,
			IdxByName:    "test1",
			IdxByClass:   "tPq41Q@test1",
		},
		{
			BBoardID:  "2_test2",
			Brdname:   "test2",
			Title:     "測試2",
			BrdAttr:   0,
			BoardType: "◎",
			Category:  "測試",
			NUser:     101,
			BMs:       []bbs.UUserID{"okcool2", "teemo2"},
			Total:     124,

			PostType: []string{"測試", "◎", ""},

			LastPostTime: 1300000000000000000,
			Gid:          3,
			Bid:          2,
			IdxByName:    "test2",
			IdxByClass:   "tPq41Q@test2",
		},
	}
	type args struct {
		startIdx string
	}
	tests := []struct {
		name                 string
		args                 args
		expectedBoardDetails []*schema.BoardDetail
		expectedNextIdx      string
		wantErr              bool
	}{
		// TODO: Add test cases.
		{
			expectedBoardDetails: expected0,
			expectedNextIdx:      "test3",
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotBoardDetails, gotNextIdx, err := loadGeneralBoards(tt.args.startIdx)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadGeneralBoards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for _, each := range gotBoardDetails {
				each.UpdateNanoTS = 0
			}
			testutil.TDeepEqual(t, "got", gotBoardDetails, tt.expectedBoardDetails)

			if gotNextIdx != tt.expectedNextIdx {
				t.Errorf("nextIdx = %v, want %v", gotNextIdx, tt.expectedNextIdx)
			}
		})
		wg.Wait()
	}
}
