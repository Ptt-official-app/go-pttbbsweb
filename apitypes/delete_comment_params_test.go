package apitypes

import (
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

func TestDeleteCommentParams_ToComment(t *testing.T) {
	setupTest()
	defer teardownTest()

	comment0 := &schema.Comment{}
	*comment0 = *testComments0[0]

	expected0 := &schema.Comment{
		BBoardID:     bbs.BBoardID("test"),
		ArticleID:    bbs.ArticleID("test"),
		CommentID:    types.CommentID("EYFhmOhREkA:Es26f7U0EXdr7Gp4a9N8pQ"),
		TheType:      ptttype.COMMENT_TYPE_DELETED,
		DeleteReason: "誤植",
		Owner:        bbs.UUserID("foolfighter"),
		CreateTime:   types.NanoTS(1261396680001000000),
		SortTime:     types.NanoTS(1261396680001000000),
		Content: [][]*types.Rune{
			{
				{
					Utf8:   "teemocogs 刪除 foolfighter 的推文: 誤植",
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
					DBCS:   []byte("teemocogs \xa7R\xb0\xa3 foolfighter \xaa\xba\xb1\xc0\xa4\xe5: \xbb~\xb4\xd3"),
				},
			},
		},
		MD5:          "0xGS9JSnmb4xuerU1DQF1w",
		UpdateNanoTS: types.NanoTS(1434567890000000000),
		DBCS:         []byte("\x1b[1;30m(teemocogs \xa7R\xb0\xa3 foolfighter \xaa\xba\xb1\xc0\xa4\xe5: \xbb~\xb4\xd3)\x1b[m\r"),
	}

	type args struct {
		comment      *schema.Comment
		userID       bbs.UUserID
		remoteAddr   string
		updateNanoTS types.NanoTS
	}
	tests := []struct {
		name               string
		d                  *DeleteCommentParams
		args               args
		expectedNewComment *schema.Comment
	}{
		// TODO: Add test cases.
		{
			d:                  &DeleteCommentParams{CommentID: comment0.CommentID, Reason: "誤植"},
			args:               args{comment: comment0, userID: "teemocogs", remoteAddr: "127.0.0.1", updateNanoTS: 1434567890000000000},
			expectedNewComment: expected0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d
			gotNewComment := d.ToComment(tt.args.comment, tt.args.userID, tt.args.remoteAddr, tt.args.updateNanoTS)

			testutil.TDeepEqual(t, "got", gotNewComment, tt.expectedNewComment)
		})
	}
}
