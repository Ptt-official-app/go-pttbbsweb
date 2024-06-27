package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"go.mongodb.org/mongo-driver/bson"
)

func TestUpdateContentBlocks(t *testing.T) {
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

	type args struct {
		contentBlocks []*ContentBlock
		updateNanoTS  types.NanoTS
	}
	tests := []struct {
		name     string
		args     args
		expected []*ContentBlock
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{contentBlocks: contentBlocks, updateNanoTS: 1234567890000000000},
			expected: contentBlocks,
		},
		{
			args:     args{contentBlocks: contentBlocks, updateNanoTS: 1234567890000000000},
			expected: contentBlocks,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if err := UpdateContentBlocks(tt.args.contentBlocks, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("UpdateContentBlocks() error = %v, wantErr %v", err, tt.wantErr)
			}

			contentBlock0 := tt.args.contentBlocks[0]
			var got []*ContentBlock
			query := bson.M{
				CONTENT_BLOCK_BBOARD_ID_b:  contentBlock0.BBoardID,
				CONTENT_BLOCK_ARTICLE_ID_b: contentBlock0.ArticleID,
				CONTENT_BLOCK_CONTENT_ID_b: contentBlock0.ContentID,
			}
			sort := bson.D{
				{Key: CONTENT_BLOCK_IDX_b, Value: 1},
			}
			err := ContentBlock_c.Find(query, 0, &got, nil, sort)
			if err != nil {
				t.Errorf("UpdateContentBlocks: unable to find: e: %v", err)
				return
			}

			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
		wg.Wait()
	}
}
