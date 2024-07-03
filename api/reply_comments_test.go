package api

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/go-pttbbsweb/apitypes"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestReplyComments(t *testing.T) {
	setupTest()
	defer teardownTest()

	boardSummaries_b := []*bbs.BoardSummary{testBoardSummarySYSOP_b}
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 123456890000000000)

	createArticlePath0 := &CreateArticlePath{
		FBoardID: "SYSOP",
	}

	createArticleRune0 := &types.Rune{
		Utf8:   "測試0",
		Color0: types.DefaultColor,
		Color1: types.DefaultColor,
	}

	createArticleRune1 := &types.Rune{
		Utf8:   "測試1",
		Color0: types.DefaultColor,
		Color1: types.DefaultColor,
	}

	createArticleParams0 := &CreateArticleParams{
		PostType: "測試",
		Title:    "this is a test",
		Content: [][]*types.Rune{
			{createArticleRune0, createArticleRune1},
		},
	}

	gotCreateArticleResult0, _, _ := CreateArticle(testIP, "SYSOP", createArticleParams0, createArticlePath0, nil)
	gotCreateArticle0, _ := gotCreateArticleResult0.(CreateArticleResult)

	createCommentParams0 := &CreateCommentParams{
		CommentType: ptttype.COMMENT_TYPE_RECOMMEND,
		Content:     "test123",
	}
	createCommentPath0 := &CreateCommentPath{
		FBoardID:   gotCreateArticle0.FBoardID,
		FArticleID: gotCreateArticle0.ArticleID,
	}

	gotCreateCommentResult0, _, err := CreateComment(testIP, "SYSOP", createCommentParams0, createCommentPath0, nil)
	createCommentResult0, _ := gotCreateCommentResult0.(CreateCommentResult)
	logrus.Infof("createCommentResult: %v gotCreateCommentResult0: %v e: %v", createCommentResult0, gotCreateCommentResult0, err)

	content0 := [][]*types.Rune{
		{
			{
				Utf8:   "測試2",
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{
			{
				Utf8:   "測試3",
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
	}

	params0 := &ReplyCommentsParams{
		TheList: []*apitypes.ReplyCommentParams{
			{
				CommentID: createCommentResult0.CommentID,
				Content:   content0,
			},
		},
	}

	path0 := &ReplyCommentsPath{
		FBoardID:   gotCreateArticle0.FBoardID,
		FArticleID: gotCreateArticle0.ArticleID,
	}

	result0 := &ReplyCommentsResult{
		Success: true,

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
		expectedResult     *ReplyCommentsResult
		expectedStatusCode int
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: params0, path: path0},
			expectedResult:     result0,
			expectedStatusCode: 200,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotResult, gotStatusCode, err := ReplyComments(tt.args.remoteAddr, tt.args.userID, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReplyComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotResult, tt.expectedResult)

			testutil.TDeepEqual(t, "got", gotStatusCode, tt.expectedStatusCode)
		})
		wg.Wait()
	}
}
