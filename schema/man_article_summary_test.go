package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/mand"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func TestNewManArticleSummaryFromPB(t *testing.T) {
	entry0 := &mand.Entry{
		Path:  "M.1234567890.A.ABC",
		Title: "這是 title",
		IsDir: false,
	}
	expected0 := &ManArticleSummary{
		BBoardID:   "10_WhoAmI",
		LevelIdx:   "",
		ArticleID:  "M.1234567890.A.ABC",
		CreateTime: 1234567890000000000,
		MTime:      1234567890000000000,

		Title:        "這是 title",
		UpdateNanoTS: 1234567890000000000,
		Idx:          0,
	}
	type args struct {
		entry        *mand.Entry
		boardID      bbs.BBoardID
		levelIdx     types.ManArticleID
		updateNanoTS types.NanoTS
		idx          int
	}
	tests := []struct {
		name                   string
		args                   args
		expectedArticleSummary *ManArticleSummary
	}{
		// TODO: Add test cases.
		{
			args:                   args{entry: entry0, boardID: "10_WhoAmI", levelIdx: "", updateNanoTS: 1234567890000000000},
			expectedArticleSummary: expected0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotArticleSummary := NewManArticleSummaryFromPB(tt.args.entry, tt.args.boardID, tt.args.levelIdx, tt.args.updateNanoTS, tt.args.idx)

			testutil.TDeepEqual(t, "got", gotArticleSummary, tt.expectedArticleSummary)
		})
	}
}

func TestUpdateManArticleSummaries(t *testing.T) {
	setupTest()
	defer teardownTest()

	a0 := []*ManArticleSummary{
		{
			BBoardID:   "10_WhoAmI",
			LevelIdx:   "",
			ArticleID:  "M.1234567890.A.ABC",
			CreateTime: 1234567890000000000,
			MTime:      1234567890000000000,

			Title:        "這是 title",
			UpdateNanoTS: 1234567890000000000,
			Idx:          0,
		},
		{
			BBoardID:   "10_WhoAmI",
			LevelIdx:   "",
			ArticleID:  "M.1234567890.A.ABD",
			CreateTime: 1234567890000000000,
			MTime:      1234567890000000000,

			Title:        "這是 title-2",
			UpdateNanoTS: 1234567890000000000,
			Idx:          1,
		},
	}

	type args struct {
		articleSummaries []*ManArticleSummary
		updateNanoTS     types.NanoTS
	}
	tests := []struct {
		name     string
		args     args
		expected []*ManArticleSummary
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{articleSummaries: a0, updateNanoTS: 1234567890000000000},
			expected: a0,
		},
		{
			args:     args{articleSummaries: a0, updateNanoTS: 1234567890000000000},
			expected: a0,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if err := UpdateManArticleSummaries(tt.args.articleSummaries, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("UpdateManArticleSummaries() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := GetManArticleSummaries(tt.args.articleSummaries[0].BBoardID, tt.args.articleSummaries[0].LevelIdx)
			if err != nil {
				t.Errorf("UpdateManArticleSummaries: unable to get: e: %v", err)
			}
			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
		wg.Wait()
	}
}

func TestRemoveManArticleSummaries(t *testing.T) {
	setupTest()
	defer teardownTest()

	a0 := []*ManArticleSummary{
		{
			BBoardID:   "10_WhoAmI",
			LevelIdx:   "",
			ArticleID:  "M.1234567890.A.ABC",
			CreateTime: 1234567890000000000,
			MTime:      1234567890000000000,

			Title:        "這是 title",
			UpdateNanoTS: 1234567890000000000,
			Idx:          0,
		},
		{
			BBoardID:   "10_WhoAmI",
			LevelIdx:   "",
			ArticleID:  "M.1234567890.A.ABD",
			CreateTime: 1234567890000000000,
			MTime:      1234567890000000000,

			Title:        "這是 title-2",
			UpdateNanoTS: 1234567890000000000,
			Idx:          1,
		},
	}
	_ = UpdateManArticleSummaries(a0, 1234567890000000000)

	type args struct {
		boardID  bbs.BBoardID
		levelIdx types.ManArticleID
		idx      int
	}
	tests := []struct {
		name     string
		args     args
		expected []*ManArticleSummary
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{boardID: "10_WhoAmI", levelIdx: "", idx: 1},
			expected: a0[:1],
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if err := RemoveManArticleSummaries(tt.args.boardID, tt.args.levelIdx, tt.args.idx); (err != nil) != tt.wantErr {
				t.Errorf("RemoveManArticleSummaries() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := GetManArticleSummaries(tt.args.boardID, tt.args.levelIdx)
			if err != nil {
				t.Errorf("RemoveManArticleSummaries: e: %v", err)
			}
			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
		wg.Wait()
	}
}
