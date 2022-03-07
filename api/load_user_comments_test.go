package api

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/boardd"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	logrus.Infof("TestLoadUserArticles: posts: %v", len(posts))

	updateNanoTS := types.NowNanoTS()
	_, _ = DeserializePBArticlesAndUpdateDB("10_WhoAmI", posts, updateNanoTS, false)

	ctx = context.Background()
	brdname = &boardd.BoardRef_Name{Name: "SYSOP"}
	req = &boardd.ListRequest{
		Ref:          &boardd.BoardRef{Ref: brdname},
		IncludePosts: true,
		Offset:       0,
		Length:       100 + 1,
	}
	resp, _ = boardd.Cli.List(ctx, req)

	posts = resp.Posts

	logrus.Infof("TestLoadUserArticles: posts: %v", len(posts))

	updateNanoTS = types.NowNanoTS()
	_, _ = DeserializePBArticlesAndUpdateDB("1_SYSOP", posts, updateNanoTS, false)

	// get article detail
	articleParams := &GetArticleDetailParams{}
	articlePath0 := &GetArticleDetailPath{
		FBoardID:   apitypes.FBoardID("WhoAmI"),
		FArticleID: apitypes.FArticleID("M.1607937174.A.081"),
	}
	_, _, _ = GetArticleDetail(testIP, "SYSOP", articleParams, articlePath0, nil)

	articleParams2 := &GetArticleDetailParams{}
	articlePath2 := &GetArticleDetailPath{
		FBoardID:   apitypes.FBoardID("WhoAmI"),
		FArticleID: apitypes.FArticleID("M.1608388506.A.85D"),
	}
	_, _, _ = GetArticleDetail(testIP, "SYSOP", articleParams2, articlePath2, nil)

	articleParams3 := &GetArticleDetailParams{}
	articlePath3 := &GetArticleDetailPath{
		FBoardID:   apitypes.FBoardID("WhoAmI"),
		FArticleID: apitypes.FArticleID("M.1584665022.A.ED0"),
	}
	_, _, _ = GetArticleDetail(testIP, "SYSOP", articleParams3, articlePath3, nil)

	time.Sleep(3 * time.Second)

	logrus.Infof("TestLoadArticleComments: get article detail: after sleep")

	// params
	paramsLoadGeneralArticles := NewLoadGeneralArticlesParams()
	pathLoadGeneralArticles := &LoadGeneralArticlesPath{FBoardID: "WhoAmI"}
	LoadGeneralArticles("localhost", "SYSOP", paramsLoadGeneralArticles, pathLoadGeneralArticles, &gin.Context{})

	paramsLoadGeneralArticles = NewLoadGeneralArticlesParams()
	pathLoadGeneralArticles = &LoadGeneralArticlesPath{FBoardID: "SYSOP"}
	LoadGeneralArticles("localhost", "SYSOP", paramsLoadGeneralArticles, pathLoadGeneralArticles, &gin.Context{})

	// tests
	params0 := NewLoadUserCommentsParams()
	path0 := &LoadUserCommentsPath{
		UserID: "kumori",
	}

	expected0 := &LoadUserCommentsResult{
		List: []*apitypes.ArticleComment{
			{ // 0
				FBoardID:          "WhoAmI",
				FArticleID:        "M.1584665022.A.ED0",
				CreateTime:        1584665022,
				MTime:             1644506386,
				Recommend:         17,
				NComments:         21,
				Owner:             "hellohiro",
				Title:             "為何打麻將叫賭博但買股票叫投資？",
				Money:             0,
				Class:             "問卦",
				URL:               "http://localhost:3457/bbs/board/WhoAmI/article/M.1584665022.A.ED0",
				Read:              types.READ_STATUS_MTIME,
				Idx:               "c#1584669120000000000@Ff3gKzAhgAA:2wXn35Dn9hRYtaN-RZF3GA",
				TheType:           apitypes.ARTICLE_COMMENT_TYPE_COMMENT,
				CommentID:         "Ff3gKzAhgAA:2wXn35Dn9hRYtaN-RZF3GA",
				CommentType:       ptttype.COMMENT_TYPE_COMMENT,
				CommentCreateTime: 1584669120,
				Comment: [][]*types.Rune{
					{
						{
							Utf8:    "拿回",
							Color0:  types.DefaultColor,
							Color1:  types.DefaultColor,
							DBCSStr: "拿回",
						},
					},
				},
			},
			{ // 1
				FBoardID:          "WhoAmI",
				FArticleID:        "M.1584665022.A.ED0",
				CreateTime:        1584665022,
				MTime:             1644506386,
				Recommend:         17,
				NComments:         21,
				Owner:             "hellohiro",
				Title:             "為何打麻將叫賭博但買股票叫投資？",
				Money:             0,
				Class:             "問卦",
				URL:               "http://localhost:3457/bbs/board/WhoAmI/article/M.1584665022.A.ED0",
				Read:              types.READ_STATUS_MTIME,
				Idx:               "c#1584669060001000000@Ff3gHTfpakA:pSRX0Ox7KFM1nip7dmdIeg",
				TheType:           apitypes.ARTICLE_COMMENT_TYPE_COMMENT,
				CommentID:         "Ff3gHTfpakA:pSRX0Ox7KFM1nip7dmdIeg",
				CommentType:       ptttype.COMMENT_TYPE_COMMENT,
				CommentCreateTime: 1584669060,
				Comment: [][]*types.Rune{
					{
						{
							Utf8:    "得到本金 以及利息",
							Color0:  types.DefaultColor,
							Color1:  types.DefaultColor,
							DBCSStr: "得到本金 以及利息",
						},
					},
				},
			},
			{ // 2
				FBoardID:          "WhoAmI",
				FArticleID:        "M.1584665022.A.ED0",
				CreateTime:        1584665022,
				MTime:             1644506386,
				Recommend:         17,
				NComments:         21,
				Owner:             "hellohiro",
				Title:             "為何打麻將叫賭博但買股票叫投資？",
				Money:             0,
				Class:             "問卦",
				URL:               "http://localhost:3457/bbs/board/WhoAmI/article/M.1584665022.A.ED0",
				Read:              types.READ_STATUS_MTIME,
				Idx:               "c#1584669000002000000@Ff3gDz-xVIA:am4iq8LtZTp8p_GYd_G9-Q",
				TheType:           apitypes.ARTICLE_COMMENT_TYPE_COMMENT,
				CommentID:         "Ff3gDz-xVIA:am4iq8LtZTp8p_GYd_G9-Q",
				CommentType:       ptttype.COMMENT_TYPE_COMMENT,
				CommentCreateTime: 1584669000,
				Comment: [][]*types.Rune{
					{
						{
							Utf8:    "賺 於是你借他錢 後來他真的賺了 你就可以",
							Color0:  types.DefaultColor,
							Color1:  types.DefaultColor,
							DBCSStr: "賺 於是你借他錢 後來他真的賺了 你就可以",
						},
					},
				},
			},
			{ // 3
				FBoardID:          "WhoAmI",
				FArticleID:        "M.1584665022.A.ED0",
				CreateTime:        1584665022,
				MTime:             1644506386,
				Recommend:         17,
				NComments:         21,
				Owner:             "hellohiro",
				Title:             "為何打麻將叫賭博但買股票叫投資？",
				Money:             0,
				Class:             "問卦",
				URL:               "http://localhost:3457/bbs/board/WhoAmI/article/M.1584665022.A.ED0",
				Read:              types.READ_STATUS_MTIME,
				Idx:               "c#1584668940001000000@Ff3gAUdaukA:vgSdv2gWw4H7hFaJF_An7A",
				TheType:           apitypes.ARTICLE_COMMENT_TYPE_COMMENT,
				CommentID:         "Ff3gAUdaukA:vgSdv2gWw4H7hFaJF_An7A",
				CommentType:       ptttype.COMMENT_TYPE_COMMENT,
				CommentCreateTime: 1584668940,
				Comment: [][]*types.Rune{
					{
						{
							Utf8:    "發行一張紙跟你借錢 你認為這間公司以後會",
							Color0:  types.DefaultColor,
							Color1:  types.DefaultColor,
							DBCSStr: "發行一張紙跟你借錢 你認為這間公司以後會",
						},
					},
				},
			},
			{ // 4
				FBoardID:          "WhoAmI",
				FArticleID:        "M.1584665022.A.ED0",
				CreateTime:        1584665022,
				MTime:             1644506386,
				Recommend:         17,
				NComments:         21,
				Owner:             "hellohiro",
				Title:             "為何打麻將叫賭博但買股票叫投資？",
				Money:             0,
				Class:             "問卦",
				URL:               "http://localhost:3457/bbs/board/WhoAmI/article/M.1584665022.A.ED0",
				Read:              types.READ_STATUS_MTIME,
				Idx:               "c#1584668880000000000@Ff3f808EIAA:gPPKXXc-6zGdsceRFTufkg",
				TheType:           apitypes.ARTICLE_COMMENT_TYPE_COMMENT,
				CommentID:         "Ff3f808EIAA:gPPKXXc-6zGdsceRFTufkg",
				CommentType:       ptttype.COMMENT_TYPE_COMMENT,
				CommentCreateTime: 1584668880,
				Comment: [][]*types.Rune{
					{
						{
							Utf8:    "股票的意義是 一間公司想生產東西現在沒錢",
							Color0:  types.DefaultColor,
							Color1:  types.DefaultColor,
							DBCSStr: "股票的意義是 一間公司想生產東西現在沒錢",
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
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
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
		wg.Wait()
	}
}
