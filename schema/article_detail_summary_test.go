package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

func TestGetArticleDetailSummary(t *testing.T) {
	setupTest()
	defer teardownTest()

	updateNanoTS := types.NanoTS(1234567890000000000)
	articleSummary0 := &ArticleSummary{
		BBoardID:   bbs.BBoardID("board0"),
		ArticleID:  bbs.ArticleID("article0"),
		IsDeleted:  false,
		CreateTime: types.NanoTS(1234567890000000000),
		MTime:      types.NanoTS(1234567889000000000),

		Recommend:    8,
		Owner:        bbs.UUserID("okcool"),
		Title:        "然後呢？～",
		Money:        3,
		Class:        "問題",
		Filemode:     0,
		UpdateNanoTS: updateNanoTS,
	}

	UpdateArticleSummaries([]*ArticleSummary{articleSummary0}, updateNanoTS)

	articleContent := &ArticleContentInfo{
		ContentMD5:          "testmd5",
		IP:                  "127.0.0.1",
		Host:                "localhost",
		BBS:                 "ptt",
		ContentUpdateNanoTS: types.NanoTS(1234567890000000000),
	}

	_ = UpdateArticleContentInfo(bbs.BBoardID("board0"), bbs.ArticleID("article0"), articleContent)

	expected0 := &ArticleDetailSummary{
		BBoardID:   bbs.BBoardID("board0"),
		ArticleID:  bbs.ArticleID("article0"),
		IsDeleted:  false,
		CreateTime: types.NanoTS(1234567890000000000),
		MTime:      types.NanoTS(1234567889000000000),

		Recommend:    8,
		Owner:        bbs.UUserID("okcool"),
		Title:        "然後呢？～",
		Money:        3,
		Class:        "問題",
		Filemode:     0,
		UpdateNanoTS: updateNanoTS,

		ContentMD5:          "testmd5",
		ContentUpdateNanoTS: types.NanoTS(1234567890000000000),
		IP:                  "127.0.0.1",
		Host:                "localhost",
		BBS:                 "ptt",
	}

	type args struct {
		bboardID  bbs.BBoardID
		articleID bbs.ArticleID
	}
	tests := []struct {
		name           string
		args           args
		expectedResult *ArticleDetailSummary
		wantErr        bool
		err            error
	}{
		// TODO: Add test cases.
		{
			args:           args{bboardID: bbs.BBoardID("board0"), articleID: bbs.ArticleID("article0")},
			expectedResult: expected0,
		},
		{
			args: args{bboardID: bbs.BBoardID("board1"), articleID: bbs.ArticleID("article0")},
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotResult, err := GetArticleDetailSummary(tt.args.bboardID, tt.args.articleID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArticleDetailSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != tt.err {
				t.Errorf("GetArticleDetailSummary: e: %v expected: %v", err, tt.err)
			}
			testutil.TDeepEqual(t, "got", gotResult, tt.expectedResult)
		})
		wg.Wait()
	}
}

