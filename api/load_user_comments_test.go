package api

import (
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/gin-gonic/gin"
)

func TestLoadUserComments(t *testing.T) {
	setupTest()
	defer teardownTest()

	// board
	boardSummaries_b := []*bbs.BoardSummary{testBoardSummaryWhoAmI_b, testBoardSummarySYSOP_b}
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 123456890000000000)

	// articles
	update0 := &schema.UserReadArticle{UserID: "SYSOP", BoardID: "10_WhoAmI", ArticleID: "19bWBI4Z", UpdateNanoTS: types.Time8(1534567891).ToNanoTS()}
	update1 := &schema.UserReadArticle{UserID: "SYSOP", BoardID: "10_WhoAmI", ArticleID: "1VrooM21", UpdateNanoTS: types.Time8(1234567800).ToNanoTS()}

	_, _ = schema.UserReadArticle_c.Update(update0, update0)
	_, _ = schema.UserReadArticle_c.Update(update1, update1)

	paramsLoadGeneralArticles := NewLoadGeneralArticlesParams()
	pathLoadGeneralArticles := &LoadGeneralArticlesPath{FBoardID: "WhoAmI"}
	LoadGeneralArticles("localhost", "SYSOP", paramsLoadGeneralArticles, pathLoadGeneralArticles, &gin.Context{})

	paramsLoadGeneralArticles = NewLoadGeneralArticlesParams()
	pathLoadGeneralArticles = &LoadGeneralArticlesPath{FBoardID: "SYSOP"}
	LoadGeneralArticles("localhost", "SYSOP", paramsLoadGeneralArticles, pathLoadGeneralArticles, &gin.Context{})

	articleParams := &GetArticleDetailParams{}
	articlePath := &GetArticleDetailPath{
		FBoardID:   apitypes.FBoardID("WhoAmI"),
		FArticleID: apitypes.FArticleID("M.1607937174.A.081"),
	}
	_, _, _ = GetArticleDetail(testIP, "SYSOP", articleParams, articlePath, nil)

	// tests
	params0 := NewLoadUserCommentsParams()
	path0 := &LoadUserCommentsPath{
		UserID: "chhsiao123",
	}

	expected0 := &LoadUserCommentsResult{
		List: []*apitypes.ArticleComment{
			{
				FBoardID:          "WhoAmI",
				FArticleID:        "M.1607937174.A.081",
				CreateTime:        1607937174,
				MTime:             1607937100,
				Recommend:         3,
				NComments:         3,
				Owner:             "teemo",
				Title:             "再來呢？～",
				Money:             12,
				Class:             "問題",
				URL:               "http://localhost:3457/bbs/WhoAmI/M.1607937174.A.081.html",
				Read:              types.READ_STATUS_MTIME,
				Idx:               "c#1608388624000000000@FlIk7pJMoAA:cLGi8fC4fapuiBkTXHU2OA",
				TheType:           apitypes.ARTICLE_COMMENT_TYPE_COMMENT,
				CommentID:         "FlIk7pJMoAA:cLGi8fC4fapuiBkTXHU2OA",
				CommentType:       ptttype.COMMENT_TYPE_BOO,
				CommentCreateTime: 1608388624,
				Comment: [][]*types.Rune{
					{
						{
							Utf8:   "噓～",
							Color0: types.DefaultColor,
							Color1: types.DefaultColor,
							Big5:   []byte("\xbcN\xa1\xe3                                                "),
							DBCS:   []byte("\xbcN\xa1\xe3                                                "),
						},
					},
				},
			},
			{
				FBoardID:          "WhoAmI",
				FArticleID:        "M.1607937174.A.081",
				CreateTime:        1607937174,
				MTime:             1607937100,
				Recommend:         3,
				NComments:         3,
				Owner:             "teemo",
				Title:             "再來呢？～",
				Money:             12,
				Class:             "問題",
				URL:               "http://localhost:3457/bbs/WhoAmI/M.1607937174.A.081.html",
				Read:              types.READ_STATUS_MTIME,
				Idx:               "c#1608388560000000000@FlIk36uaIAA:FQaNH8WkdAbEGD7yp2Zkvg",
				TheType:           apitypes.ARTICLE_COMMENT_TYPE_COMMENT,
				CommentID:         "FlIk36uaIAA:FQaNH8WkdAbEGD7yp2Zkvg",
				CommentType:       ptttype.COMMENT_TYPE_RECOMMEND,
				CommentCreateTime: 1608388560,
				Comment: [][]*types.Rune{
					{
						{
							Utf8:   "推",
							Color0: types.DefaultColor,
							Color1: types.DefaultColor,
							Big5:   []byte("\xb1\xc0                                                  "),
							DBCS:   []byte("\xb1\xc0                                                  "),
						},
					},
				},
			},
		},
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
			args:               args{remoteAddr: "localhost", userID: "SYSOP", params: params0, path: path0},
			expectedResult:     expected0,
			expectedStatusCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotStatusCode, err := LoadUserComments(tt.args.remoteAddr, tt.args.userID, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadUserComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotResult, tt.expectedResult)
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("LoadUserComments() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
	}
}
