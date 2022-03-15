package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func TestUpdateManArticleContentInfo(t *testing.T) {
	setupTest()
	defer teardownTest()

	contentInfo0 := &ManArticleContentInfo{
		ContentMD5:          "contentMD5",
		ContentID:           "contentID",
		ContentUpdateNanoTS: 1234567890000000000,
	}

	type args struct {
		bboardID    bbs.BBoardID
		articleID   types.ManArticleID
		contentInfo *ManArticleContentInfo
	}
	tests := []struct {
		name     string
		args     args
		expected *ManArticleContentInfo
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{bboardID: "10_WhoAmI", articleID: "D1D2/M.1234567890.A.ABC", contentInfo: contentInfo0},
			expected: contentInfo0,
		},
		{
			args:     args{bboardID: "10_WhoAmI", articleID: "D1D2/M.1234567890.A.ABC", contentInfo: contentInfo0},
			expected: contentInfo0,
			wantErr:  true,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if err := UpdateManArticleContentInfo(tt.args.bboardID, tt.args.articleID, tt.args.contentInfo); (err != nil) != tt.wantErr {
				t.Errorf("UpdateManArticleContentInfo() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := GetManArticleContentInfo(tt.args.bboardID, tt.args.articleID, true)
			if err != nil {
				t.Errorf("UpdateManArticleContentInfo: unable to get: e: %v", err)

				testutil.TDeepEqual(t, "got", got, tt.expected)
			}
		})
		wg.Wait()
	}
}
