package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

func TestGetAllCommentMD5s(t *testing.T) {
	setupTest()
	defer teardownTest()

	_ = UpdateComments(testComments0, types.NanoTS(1334567890000000000))

	expected0 := []*CommentMD5{
		{
			BBoardID:   "test",
			ArticleID:  "test",
			CommentID:  "EYFhmOhREkA:Es26f7U0EXdr7Gp4a9N8pQ",
			TheType:    ptttype.COMMENT_TYPE_RECOMMEND,
			CreateTime: 1261396680001000000,
			MD5:        "Es26f7U0EXdr7Gp4a9N8pQ",

			SortTime: 1261396680001000000,
			TheDate:  "12/21 19:58",
		},
		{
			BBoardID:   "test",
			ArticleID:  "test",
			CommentID:  "EYFhmOhgVIA:gmrKWXE7BjV-1U89GcPqHg",
			TheType:    ptttype.COMMENT_TYPE_RECOMMEND,
			CreateTime: 1261396680002000000,
			MD5:        "gmrKWXE7BjV-1U89GcPqHg",

			SortTime: 1261396680002000000,
			TheDate:  "12/21 19:58",
		},
		{
			BBoardID:   "test",
			ArticleID:  "test",
			CommentID:  "EYFhmOhvlsA:cpqbGyLoF_jIyITF4bv-rQ",
			TheType:    ptttype.COMMENT_TYPE_RECOMMEND,
			CreateTime: 1261396680003000000,
			MD5:        "cpqbGyLoF_jIyITF4bv-rQ",

			SortTime: 1261396680003000000,
			TheDate:  "12/21 19:58",
		},
		{
			BBoardID:   "test",
			ArticleID:  "test",
			CommentID:  "EYFhmOhvlsA:cpqbGyLoF_jIyITF4bv-rQ:R",
			TheType:    ptttype.COMMENT_TYPE_REPLY,
			CreateTime: 1261396680003100000,
			MD5:        "VMu8YlVFJ4k06pYnUILy4w",

			EditNanoTS: 1608551574000000000,
			SortTime:   1261396680003100000,
			TheDate:    "",
		},
		{
			BBoardID:   "test",
			ArticleID:  "test",
			CommentID:  "EYFuT-Ew6AA:ALE6XIa5ARhXunryJTB3xg",
			TheType:    ptttype.COMMENT_TYPE_RECOMMEND,
			CreateTime: 1261410660000000000,
			MD5:        "ALE6XIa5ARhXunryJTB3xg",

			SortTime: 1261410660000000000,
			TheDate:  "12/21 23:51",
		},
	}

	type args struct {
		boardID   bbs.BBoardID
		articleID bbs.ArticleID
	}
	tests := []struct {
		name                string
		args                args
		expectedCommentMD5s []*CommentMD5
		wantErr             bool
	}{
		// TODO: Add test cases.
		{
			args:                args{boardID: "test", articleID: "test"},
			expectedCommentMD5s: expected0,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotCommentMD5s, err := GetAllCommentMD5s(tt.args.boardID, tt.args.articleID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllCommentMD5s() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotCommentMD5s, tt.expectedCommentMD5s)
		})
		wg.Wait()
	}
}
