package api

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/gin-gonic/gin"
)

func TestLoadBottomArticles(t *testing.T) {
	setupTest()
	defer teardownTest()

	boardSummaries_b := []*bbs.BoardSummary{testBoardSummaryWhoAmI_b}
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 123456890000000000)

	update0 := &schema.UserReadArticle{UserID: "SYSOP", BoardID: "10_WhoAmI", ArticleID: "19bWBI4Z", UpdateNanoTS: types.Time8(1534567891).ToNanoTS()}
	update1 := &schema.UserReadArticle{UserID: "SYSOP", BoardID: "10_WhoAmI", ArticleID: "1VrooM21", UpdateNanoTS: types.Time8(1234567800).ToNanoTS()}

	_, _ = schema.UserReadArticle_c.Update(update0, update0)
	_, _ = schema.UserReadArticle_c.Update(update1, update1)

	path := &LoadBottomArticlesPath{FBoardID: "WhoAmI"}
	expectedResult := &LoadGeneralArticlesResult{
		List: []*apitypes.ArticleSummary{
			{
				FBoardID:   apitypes.FBoardID("WhoAmI"),
				ArticleID:  apitypes.FArticleID("M.1607937174.A.081"),
				IsDeleted:  false,
				CreateTime: types.Time8(1607937174),
				MTime:      types.Time8(1607937100),
				Recommend:  3,
				Owner:      "teemo",
				Title:      "再來呢？～",
				Class:      "問題",
				Money:      12,
				Filemode:   0,
				URL:        "http://localhost:3457/bbs/WhoAmI/M.1607937174.A.081.html",
				Read:       false,
			},
			{
				FBoardID:   apitypes.FBoardID("WhoAmI"),
				ArticleID:  apitypes.FArticleID("M.1234567890.A.123"),
				IsDeleted:  false,
				CreateTime: types.Time8(1234567890),
				MTime:      types.Time8(1234567889),
				Recommend:  8,
				Owner:      "okcool",
				Title:      "然後呢？～",
				Class:      "問題",
				Money:      3,
				Filemode:   0,
				URL:        "http://localhost:3457/bbs/WhoAmI/M.1234567890.A.123.html",
				Read:       true,
			},
		},
		NextIdx: "1234560000@19bUG021",
	}

	type args struct {
		remoteAddr string
		userID     bbs.UUserID
		params     interface{}
		path       interface{}
		c          *gin.Context
	}
	tests := []struct {
		name               string
		args               args
		expectedResult     interface{}
		expectedStatusCode int
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				remoteAddr: testIP,
				userID:     "SYSOP",
				path:       path,
				c:          &gin.Context{},
			},
			expectedResult:     expectedResult,
			expectedStatusCode: 200,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotResult, gotStatusCode, err := LoadBottomArticles(tt.args.remoteAddr, tt.args.userID, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadBottomArticles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotResult, tt.expectedResult)
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("LoadBottomArticles() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}
