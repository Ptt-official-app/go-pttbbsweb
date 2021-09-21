package schema

import (
	"reflect"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func TestGetAllCommentDBCSs(t *testing.T) {
	setupTest()
	defer teardownTest()

	_ = UpdateComments(testComments0, types.NanoTS(1334567890000000000))

	expected0 := []*CommentDBCS{
		{
			BBoardID:  "test",
			ArticleID: "test",
			CommentID: "EYFhmOhREkA:Es26f7U0EXdr7Gp4a9N8pQ",
			SortTime:  1261396680001000000,
			DBCS:      []byte("\x1b[1;37m\xb1\xc0 \x1b[33mfoolfighter\x1b[m\x1b[33m: \xa4g\xac_\xc1T\xb4X\xa6\xa8\xa6\xb3\xa5x\xa5_\xa5\xab\xa4\xe1\xc4y\xb0\xda\xa1H                        \x1b[m 12/21 19:58\r"),
		},
		{
			BBoardID:  "test",
			ArticleID: "test",
			CommentID: "EYFhmOhgVIA:gmrKWXE7BjV-1U89GcPqHg",
			SortTime:  1261396680002000000,
			DBCS:      []byte("\x1b[1;37m\xb1\xc0 \x1b[33myehpi\x1b[m\x1b[33m: \xb6\xc0\xa6\xb3\xb6W\xa4j\xa4\xe4\xb2\xbc\xa1]\xaa\xab\xb2z\xa1^                                    \x1b[m 12/21 19:58\r"),
		},
		{
			BBoardID:  "test",
			ArticleID: "test",
			CommentID: "EYFhmOhvlsA:cpqbGyLoF_jIyITF4bv-rQ",
			SortTime:  1261396680003000000,
			DBCS:      []byte("\x1b[1;37m\xb1\xc0 \x1b[33mlockeyman\x1b[m\x1b[33m: \xbcP                                                  \x1b[m 12/21 19:58\r"),
		},
		{
			BBoardID:  "test",
			ArticleID: "test",
			CommentID: "EYFhmOhvlsA:cpqbGyLoF_jIyITF4bv-rQ:R",
			SortTime:  1261396680003100000,
			DBCS:      []byte("\r\ntest123123\r\n\r\ntest124124\r\n\r\ntest125125\r\n\r"),
		},
		{
			BBoardID:  "test",
			ArticleID: "test",
			CommentID: "EYFuT-Ew6AA:ALE6XIa5ARhXunryJTB3xg",
			SortTime:  1261410660000000000,
			DBCS:      []byte("\x1b[1;37m\xb1\xc0 \x1b[33mdeathdancer\x1b[m\x1b[33m: \xaa\xfc\xa5_\xa4~\xa8S\xa6\xb3\xbf\xe9                                      \x1b[m 12/21 23:51\r"),
		},
	}

	type args struct {
		boardID   bbs.BBoardID
		articleID bbs.ArticleID
	}
	tests := []struct {
		name                string
		args                args
		expectedCommentDBCS []*CommentDBCS
		wantErr             bool
	}{
		// TODO: Add test cases.
		{
			args:                args{boardID: "test", articleID: "test"},
			expectedCommentDBCS: expected0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCommentDBCS, err := GetAllCommentDBCSs(tt.args.boardID, tt.args.articleID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllCommentDBCSs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCommentDBCS, tt.expectedCommentDBCS) {
				t.Errorf("GetAllCommentDBCSs() = %v, want %v", gotCommentDBCS, tt.expectedCommentDBCS)
			}
		})
	}
}
