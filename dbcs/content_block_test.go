package dbcs

import (
	"reflect"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func TestParseContentBlocks(t *testing.T) {
	setupTest()
	defer teardownTest()

	contentBlocks := []*schema.ContentBlock{
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

	type args struct {
		boardID      bbs.BBoardID
		articleID    bbs.ArticleID
		content      [][]*types.Rune
		contentMD5   string
		updateNanoTS types.NanoTS
	}
	tests := []struct {
		name                  string
		args                  args
		expectedContentID     types.ContentID
		expectedContentBlocks []*schema.ContentBlock
	}{
		// TODO: Add test cases.
		{
			args: args{
				boardID:      "boardID0",
				articleID:    "articleID0",
				content:      testContent11Utf8,
				contentMD5:   "IKCj3KzpwP5pcJxOAPNDNQ",
				updateNanoTS: 1234567890000000000,
			},
			expectedContentID:     "ESIQ9HaNtAA:IKCj3KzpwP5pcJxOAPNDNQ",
			expectedContentBlocks: contentBlocks,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotContentID, gotContentBlocks := ParseContentBlocks(tt.args.boardID, tt.args.articleID, tt.args.content, tt.args.contentMD5, tt.args.updateNanoTS)
			if !reflect.DeepEqual(gotContentID, tt.expectedContentID) {
				t.Errorf("ParseContentBlocks() gotContentID = %v, want %v", gotContentID, tt.expectedContentID)
			}
			testutil.TDeepEqual(t, "got", gotContentBlocks, tt.expectedContentBlocks)
		})
		wg.Wait()
	}
}
