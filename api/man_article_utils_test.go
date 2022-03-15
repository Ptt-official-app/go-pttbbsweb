package api

import (
	"reflect"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/mand"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func TestDeserializePBManArticlesAndUpdateDB(t *testing.T) {
	type args struct {
		boardID      bbs.BBoardID
		parentID     types.ManArticleID
		entries      []*mand.Entry
		updateNanoTS types.NanoTS
	}
	tests := []struct {
		name                     string
		args                     args
		expectedArticleSummaries []*schema.ManArticleSummary
		wantErr                  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotArticleSummaries, err := DeserializePBManArticlesAndUpdateDB(tt.args.boardID, tt.args.parentID, tt.args.entries, tt.args.updateNanoTS)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeserializePBManArticlesAndUpdateDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotArticleSummaries, tt.expectedArticleSummaries) {
				t.Errorf("DeserializePBManArticlesAndUpdateDB() = %v, want %v", gotArticleSummaries, tt.expectedArticleSummaries)
			}
		})
	}
}

func TestUpdateManArticleContentInfo(t *testing.T) {
	type args struct {
		boardID      bbs.BBoardID
		articleID    types.ManArticleID
		content      [][]*types.Rune
		contentMD5   string
		updateNanoTS types.NanoTS
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateManArticleContentInfo(tt.args.boardID, tt.args.articleID, tt.args.content, tt.args.contentMD5, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("UpdateManArticleContentInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTryGetManArticleContentInfo(t *testing.T) {
	type args struct {
		userID    bbs.UUserID
		bboardID  bbs.BBoardID
		articleID types.ManArticleID
		c         *gin.Context
		isSystem  bool
		isContent bool
	}
	tests := []struct {
		name                         string
		args                         args
		expectedContent              [][]*types.Rune
		expectedContentMD5           string
		expectedArticleDetailSummary *schema.ManArticleDetailSummary
		wantErr                      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotContent, gotContentMD5, gotArticleDetailSummary, err := TryGetManArticleContentInfo(tt.args.userID, tt.args.bboardID, tt.args.articleID, tt.args.c, tt.args.isSystem, tt.args.isContent)
			if (err != nil) != tt.wantErr {
				t.Errorf("TryGetManArticleContentInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotContent, tt.expectedContent) {
				t.Errorf("TryGetManArticleContentInfo() gotContent = %v, want %v", gotContent, tt.expectedContent)
			}
			if gotContentMD5 != tt.expectedContentMD5 {
				t.Errorf("TryGetManArticleContentInfo() gotContentMD5 = %v, want %v", gotContentMD5, tt.expectedContentMD5)
			}
			if !reflect.DeepEqual(gotArticleDetailSummary, tt.expectedArticleDetailSummary) {
				t.Errorf("TryGetManArticleContentInfo() gotArticleDetailSummary = %v, want %v", gotArticleDetailSummary, tt.expectedArticleDetailSummary)
			}
		})
	}
}
