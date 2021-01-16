package api

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func TestGetArticleDetail(t *testing.T) {
	setupTest()
	defer teardownTest()

	schema.UserReadArticle_c.Drop()
	schema.Article_c.Drop()
	schema.Comment_c.Drop()

	defer func() {
		schema.UserReadArticle_c.Drop()
		schema.Article_c.Drop()
		schema.Comment_c.Drop()
	}()

	params := &GetArticleDetailParams{}
	path0 := &GetArticleDetailPath{
		BBoardID:  bbs.BBoardID("10_WhoAmI"),
		ArticleID: bbs.ArticleID("19bWBI4ZSYSOP"),
	}

	expectedResult0 := &GetArticleDetailResult{
		BBoardID:   bbs.BBoardID("10_WhoAmI"),
		ArticleID:  bbs.ArticleID("19bWBI4ZSYSOP"),
		Owner:      bbs.UUserID("SYSOP"),
		CreateTime: types.Time8(1234567890),

		URL:  "http://localhost:3457/bbs/10_WhoAmI/M.1234567890.A.123.html",
		Read: true,

		Brdname: "WhoAmI",
		Content: testContent3Utf8,
		IP:      "172.22.0.1",
		Host:    "",
		BBS:     "批踢踢 docker(pttdocker.test)",
	}

	expectedArticleDetailSummary0 := &schema.ArticleDetailSummary{
		BBoardID:              bbs.BBoardID("10_WhoAmI"),
		ArticleID:             bbs.ArticleID("19bWBI4ZSYSOP"),
		ContentMD5:            "2hfH0_ofkV5BgHJMr1H2tg",
		ContentMTime:          types.NanoTS(1607937174000000000),
		FirstCommentsLastTime: types.NanoTS(0),
	}

	path1 := &GetArticleDetailPath{
		BBoardID:  bbs.BBoardID("10_WhoAmI"),
		ArticleID: bbs.ArticleID("1VrooM21SYSOP"),
	}

	expectedResult1 := &GetArticleDetailResult{
		BBoardID:   bbs.BBoardID("10_WhoAmI"),
		ArticleID:  bbs.ArticleID("1VrooM21SYSOP"),
		Owner:      bbs.UUserID("SYSOP"),
		CreateTime: types.Time8(1607937174),

		URL:  "http://localhost:3457/bbs/10_WhoAmI/M.1607937174.A.081.html",
		Read: true,

		Brdname: "WhoAmI",
		Content: testContent4Utf8,
		IP:      "172.22.0.1",
		Host:    "",
		BBS:     "批踢踢 docker(pttdocker.test)",
	}

	expectedArticleDetailSummary1 := &schema.ArticleDetailSummary{
		BBoardID:     bbs.BBoardID("10_WhoAmI"),
		ArticleID:    bbs.ArticleID("1VrooM21SYSOP"),
		ContentMTime: types.NanoTS(1234567890000000000),
		ContentMD5:   "TD1vkp4KtB5bqVEFubzuOw",

		FirstCommentsMD5:      "3fjMk__1yvzpuEgq8jfdmg",
		FirstCommentsLastTime: types.NanoTS(1608388620000000000),
	}

	c := &gin.Context{}
	c.Request = &http.Request{URL: &url.URL{Path: "/api/boards/10_WhoAmI/articles/19bWBI4ZSYSOP"}}
	type args struct {
		remoteAddr string
		userID     bbs.UUserID
		params     interface{}
		path       interface{}
		c          *gin.Context
		boardID    bbs.BBoardID
		articleID  bbs.ArticleID
	}
	tests := []struct {
		name                         string
		args                         args
		expectedResult               interface{}
		expectedFirstComments        []*schema.Comment
		expectedArticleDetailSummary *schema.ArticleDetailSummary

		expectedStatusCode int
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			name: "0th-19bWBI4ZSYSOP",
			args: args{
				remoteAddr: "localhost",
				userID:     "chhsiao123",
				params:     params,
				path:       path0,
				c:          c,
				boardID:    "10_WhoAmI",
				articleID:  "19bWBI4ZSYSOP",
			},
			expectedArticleDetailSummary: expectedArticleDetailSummary0,
			expectedResult:               expectedResult0,
			expectedStatusCode:           200,
		},
		{
			name: "1st-19bWBI4ZSYSOP",
			args: args{
				remoteAddr: "localhost",
				userID:     "chhsiao123",
				params:     params,
				path:       path0,
				c:          c,
				boardID:    "10_WhoAmI",
				articleID:  "19bWBI4ZSYSOP",
			},
			expectedArticleDetailSummary: expectedArticleDetailSummary0,
			expectedResult:               expectedResult0,
			expectedStatusCode:           200,
		},
		{
			name: "0th-1VrooM21SYSOP",
			args: args{
				remoteAddr: "localhost",
				userID:     "chhsiao123",
				params:     params,
				path:       path1,
				c:          c,
				boardID:    "10_WhoAmI",
				articleID:  "1VrooM21SYSOP",
			},
			expectedFirstComments:        testFirstComments4,
			expectedResult:               expectedResult1,
			expectedStatusCode:           200,
			expectedArticleDetailSummary: expectedArticleDetailSummary1,
		},
		{
			name: "1st-1VrooM21SYSOP",
			args: args{
				remoteAddr: "localhost",
				userID:     "chhsiao123",
				params:     params,
				path:       path1,
				c:          c,
				boardID:    "10_WhoAmI",
				articleID:  "1VrooM21SYSOP",
			},
			expectedFirstComments:        testFirstComments4,
			expectedResult:               expectedResult1,
			expectedStatusCode:           200,
			expectedArticleDetailSummary: expectedArticleDetailSummary1,
		},
	}

	for _, tt := range tests {

		gotResult, gotStatusCode, err := GetArticleDetail(tt.args.remoteAddr, tt.args.userID, tt.args.params, tt.args.path, tt.args.c)
		if (err != nil) != tt.wantErr {
			t.Errorf("GetArticleDetail() error = %v, wantErr %v", err, tt.wantErr)
			return
		}

		testutil.TDeepEqual(t, "result", gotResult, tt.expectedResult)

		if gotStatusCode != tt.expectedStatusCode {
			t.Errorf("GetArticleDetail() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
		}

		query := bson.M{
			schema.COMMENT_BBOARD_ID_b:  tt.args.boardID,
			schema.COMMENT_ARTICLE_ID_b: tt.args.articleID,
		}
		var gotComments []*schema.Comment
		_ = schema.Comment_c.Find(query, 0, &gotComments, nil, nil)

		for _, each := range gotComments {
			each.UpdateNanoTS = 0
		}
		testutil.TDeepEqual(t, "comments", gotComments, tt.expectedFirstComments)

		gotArticleDetailSummary, err := schema.GetArticleDetailSummary(tt.args.boardID, tt.args.articleID)
		if gotArticleDetailSummary != nil {
			gotArticleDetailSummary.ContentUpdateNanoTS = 0
			gotArticleDetailSummary.FirstCommentsUpdateNanoTS = 0
		}

		testutil.TDeepEqual(t, "article-detail-summary", gotArticleDetailSummary, tt.expectedArticleDetailSummary)
	}
}
