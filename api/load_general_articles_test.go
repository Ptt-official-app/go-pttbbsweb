package api

import (
	"context"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/boardd"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestLoadGeneralArticles(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = deserializeUserDetailAndUpdateDB(testUserSYSOP_b, 123456890000000000)
	_, _ = deserializeUserDetailAndUpdateDB(testUserChhsiao123_b, 123456891000000000)

	boardSummaries_b := []*bbs.BoardSummary{testBoardSummaryWhoAmI_b}
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 123456890000000000)

	update0 := &schema.UserArticle{UserID: "SYSOP", BoardID: "10_WhoAmI", ArticleID: "1VtWRel9", ReadUpdateNanoTS: types.Time8(1608386300).ToNanoTS()}
	update1 := &schema.UserArticle{UserID: "SYSOP", BoardID: "10_WhoAmI", ArticleID: "19bWBI4Z", ReadUpdateNanoTS: types.Time8(1234567990).ToNanoTS()}

	_, _ = schema.UserArticle_c.Update(update0, update0)
	_, _ = schema.UserArticle_c.Update(update1, update1)

	// load articles
	ctx := context.Background()
	brdname := &boardd.BoardRef_Name{Name: "WhoAmI"}
	req := &boardd.ListRequest{
		Ref:          &boardd.BoardRef{Ref: brdname},
		IncludePosts: true,
		Offset:       0,
		Length:       100 + 1,
	}
	resp, _ := boardd.Cli.List(ctx, req)

	posts := resp.Posts

	logrus.Infof("TestLoadGeneralArticles: posts: %v", len(posts))

	updateNanoTS := types.NowNanoTS()
	_, _ = DeserializePBArticlesAndUpdateDB("10_WhoAmI", posts, updateNanoTS, false)

	// params
	params0 := NewLoadGeneralArticlesParams()
	path := &LoadGeneralArticlesPath{FBoardID: "WhoAmI"}
	expected0 := &LoadGeneralArticlesResult{
		List: []*apitypes.ArticleSummary{
			{
				FBoardID:   apitypes.FBoardID("WhoAmI"),
				ArticleID:  apitypes.FArticleID("M.1608388506.A.85D"),
				IsDeleted:  false,
				CreateTime: types.Time8(1608388506),
				MTime:      types.Time8(1608386280),
				Recommend:  9,
				Owner:      "SYSOP",
				Title:      "所以特殊字真的是有綠色的∼",
				Class:      "閒聊",
				Money:      0,
				Filemode:   0,
				URL:        "http://localhost:3457/bbs/board/WhoAmI/article/M.1608388506.A.85D",
				Read:       false,
				Idx:        "1608388506@1VtW-QXT",

				TokenUser: "SYSOP",
				Editable:  true,
				Deletable: true,
			},
			{
				FBoardID:   apitypes.FBoardID("WhoAmI"),
				ArticleID:  apitypes.FArticleID("M.1608386280.A.BC9"),
				IsDeleted:  false,
				CreateTime: types.Time8(1608386280),
				MTime:      types.Time8(1608386280),
				Recommend:  9,
				Owner:      "SYSOP",
				Title:      "測試一下特殊字～",
				Class:      "心得",
				Money:      0,
				Filemode:   0,
				URL:        "http://localhost:3457/bbs/board/WhoAmI/article/M.1608386280.A.BC9",
				Read:       true,
				Idx:        "1608386280@1VtWRel9",

				TokenUser: "SYSOP",
				Editable:  true,
				Deletable: true,
			},
			{
				FBoardID:   apitypes.FBoardID("WhoAmI"),
				ArticleID:  apitypes.FArticleID("M.1607937174.A.081"),
				IsDeleted:  false,
				CreateTime: types.Time8(1607937174),
				MTime:      types.Time8(1607937100),
				Recommend:  3,
				Owner:      "teemo",
				Title:      "新書的情報",
				Class:      "閒聊",
				Money:      0,
				Filemode:   0,
				URL:        "http://localhost:3457/bbs/board/WhoAmI/article/M.1607937174.A.081",
				Read:       false,
				Idx:        "1607937174@1VrooM21",

				TokenUser: "SYSOP",
				Editable:  false,
				Deletable: true,
			},
			{
				FBoardID:   apitypes.FBoardID("WhoAmI"),
				ArticleID:  apitypes.FArticleID("M.1607202240.A.30D"),
				IsDeleted:  false,
				CreateTime: types.Time8(1607202240),
				MTime:      types.Time8(1607202240),
				Recommend:  23,
				Owner:      "cheinshin",
				Title:      "TVBS六都民調 侯奪冠、盧升第四、柯墊底",
				Class:      "新聞",
				Money:      0,
				Filemode:   0,
				URL:        "http://localhost:3457/bbs/board/WhoAmI/article/M.1607202240.A.30D",
				Read:       false,
				Idx:        "1607202240@1Vo_N0CD",

				TokenUser: "SYSOP",
				Editable:  false,
				Deletable: true,
			},
			{
				FBoardID:   apitypes.FBoardID("WhoAmI"),
				ArticleID:  apitypes.FArticleID("M.1584665022.A.ED0"),
				IsDeleted:  false,
				CreateTime: types.Time8(1584665022),
				MTime:      types.Time8(1644506386),
				Recommend:  17,
				Owner:      "hellohiro",
				Title:      "為何打麻將叫賭博但買股票叫投資？",
				Class:      "問卦",
				Money:      0,
				Filemode:   0,
				URL:        "http://localhost:3457/bbs/board/WhoAmI/article/M.1584665022.A.ED0",
				Read:       false,
				Idx:        "1584665022@1UT16-xG",

				TokenUser: "SYSOP",
				Editable:  false,
				Deletable: true,
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
				Money:      0,
				Filemode:   0,
				URL:        "http://localhost:3457/bbs/board/WhoAmI/article/M.1234567890.A.123",
				Read:       true,
				Idx:        "1234567890@19bWBI4Z",

				TokenUser: "SYSOP",
				Editable:  false,
				Deletable: true,
			},
			{
				FBoardID:   apitypes.FBoardID("WhoAmI"),
				ArticleID:  apitypes.FArticleID("M.1234560000.A.081"),
				IsDeleted:  false,
				CreateTime: types.Time8(1234560000),
				MTime:      types.Time8(1234560000),
				Recommend:  13,
				Owner:      "SYSOP",
				Title:      "這是 SYSOP",
				Class:      "問題",
				Money:      0,
				Filemode:   0,
				URL:        "http://localhost:3457/bbs/board/WhoAmI/article/M.1234560000.A.081",
				Read:       false,
				Idx:        "1234560000@19bUG021",

				TokenUser: "SYSOP",
				Editable:  true,
				Deletable: true,
			},
		},
		NextIdx: "",

		TokenUser: "SYSOP",
	}

	params1 := &LoadGeneralArticlesParams{
		Keyword:    "然後",
		Max:        200,
		Descending: true,
	}
	expected1 := &LoadGeneralArticlesResult{
		List: []*apitypes.ArticleSummary{
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
				Money:      0,
				Filemode:   0,
				URL:        "http://localhost:3457/bbs/board/WhoAmI/article/M.1234567890.A.123",
				Read:       true,
				Idx:        "1234567890@19bWBI4Z",

				TokenUser: "SYSOP",
				Deletable: true,
			},
		},

		TokenUser: "SYSOP",
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
			args:               args{remoteAddr: "localhost", userID: "SYSOP", params: params0, path: path, c: &gin.Context{}},
			expectedResult:     expected0,
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: "localhost", userID: "SYSOP", params: params1, path: path, c: &gin.Context{}},
			expectedResult:     expected1,
			expectedStatusCode: 200,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			user := &UserInfo{UserID: tt.args.userID, IsOver18: true}
			gotResult, gotStatusCode, err := LoadGeneralArticles(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadGeneralArticles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			testutil.TDeepEqual(t, "result", gotResult, tt.expectedResult)

			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("LoadGeneralArticles() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}
