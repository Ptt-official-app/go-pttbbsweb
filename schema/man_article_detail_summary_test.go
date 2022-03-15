package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func TestGetManArticleDetailSummary(t *testing.T) {
	setupTest()
	defer teardownTest()

	a0 := []*ManArticleSummary{
		{
			BBoardID:   "10_WhoAmI",
			LevelIdx:   "D1D2",
			ArticleID:  "D1D2/M.1234567890.A.ABC",
			CreateTime: 1234567890000000000,
			MTime:      1234567890000000000,

			Title:        "這是 title",
			UpdateNanoTS: 1234567890000000000,
			Idx:          0,
		},
		{
			BBoardID:   "10_WhoAmI",
			LevelIdx:   "D1D2",
			ArticleID:  "D1D2/M.1234567890.A.ABD",
			CreateTime: 1234567890000000000,
			MTime:      1234567890000000000,

			Title:        "這是 title-2",
			UpdateNanoTS: 1234567890000000000,
			Idx:          1,
		},
	}
	_ = UpdateManArticleSummaries(a0, 1234567890000000000)

	expected0 := &ManArticleDetailSummary{
		BBoardID:   "10_WhoAmI",
		LevelIdx:   "D1D2",
		ArticleID:  "D1D2/M.1234567890.A.ABC",
		CreateTime: 1234567890000000000,
		MTime:      1234567890000000000,

		Title:        "這是 title",
		UpdateNanoTS: 1234567890000000000,
		Idx:          0,
	}

	expected1 := &ManArticleDetailSummary{
		BBoardID:   "10_WhoAmI",
		LevelIdx:   "D1D2",
		ArticleID:  "D1D2/M.1234567890.A.ABD",
		CreateTime: 1234567890000000000,
		MTime:      1234567890000000000,

		Title:        "這是 title-2",
		UpdateNanoTS: 1234567890000000000,
		Idx:          1,
	}

	type args struct {
		bboardID  bbs.BBoardID
		articleID types.ManArticleID
	}
	tests := []struct {
		name           string
		args           args
		expectedResult *ManArticleDetailSummary
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			args:           args{bboardID: "10_WhoAmI", articleID: "D1D2/M.1234567890.A.ABC"},
			expectedResult: expected0,
		},
		{
			args:           args{bboardID: "10_WhoAmI", articleID: "D1D2/M.1234567890.A.ABD"},
			expectedResult: expected1,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotResult, err := GetManArticleDetailSummary(tt.args.bboardID, tt.args.articleID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetManArticleDetailSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotResult, tt.expectedResult)
		})
		wg.Wait()
	}
}

func TestGetManArticleDetailSummaries(t *testing.T) {
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

	expected0 := []*ManArticleDetailSummary{
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
		boardID  bbs.BBoardID
		levelIdx types.ManArticleID
	}
	tests := []struct {
		name           string
		args           args
		expectedResult []*ManArticleDetailSummary
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			args:           args{boardID: "10_WhoAmI", levelIdx: ""},
			expectedResult: expected0,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotResult, err := GetManArticleDetailSummaries(tt.args.boardID, tt.args.levelIdx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetManArticleDetailSummaries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotResult, tt.expectedResult)
		})
		wg.Wait()
	}
}
