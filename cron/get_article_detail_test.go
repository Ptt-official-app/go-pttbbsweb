package cron

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func TestTryGetArticleContentInfo(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _, _ = loadGeneralArticles("10_WhoAmI", "")
	type args struct {
		boardID   bbs.BBoardID
		articleID bbs.ArticleID
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{boardID: "10_WhoAmI", articleID: "1VrooM21"},
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if err := TryGetArticleContentInfo(tt.args.boardID, tt.args.articleID); (err != nil) != tt.wantErr {
				t.Errorf("TryGetArticleContentInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		wg.Wait()
	}
}
