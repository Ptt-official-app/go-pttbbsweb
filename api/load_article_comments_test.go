package api

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/boardd"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestLoadArticleComments(t *testing.T) {
	setupTest()
	defer teardownTest()

	boardSummaries_b := []*bbs.BoardSummary{testBoardSummaryWhoAmI_b}
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 123456890000000000)

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

	// get article detail
	articleParams := &GetArticleDetailParams{}
	articlePath := &GetArticleDetailPath{
		FBoardID:   apitypes.FBoardID("WhoAmI"),
		FArticleID: apitypes.FArticleID("M.1607937174.A.081"),
	}
	_, _, _ = GetArticleDetail(testIP, "SYSOP", articleParams, articlePath, nil)

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
	comments := testUtf8FullAPIComments6

	params0 := NewLoadArticleCommentsParams()
	params0.Descending = false
	path0 := &LoadArticleCommentsPath{
		FBoardID:   "WhoAmI",
		FArticleID: "M.1584665022.A.ED0",
	}
	expected0 := &LoadArticleCommentsResult{
		List: comments,
	}

	params1 := &LoadArticleCommentsParams{
		Descending: true,
		Max:        2,
	}
	expected1 := &LoadArticleCommentsResult{
		List: []*apitypes.Comment{
			comments[20],
			comments[19],
		},
		NextIdx: "1584671580000000000@Ff3iZ_OOmAA:LopnABllsd8x2w5MEenj3w",
	}

	params2 := &LoadArticleCommentsParams{
		StartIdx:   "1584668640000000000@Ff3fu23mwAA:M-lz7qoajllkXbuzxNL2Ww",
		Descending: true,
		Max:        2,
	}
	expected2 := &LoadArticleCommentsResult{
		List: []*apitypes.Comment{
			comments[0],
		},
	}

	params3 := &LoadArticleCommentsParams{
		Descending: true,
		Max:        200,
	}
	expected3 := &LoadArticleCommentsResult{
		List: []*apitypes.Comment{
			comments[20],
			comments[19],
			comments[18],
			comments[17],
			comments[16],
			comments[15],
			comments[14],
			comments[13],
			comments[12],
			comments[11],
			comments[10],
			comments[9],
			comments[8],
			comments[7],
			comments[6],
			comments[5],
			comments[4],
			comments[3],
			comments[2],
			comments[1],
			comments[0],
		},
	}

	params4 := &LoadArticleCommentsParams{
		Descending: false,
		Max:        2,
	}
	expected4 := &LoadArticleCommentsResult{
		List: []*apitypes.Comment{
			comments[0],
			comments[1],
		},
		NextIdx: "1584668700001000000@Ff3fyWY9WkA:8M52DhmjUJDOUbeiKRdKwg",
	}

	params5 := &LoadArticleCommentsParams{
		StartIdx:   "1644506386000000000@FtJ12FhkNAA:4oMXnXegENBL9Dn6-SddrA",
		Descending: false,
		Max:        2,
	}
	expected5 := &LoadArticleCommentsResult{
		List: []*apitypes.Comment{
			comments[20],
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
		expectedResult     *LoadArticleCommentsResult
		expectedStatusCode int
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: params0, path: path0},
			expectedResult:     expected0,
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: params1, path: path0},
			expectedResult:     expected1,
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: params2, path: path0},
			expectedResult:     expected2,
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: params3, path: path0},
			expectedResult:     expected3,
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: params4, path: path0},
			expectedResult:     expected4,
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: params5, path: path0},
			expectedResult:     expected5,
			expectedStatusCode: 200,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotResult, gotStatusCode, err := LoadArticleComments(tt.args.remoteAddr, tt.args.userID, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadArticleComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			testutil.TDeepEqual(t, "got", gotResult, tt.expectedResult)

			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("LoadArticleComments() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}
