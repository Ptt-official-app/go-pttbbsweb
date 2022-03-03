package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func TestResetArticleIsBottom(t *testing.T) {
	setupTest()
	defer teardownTest()

	articleSummary0 := &ArticleSummaryWithRegex{
		BBoardID:     "board0",
		ArticleID:    "1VtW-QXT",
		CreateTime:   1608388506000000000,
		MTime:        1608388508000000000,
		Recommend:    12,
		Owner:        "SYSOP",
		FullTitle:    "Re:Re:Re: Fw:[轉錄] [爆卦]這裡是八卦板",
		Title:        "這裡是八卦板",
		Class:        "爆卦",
		TitleRegex:   []string{"這", "裡", "是", "八", "卦", "板", "這裡", "裡是", "是八", "八卦", "卦板", "這裡是", "裡是八", "是八卦", "八卦板", "這裡是八", "裡是八卦", "是八卦板", "這裡是八卦", "裡是八卦板"},
		SubjectType:  ptttype.SUBJECT_REPLY,
		Idx:          "1608388506@1VtW-QXT",
		UpdateNanoTS: 1734567890000000000,
		IsBottom:     true,
	}

	expected0 := &ArticleSummary{
		BBoardID:     "board0",
		ArticleID:    "1VtW-QXT",
		CreateTime:   1608388506000000000,
		MTime:        1608388508000000000,
		Recommend:    12,
		Owner:        "SYSOP",
		FullTitle:    "Re:Re:Re: Fw:[轉錄] [爆卦]這裡是八卦板",
		Title:        "這裡是八卦板",
		Class:        "爆卦",
		SubjectType:  ptttype.SUBJECT_REPLY,
		Idx:          "1608388506@1VtW-QXT",
		UpdateNanoTS: 1734567890000000000,
		IsBottom:     true,
	}
	_ = UpdateArticleSummaryWithRegexes([]*ArticleSummaryWithRegex{articleSummary0}, 1234567890000000000)

	type args struct {
		boardID bbs.BBoardID
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool

		before []*ArticleSummary
		after  []*ArticleSummary
	}{
		// TODO: Add test cases.
		{
			args:   args{boardID: "board0"},
			before: []*ArticleSummary{expected0},
			after:  nil,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			got, _ := GetBottomArticleSummaries(tt.args.boardID)
			testutil.TDeepEqual(t, "before-reset", got, tt.before)
			if err := ResetArticleIsBottom(tt.args.boardID); (err != nil) != tt.wantErr {
				t.Errorf("ResetArticleIsBottom() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := GetBottomArticleSummaries(tt.args.boardID)
			if err != nil {
				t.Errorf("ResetArticleIsBottom() after-reset: e: %v", err)
			}
			testutil.TDeepEqual(t, "after-reset", got, tt.after)
		})
		wg.Wait()
	}
}
