package apitypes

import (
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func TestReplyCommentParams_ToComment(t *testing.T) {
	setupTest()
	defer teardownTest()

	replyCommentParams0 := &ReplyCommentParams{
		CommentID: "testCommmentID0",
		Content: [][]*types.Rune{
			{
				{
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
				},
			},
			{
				{
					Utf8:   "test123123",
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
				},
			},
			{
				{
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
				},
			},
			{
				{
					Utf8:   "test124124",
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
				},
			},
			{
				{
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
				},
			},
			{
				{
					Utf8:   "test125125",
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
				},
			},
		},
	}

	expectedContent0 := [][]*types.Rune{
		{
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{
			{
				Utf8:   "test123123",
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("test123123\r"),
			},
		},
		{
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{
			{
				Utf8:   "test124124",
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("test124124\r"),
			},
		},
		{
			{
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\r"),
			},
		},
		{
			{
				Utf8:   "test125125",
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("test125125\r"),
			},
		},
	}

	expected0 := &schema.Comment{
		BBoardID:      "test",
		ArticleID:     "test",
		CommentID:     "testCommmentID0:R",
		TheType:       ptttype.COMMENT_TYPE_REPLY,
		RefIDs:        []types.CommentID{"testCommmentID0"},
		Owner:         "SYSOP",
		Content:       expectedContent0,
		DBCS:          []byte("\r\ntest123123\r\n\r\ntest124124\r\n\r\ntest125125\r"),
		CreateTime:    1334567890000000000,
		MD5:           "UWpO9RU5jvf1CQqXwdNLzg",
		NewCreateTime: 1334567890000000000,
		SortTime:      1234567890000100000,
		EditNanoTS:    1334567890000000000,
		UpdateNanoTS:  1334567890000000000,
	}

	type fields struct {
		CommentID types.CommentID
		Content   [][]*types.Rune
	}
	type args struct {
		userID          bbs.UUserID
		remoteAddr      string
		boardID         bbs.BBoardID
		articleID       bbs.ArticleID
		commentSortTime types.NanoTS
		updateNanoTS    types.NanoTS
	}
	tests := []struct {
		name                 string
		r                    *ReplyCommentParams
		args                 args
		expectedReplyComment *schema.Comment
	}{
		// TODO: Add test cases.
		{
			r:                    replyCommentParams0,
			args:                 args{userID: "SYSOP", boardID: "test", articleID: "test", commentSortTime: 1234567890000000000, updateNanoTS: 1334567890000000000},
			expectedReplyComment: expected0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.r
			gotReplyComment := r.ToComment(tt.args.userID, tt.args.remoteAddr, tt.args.boardID, tt.args.articleID, tt.args.commentSortTime, tt.args.updateNanoTS)

			testutil.TDeepEqual(t, "got", gotReplyComment, tt.expectedReplyComment)
		})
	}
}
