package schema

import (
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func TestUpdateArticleFirstComments(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer Article_c.Drop()

	articleContent := &ArticleContentInfo{
		ContentMD5:          "test1",
		IP:                  "127.0.0.1",
		Host:                "localhost",
		BBS:                 "ptt",
		ContentUpdateNanoTS: types.NanoTS(1234567890000000000),
	}

	_ = UpdateArticleContentInfo(bbs.BBoardID("board0"), bbs.ArticleID("article0"), articleContent)

	firstComments0 := &ArticleFirstComments{
		BBoardID:  bbs.BBoardID("board0"),
		ArticleID: bbs.ArticleID("article0"),

		FirstCommentsMD5:          "testmd5",
		FirstCommentsLastTime:     types.NanoTS(1234567890000000000),
		FirstCommentsUpdateNanoTS: types.NanoTS(1234567890000000000),
	}

	firstComments1 := &ArticleFirstComments{
		BBoardID:  bbs.BBoardID("board0"),
		ArticleID: bbs.ArticleID("article1"),

		FirstCommentsMD5:          "testmd5",
		FirstCommentsLastTime:     types.NanoTS(1234567890000000000),
		FirstCommentsUpdateNanoTS: types.NanoTS(1234567890000000000),
	}

	firstComments2 := &ArticleFirstComments{
		BBoardID:  bbs.BBoardID("board0"),
		ArticleID: bbs.ArticleID("article0"),

		FirstCommentsMD5:          "testmd5",
		FirstCommentsLastTime:     types.NanoTS(1234567890000000000),
		FirstCommentsUpdateNanoTS: types.NanoTS(1234567890000000001),
	}

	type args struct {
		articleFirstComments *ArticleFirstComments
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		// TODO: Add test cases.
		{
			args: args{articleFirstComments: firstComments0},
		},
		{
			args:    args{articleFirstComments: firstComments0},
			wantErr: true,
			err:     ErrNoMatch,
		},
		{
			args:    args{articleFirstComments: firstComments1},
			wantErr: true,
			err:     ErrNoMatch,
		},
		{
			args: args{articleFirstComments: firstComments2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UpdateArticleFirstComments(tt.args.articleFirstComments)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateArticleFirstComments() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != tt.err {
				t.Errorf("UpdateArticleFirstComments: e: %v expeted: %v", err, tt.err)
			}
		})
	}
}
