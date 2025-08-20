package schema

import (
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func TestUpdateArticleContentMTime(t *testing.T) {
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

	articleContentMTime := &ArticleContentMTime{
		ContentMTime: types.NanoTS(1234567890000000000),
	}

	type args struct {
		bboardID     bbs.BBoardID
		articleID    bbs.ArticleID
		contentMTime types.NanoTS
	}
	tests := []struct {
		name     string
		args     args
		expected *ArticleContentMTime
		wantErr  bool
		err      error
	}{
		// TODO: Add test cases.
		{
			args: args{
				bboardID:     bbs.BBoardID("board1"),
				articleID:    bbs.ArticleID("article1"),
				contentMTime: types.NanoTS(1234567890000000000),
			},
			wantErr: true,
			err:     ErrNoMatch,
		},
		{
			args: args{
				bboardID:     bbs.BBoardID("board0"),
				articleID:    bbs.ArticleID("article0"),
				contentMTime: types.NanoTS(1234567890000000000),
			},
			expected: articleContentMTime,
			wantErr:  false,
		},
		{
			args: args{
				bboardID:     bbs.BBoardID("board0"),
				articleID:    bbs.ArticleID("article0"),
				contentMTime: types.NanoTS(1234567890000000000),
			},
			wantErr:  true,
			err:      ErrNoMatch,
			expected: articleContentMTime,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UpdateArticleContentMTime(tt.args.bboardID, tt.args.articleID, tt.args.contentMTime)

			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateArticleContentMTime() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != tt.err {
				t.Errorf("UpdateArticleContentMTime: e: %v expected: %v", err, tt.err)
			}

			got, _ := GetArticleContentMTime(tt.args.bboardID, tt.args.articleID)

			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
	}
}
