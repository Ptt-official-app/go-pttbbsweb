package api

import (
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/gin-gonic/gin"
)

func TestCreateComment(t *testing.T) {
	setupTest()
	defer teardownTest()

	params0 := &CreateCommentParams{
		CommentType: ptttype.COMMENT_TYPE_RECOMMEND,
	}
	path0 := &CreateCommentPath{
		BoardID:   "10_WhoAmI",
		ArticleID: "1VrooM21SYSOP",
	}
	expected0 := &apitypes.Comment{
		BBoardID:   "10_WhoAmI",
		ArticleID:  "1VrooM21SYSOP",
		TheType:    ptttype.COMMENT_TYPE_RECOMMEND,
		Owner:      "SYSOP",
		CommentID:  "FoL-3rSuO0A:PlQBjivX5w0hPmg3SIsGjw",
		CreateTime: 1607937324,
		SortTime:   1622139048,
		Content: [][]*types.Rune{
			{
				{
					Utf8:   "test123",
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
					Big5:   []byte("test123                                                  "),
					DBCS:   []byte("test123                                                  "),
				},
			},
		},
		Idx: "1622139048149857000@FofoE7wzVug:PlQBjivX5w0hPmg3SIsGjw",
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
		expectedResult     CreateCommentResult
		expectedStatusCode int
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				remoteAddr: testIP,
				userID:     "SYSOP",
				params:     params0,
				path:       path0,
				c:          nil},
			expectedResult:     expected0,
			expectedStatusCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotStatusCode, err := CreateComment(tt.args.remoteAddr, tt.args.userID, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, _ := gotResult.(CreateCommentResult)
			got.SortTime = tt.expectedResult.SortTime
			got.CommentID = tt.expectedResult.CommentID
			got.Idx = tt.expectedResult.Idx
			testutil.TDeepEqual(t, "got", got, tt.expectedResult)
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("CreateComment() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
	}
}
