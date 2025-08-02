package cron

import (
	"sync"
	"testing"
)

func TestLoadFullClassBoards(t *testing.T) {
	setupTest()
	defer teardownTest()
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if err := LoadFullClassBoards(); (err != nil) != tt.wantErr {
				t.Errorf("LoadFullClassBoards() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		wg.Wait()
	}
}

/*
func Test_loadFullClassBoards(t *testing.T) {
	setupTest()
	defer teardownTest()

	expected0 := &schema.BoardSummary{
		BBoardID:  "3_3..........",
		Brdname:   "3..........",
		Title:     "市民廣場     報告  站長  ㄜ！",
		BoardType: "Σ",
		Category:  "....",

		BrdAttr: 8,
		Gid:     1,
		Bid:     3,
		BMs:     []bbs.UUserID{"okcool", "teemo"},

		LastPostTime: 1234567890000000000,

		IdxByName:  "3..........",
		IdxByClass: "....@3..........",
	}

	type args struct {
		startBid ptttype.Bid
	}
	tests := []struct {
		name                 string
		args                 args
		expectedNextBid      ptttype.Bid
		expectedBoardID      bbs.BBoardID
		expectedBoardSummary *schema.BoardSummary
		wantErr              bool
	}{
		// TODO: Add test cases.
		{
			args:                 args{1},
			expectedBoardID:      "3_3..........",
			expectedBoardSummary: expected0,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			_, gotNextBid, err := loadFullClassBoards(tt.args.startBid)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadFullClassBoards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNextBid, tt.expectedNextBid) {
				t.Errorf("loadFullClassBoards() = %v, want %v", gotNextBid, tt.expectedNextBid)
			}

			boardSummary, _ := schema.GetBoardSummary(tt.expectedBoardID)
			boardSummary.UpdateNanoTS = 0
			testutil.TDeepEqual(t, "summary", boardSummary, tt.expectedBoardSummary)
		})
		wg.Wait()
	}
}
*/
