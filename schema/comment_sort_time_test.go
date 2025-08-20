package schema

import (
	"reflect"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func TestGetCommentSortTimeMapByCommentIDs(t *testing.T) {
	setupTest()
	defer teardownTest()

	_ = UpdateComments(testComments0, types.NanoTS(1334567890000000000))

	commentIDs0 := []types.CommentID{
		"EYFhmOhREkA:Es26f7U0EXdr7Gp4a9N8pQ",
		"EYFhmOhgVIA:gmrKWXE7BjV-1U89GcPqHg",
	}
	expected0 := map[types.CommentID]types.NanoTS{
		"EYFhmOhREkA:Es26f7U0EXdr7Gp4a9N8pQ": 1261396680001000000,
		"EYFhmOhgVIA:gmrKWXE7BjV-1U89GcPqHg": 1261396680002000000,
	}

	type args struct {
		boardID    bbs.BBoardID
		articleID  bbs.ArticleID
		commentIDs []types.CommentID
	}
	tests := []struct {
		name                string
		args                args
		expectedSortTimeMap map[types.CommentID]types.NanoTS
		wantErr             bool
	}{
		// TODO: Add test cases.
		{
			args:                args{boardID: "test", articleID: "test", commentIDs: commentIDs0},
			expectedSortTimeMap: expected0,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotSortTimeMap, err := GetCommentSortTimeMapByCommentIDs(tt.args.boardID, tt.args.articleID, tt.args.commentIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCommentSortTimeMapByCommentIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSortTimeMap, tt.expectedSortTimeMap) {
				t.Errorf("GetCommentSortTimeMapByCommentIDs() = %v, want %v", gotSortTimeMap, tt.expectedSortTimeMap)
			}
		})
		wg.Wait()
	}
}
