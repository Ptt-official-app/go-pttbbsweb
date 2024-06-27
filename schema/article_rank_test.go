package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

func TestUpdateArticleRank(t *testing.T) {
	setupTest()
	defer teardownTest()

	article := &ArticleSummary{BBoardID: "test_board", ArticleID: "test_article"}
	UpdateArticleSummaries([]*ArticleSummary{article}, 1234567890000000000)

	type args struct {
		boardID      bbs.BBoardID
		articleID    bbs.ArticleID
		diffRank     int
		updateNanoTS types.NanoTS
	}
	tests := []struct {
		name     string
		args     args
		expected int
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{boardID: "test_board", articleID: "test_article", diffRank: 2, updateNanoTS: 1234567890000000000},
			expected: 2,
		},
		{
			args:     args{boardID: "test_board", articleID: "test_article", diffRank: 3, updateNanoTS: 1234567890000000000},
			expected: 5,
		},
		{
			args:     args{boardID: "test_board", articleID: "test_article", diffRank: -1, updateNanoTS: 1234567890000000000},
			expected: 4,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			newRank, err := UpdateArticleRank(tt.args.boardID, tt.args.articleID, tt.args.diffRank, tt.args.updateNanoTS)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateArticleRank() error = %v, wantErr %v", err, tt.wantErr)
			}
			if newRank != tt.expected {
				t.Errorf("UpdateArticleRank: newRank: %v expected: %v", newRank, tt.expected)
			}
			summary, err := GetArticleSummary(tt.args.boardID, tt.args.articleID)
			if err != nil {
				t.Errorf("UpdateArticleRank: unable tget article summary: e: %v", err)
			}
			if summary.Rank != tt.expected {
				t.Errorf("UpdateArticleRank: summary.Rank: %v expected: %v", summary.Rank, tt.expected)
			}
		})
		wg.Wait()
	}
}
