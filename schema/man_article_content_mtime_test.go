package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

func TestUpdateManArticleContentMTime(t *testing.T) {
	setupTest()
	defer teardownTest()

	a0 := &ManArticleSummary{
		BBoardID:   "10_WhoAmI",
		LevelIdx:   "D1D2",
		ArticleID:  "D1D2/M.1234567890.A.ABC",
		CreateTime: 1234567890000000000,
		MTime:      1234567890000000000,

		Title:        "這是 title",
		UpdateNanoTS: 1234567890000000000,
		Idx:          0,
	}

	expected0 := &ManArticleContentMTime{
		ContentMTime: 1234567890000000000,
	}

	type args struct {
		bboardID     bbs.BBoardID
		articleID    types.ManArticleID
		contentMTime types.NanoTS
	}
	tests := []struct {
		name           string
		args           args
		articleSumamry *ManArticleSummary
		expected       *ManArticleContentMTime
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			args:    args{bboardID: "10_WhoAmI", articleID: "D1D2/M.1234567890.A.ABC", contentMTime: 1234567890000000000},
			wantErr: true,
		},
		{
			args:           args{bboardID: "10_WhoAmI", articleID: "D1D2/M.1234567890.A.ABC", contentMTime: 1234567890000000000},
			articleSumamry: a0,
			expected:       expected0,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if tt.articleSumamry != nil {
				_ = UpdateManArticleSummaries([]*ManArticleSummary{tt.articleSumamry}, 1234567890000000000)
			}

			if err := UpdateManArticleContentMTime(tt.args.bboardID, tt.args.articleID, tt.args.contentMTime); (err != nil) != tt.wantErr {
				t.Errorf("UpdateManArticleContentMTime() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := GetManArticleContentMTime(tt.args.bboardID, tt.args.articleID)
			if err != nil {
				t.Errorf("UpdateManArticleContentMTime: unable to get: e: %v", err)
			}
			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
		wg.Wait()
	}
}
