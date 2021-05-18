package api

import (
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/gin-gonic/gin"
)

func TestLoadArticleComments(t *testing.T) {
	setupTest()
	defer teardownTest()

	articleParams := &GetArticleDetailParams{}
	articlePath := &GetArticleDetailPath{
		BBoardID:  bbs.BBoardID("10_WhoAmI"),
		ArticleID: bbs.ArticleID("1VrooM21SYSOP"),
	}
	_, _, _ = GetArticleDetail(testIP, "SYSOP", articleParams, articlePath, nil)

	comments := []*apitypes.Comment{
		{
			BBoardID:   "10_WhoAmI",
			ArticleID:  "1VrooM21SYSOP",
			CommentID:  "FlIk7pJMoAA:cLGi8fC4fapuiBkTXHU2OA",
			TheType:    2,
			CreateTime: 1608388624,
			Owner:      "chhsiao123",
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "噓～",
						Big5:   []byte("\xbcN\xa1\xe3                                                "),
						Color0: types.Color{Foreground: 37, Background: 40},
						Color1: types.Color{Foreground: 37, Background: 40},
						DBCS:   []byte("\xbcN\xa1\xe3                                                "),
					},
				},
			},
		},
		{
			BBoardID:   "10_WhoAmI",
			ArticleID:  "1VrooM21SYSOP",
			CommentID:  "FlIk36uaIAA:FQaNH8WkdAbEGD7yp2Zkvg",
			TheType:    1,
			CreateTime: 1608388560,
			Owner:      "chhsiao123",
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "推",
						Big5:   []byte("\xb1\xc0                                                  "),
						Color0: types.Color{Foreground: 37, Background: 40},
						Color1: types.Color{Foreground: 37, Background: 40},
						DBCS:   []byte("\xb1\xc0                                                  "),
					},
				},
			},
		},
		{
			BBoardID:   "10_WhoAmI",
			ArticleID:  "1VrooM21SYSOP",
			CommentID:  "FlIk0bNSyAA:3dK46zmOe5zmna12AC1gnQ",
			TheType:    3,
			CreateTime: 1608388500,
			Owner:      "SYSOP",
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "推",
						Big5:   []byte("\xb1\xc0                                                       "),
						Color0: types.Color{Foreground: 37, Background: 40},
						Color1: types.Color{Foreground: 37, Background: 40},
						DBCS:   []byte("\xb1\xc0                                                       "),
					},
				},
			},
		},
	}

	params0 := NewLoadArticleCommentsParams()
	path0 := &LoadArticleCommentsPath{
		BBoardID:  "10_WhoAmI",
		ArticleID: "1VrooM21SYSOP",
	}
	expected0 := &LoadArticleCommentsResult{
		List: []*apitypes.Comment{
			comments[0],
			comments[1],
			comments[2],
		},
	}

	params1 := &LoadArticleCommentsParams{
		Descending: true,
		Max:        2,
	}
	expected1 := &LoadArticleCommentsResult{
		List: []*apitypes.Comment{
			comments[0],
			comments[1],
		},
		NextIdx: "1608388500000000000@FlIk0bNSyAA:3dK46zmOe5zmna12AC1gnQ",
	}

	params2 := &LoadArticleCommentsParams{
		StartIdx:   "1608388500000000000@FlIk0bNSyAA:3dK46zmOe5zmna12AC1gnQ",
		Descending: true,
		Max:        2,
	}
	expected2 := &LoadArticleCommentsResult{
		List: []*apitypes.Comment{
			comments[2],
		},
	}

	params3 := &LoadArticleCommentsParams{
		Descending: false,
		Max:        200,
	}
	expected3 := &LoadArticleCommentsResult{
		List: []*apitypes.Comment{
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
			comments[2],
			comments[1],
		},
		NextIdx: "1608388624000000000@FlIk7pJMoAA:cLGi8fC4fapuiBkTXHU2OA",
	}

	params5 := &LoadArticleCommentsParams{
		StartIdx:   "1608388624000000000@FlIk7pJMoAA:cLGi8fC4fapuiBkTXHU2OA",
		Descending: false,
		Max:        2,
	}
	expected5 := &LoadArticleCommentsResult{
		List: []*apitypes.Comment{
			comments[0],
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
	}
}