func TestGetArticleDetailSummaries(t *testing.T) {
	setupTest()
	defer teardownTest()

	updateNanoTS := types.NanoTS(1234567890000000000)
	articleSummary0 := &ArticleSummary{
		BBoardID:   bbs.BBoardID("board0"),
		ArticleID:  bbs.ArticleID("article0"),
		IsDeleted:  false,
		CreateTime: types.NanoTS(1234567890000000000),
		MTime:      types.NanoTS(1234567889000000000),

		Recommend:    8,
		Owner:        bbs.UUserID("okcool"),
		Title:        "然後呢？～",
		Money:        3,
		Class:        "問題",
		Filemode:     0,
		UpdateNanoTS: updateNanoTS,
		Idx:          "1234567890@article0",
	}

	UpdateArticleSummaries([]*ArticleSummary{articleSummary0}, updateNanoTS)

	articleContent0 := &ArticleContentInfo{
		ContentMD5:          "testmd5",
		IP:                  "127.0.0.1",
		Host:                "localhost",
		BBS:                 "ptt",
		ContentUpdateNanoTS: types.NanoTS(1234567890000000000),
	}

	_ = UpdateArticleContentInfo(bbs.BBoardID("board0"), bbs.ArticleID("article0"), articleContent0)

	articleSummary1 := &ArticleSummary{
		BBoardID:   bbs.BBoardID("board0"),
		ArticleID:  bbs.ArticleID("article1"),
		IsDeleted:  false,
		CreateTime: types.NanoTS(1234567891000000000),
		MTime:      types.NanoTS(1234567892000000000),

		Recommend:    5,
		Owner:        bbs.UUserID("okcool2"),
		Title:        "然後呢？～2",
		Money:        6,
		Class:        "test",
		Filemode:     0,
		UpdateNanoTS: updateNanoTS,
		Idx:          "1234567891@article1",
	}

	UpdateArticleSummaries([]*ArticleSummary{articleSummary0, articleSummary1}, updateNanoTS)

	articleContent1 := &ArticleContentInfo{
		ContentMD5:          "testmd5",
		IP:                  "127.0.0.1",
		Host:                "localhost",
		BBS:                 "ptt",
		ContentUpdateNanoTS: types.NanoTS(1234567893000000000),
	}

	_ = UpdateArticleContentInfo(bbs.BBoardID("board0"), bbs.ArticleID("article1"), articleContent1)

	expected0 := &ArticleDetailSummary{
		BBoardID:   bbs.BBoardID("board0"),
		ArticleID:  bbs.ArticleID("article0"),
		IsDeleted:  false,
		CreateTime: types.NanoTS(1234567890000000000),
		MTime:      types.NanoTS(1234567889000000000),

		Recommend:    8,
		Owner:        bbs.UUserID("okcool"),
		Title:        "然後呢？～",
		Money:        3,
		Class:        "問題",
		Filemode:     0,
		UpdateNanoTS: updateNanoTS,

		ContentMD5:          "testmd5",
		ContentUpdateNanoTS: types.NanoTS(1234567890000000000),
		IP:                  "127.0.0.1",
		Host:                "localhost",
		BBS:                 "ptt",

		Idx: "1234567890@article0",
	}

	expected1 := &ArticleDetailSummary{
		BBoardID:   bbs.BBoardID("board0"),
		ArticleID:  bbs.ArticleID("article1"),
		IsDeleted:  false,
		CreateTime: types.NanoTS(1234567891000000000),
		MTime:      types.NanoTS(1234567892000000000),

		Recommend:    5,
		Owner:        bbs.UUserID("okcool2"),
		Title:        "然後呢？～2",
		Money:        6,
		Class:        "test",
		Filemode:     0,
		UpdateNanoTS: updateNanoTS,

		ContentMD5:          "testmd5",
		ContentUpdateNanoTS: types.NanoTS(1234567893000000000),
		IP:                  "127.0.0.1",
		Host:                "localhost",
		BBS:                 "ptt",

		Idx: "1234567891@article1",
	}

	type args struct {
		boardID     bbs.BBoardID
		startIdx    string
		descending  bool
		limit       int
		withDeleted bool
	}
	tests := []struct {
		name           string
		args           args
		expectedResult []*ArticleDetailSummary
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			args:           args{boardID: "board0", limit: 100},
			expectedResult: []*ArticleDetailSummary{expected0, expected1},
		},
		{
			args:           args{boardID: "board0", startIdx: "1234567891@article1", limit: 100},
			expectedResult: []*ArticleDetailSummary{expected1},
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotResult, err := GetArticleDetailSummaries(tt.args.boardID, tt.args.startIdx, tt.args.descending, tt.args.limit, tt.args.withDeleted)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArticleDetailSummaries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotResult, tt.expectedResult)
		})
		wg.Wait()
	}
}
