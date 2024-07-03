package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

func TestUpdateArticleContentInfo(t *testing.T) {
	setupTest()
	defer teardownTest()

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

			contentInfo, err := GetArticleContentInfo(tt.args.bboardID, tt.args.articleID, true)
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

	contentBlocks := []*ContentBlock{
		{
			BBoardID:     "boardID0",
			ArticleID:    "articleID0",
			ContentID:    "ESIQ9HaNtAA:IKCj3KzpwP5pcJxOAPNDNQ",
			Idx:          0,
			Content:      testContent11Utf8[:50],
			UpdateNanoTS: 1234567890000000000,
		},
		{
			BBoardID:     "boardID0",
			ArticleID:    "articleID0",
			ContentID:    "ESIQ9HaNtAA:IKCj3KzpwP5pcJxOAPNDNQ",
			Idx:          1,
			Content:      testContent11Utf8[50:],
			UpdateNanoTS: 1234567890000000000,
		},
	}

	contentInfo := &ArticleContentInfo{
		ContentMD5: "IKCj3KzpwP5pcJxOAPNDNQ",

		ContentID:           "ESIQ9HaNtAA:IKCj3KzpwP5pcJxOAPNDNQ",
		IP:                  "127.0.0.2",
		Host:                "localhost3",
		BBS:                 "ptt2",
		ContentUpdateNanoTS: types.NanoTS(1234567890000000001),
	}

	expectedContentInfo := &ArticleContentInfo{
		ContentMD5: "IKCj3KzpwP5pcJxOAPNDNQ",

		Content:             testContent11Utf8,
		ContentID:           "ESIQ9HaNtAA:IKCj3KzpwP5pcJxOAPNDNQ",
		IP:                  "127.0.0.2",
		Host:                "localhost3",
		BBS:                 "ptt2",
		ContentUpdateNanoTS: types.NanoTS(1234567890000000001),
	}

	type args struct {
		bboardID  bbs.BBoardID
		articleID bbs.ArticleID
	}
	tests := []struct {
		name                string
		contentInfo         *ArticleContentInfo
		contentBlocks       []*ContentBlock
		args                args
		expectedContentInfo *ArticleContentInfo
		wantErr             bool
	}{
		// TODO: Add test cases.
		{
			args: args{bboardID: bbs.BBoardID("test-non-exists"), articleID: bbs.ArticleID("test-article0")},
		},
		{
			name:                "with content-blocks",
			contentInfo:         contentInfo,
			contentBlocks:       contentBlocks,
			args:                args{bboardID: bbs.BBoardID("boardID0"), articleID: bbs.ArticleID("articleID0")},
			expectedContentInfo: expectedContentInfo,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if tt.contentInfo != nil {
				_ = UpdateArticleContentInfo(tt.args.bboardID, tt.args.articleID, tt.contentInfo)
			}
			if tt.contentBlocks != nil {
				_ = UpdateContentBlocks(tt.contentBlocks, 1234567890000000000)
			}

			gotContentInfo, err := GetArticleContentInfo(tt.args.bboardID, tt.args.articleID, true)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArticleContentInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotContentInfo, tt.expectedContentInfo)
		})
		wg.Wait()
	}
}
