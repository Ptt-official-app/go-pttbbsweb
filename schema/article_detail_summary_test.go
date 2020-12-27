package schema

import (
	"reflect"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func TestGetArticleDetailSummary(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer Article_c.Drop()

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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := GetArticleDetailSummary(tt.args.bboardID, tt.args.articleID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArticleDetailSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != tt.err {
				t.Errorf("GetArticleDetailSummary: e: %v expected: %v", err, tt.err)
			}
			if !reflect.DeepEqual(gotResult, tt.expectedResult) {
				t.Errorf("GetArticleDetailSummary() = %v, want %v", gotResult, tt.expectedResult)
			}
		})
	}
}
