package cron

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func Test_loadGeneralArticles(t *testing.T) {
	setupTest()
	defer teardownTest()

	expected0 := []*schema.ArticleSummaryWithRegex{
		{
			BBoardID:   "10_WhoAmI",
			ArticleID:  "1VrooM21",
			IsDeleted:  false,
			CreateTime: 1607937174000000000,
			MTime:      1607937100000000000,
			Recommend:  3,
			Owner:      "teemo",
			Title:      "再來呢？～",
			Class:      "問題",
			Money:      12,
			Filemode:   0,
			Idx:        "1607937174@1VrooM21",
			FullTitle:  "[問題]再來呢？～",
			TitleRegex: []string{"再", "來", "呢", "？", "～", "再來", "來呢", "呢？", "？～", "再來呢", "來呢？", "呢？～", "再來呢？", "來呢？～", "再來呢？～", "問", "題", "問題"},
		},
		{
			BBoardID:   "10_WhoAmI",
			ArticleID:  "19bWBI4Z",
			IsDeleted:  false,
			CreateTime: 1234567890000000000,
			MTime:      1234567889000000000,
			Recommend:  8,
			Owner:      "okcool",
			Title:      "然後呢？～",
			Class:      "問題",
			Money:      3,
			Filemode:   0,
			Idx:        "1234567890@19bWBI4Z",
			FullTitle:  "[問題]然後呢？～",
			TitleRegex: []string{"然", "後", "呢", "？", "～", "然後", "後呢", "呢？", "？～", "然後呢", "後呢？", "呢？～", "然後呢？", "後呢？～", "然後呢？～", "問", "題", "問題"},
		},
	}

	type args struct {
		boardID  bbs.BBoardID
		startIdx string
	}
	tests := []struct {
		name                     string
		args                     args
		expectedArticleSummaries []*schema.ArticleSummaryWithRegex
		expectedNextIdx          string
		wantErr                  bool
	}{
		// TODO: Add test cases.
		{
			args:                     args{boardID: "10_WhoAmI"},
			expectedArticleSummaries: expected0,
			expectedNextIdx:          "1234560000@19bUG021",
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotArticleSummaries, gotNextIdx, err := loadGeneralArticles(tt.args.boardID, tt.args.startIdx)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadGeneralArticles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, each := range gotArticleSummaries {
				each.UpdateNanoTS = 0
			}
			testutil.TDeepEqual(t, "got", gotArticleSummaries, tt.expectedArticleSummaries)
			if gotNextIdx != tt.expectedNextIdx {
				t.Errorf("loadGeneralArticles() gotNextIdx = %v, want %v", gotNextIdx, tt.expectedNextIdx)
			}
		})
		wg.Wait()
	}
}
