package schema

import (
	"reflect"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func TestUpdateArticleContentInfo(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer Article_c.Drop()

	contentInfo0 := &ArticleContentInfo{
		ContentMD5:          "test1",
		IP:                  "127.0.0.1",
		Host:                "localhost",
		BBS:                 "ptt",
		ContentUpdateNanoTS: types.NanoTS(1234567890000000000),
	}

	contentInfo1 := &ArticleContentInfo{
		ContentMD5:          "test2",
		IP:                  "127.0.0.2",
		Host:                "localhost2",
		BBS:                 "ptt2",
		ContentUpdateNanoTS: types.NanoTS(1234567880000000000),
	}

	contentInfo2 := &ArticleContentInfo{
		ContentMD5:          "test3",
		IP:                  "127.0.0.2",
		Host:                "localhost3",
		BBS:                 "ptt2",
		ContentUpdateNanoTS: types.NanoTS(1234567890000000001),
	}

	type args struct {
		bboardID    bbs.BBoardID
		articleID   bbs.ArticleID
		contentInfo *ArticleContentInfo
	}
	tests := []struct {
		name     string
		args     args
		expected *ArticleContentInfo
		wantErr  bool
		err      error
	}{
		// TODO: Add test cases.
		{
			args:     args{bboardID: bbs.BBoardID("board0"), articleID: bbs.ArticleID("article0"), contentInfo: contentInfo0},
			expected: contentInfo0,
		},
		{
			args:     args{bboardID: bbs.BBoardID("board0"), articleID: bbs.ArticleID("article0"), contentInfo: contentInfo1},
			expected: contentInfo0,
			wantErr:  true,
			err:      ErrNoMatch,
		},
		{
			args:     args{bboardID: bbs.BBoardID("board0"), articleID: bbs.ArticleID("article0"), contentInfo: contentInfo2},
			expected: contentInfo2,
		},
		{
			args:     args{bboardID: bbs.BBoardID("board0"), articleID: bbs.ArticleID("article1"), contentInfo: contentInfo2},
			expected: contentInfo2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UpdateArticleContentInfo(tt.args.bboardID, tt.args.articleID, tt.args.contentInfo)

			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateArticleContentInfo() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != tt.err {
				t.Errorf("UpdateArticleContentInfo: err: %v expected: %v", err, tt.err)
			}

			contentInfo, err := GetArticleContentInfo(tt.args.bboardID, tt.args.articleID)
			if err != nil {
				t.Errorf("UpdateArticleContentInfo: unable to get: e: %v", err)
			}

			testutil.TDeepEqual(t, "contentInfo", contentInfo, tt.expected)
		})
	}
}

func TestGetArticleContentInfo(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer Article_c.Drop()

	type args struct {
		bboardID  bbs.BBoardID
		articleID bbs.ArticleID
	}
	tests := []struct {
		name                string
		args                args
		expectedContentInfo *ArticleContentInfo
		wantErr             bool
	}{
		// TODO: Add test cases.
		{
			args: args{bboardID: bbs.BBoardID("test-non-exists"), articleID: bbs.ArticleID("test-article0")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotContentInfo, err := GetArticleContentInfo(tt.args.bboardID, tt.args.articleID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArticleContentInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotContentInfo, tt.expectedContentInfo) {
				t.Errorf("GetArticleContentInfo() = %v, want %v", gotContentInfo, tt.expectedContentInfo)
			}
		})
	}
}
