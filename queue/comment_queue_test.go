// +build !noqueue

package queue

import (
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"go.mongodb.org/mongo-driver/bson"
)

func TestQueueCommentDBCS(t *testing.T) {
	setupTest()
	defer teardownTest()

	schema.Comment_c.Drop()
	defer schema.Comment_c.Drop()

	//move setupTest inside
	type args struct {
		bboardID              bbs.BBoardID
		articleID             bbs.ArticleID
		ownerID               bbs.UUserID
		commentDBCS           []byte
		firstCommentsLastTime types.NanoTS
		updateNanoTS          types.NanoTS
	}
	tests := []struct {
		name     string
		args     args
		expected []*schema.Comment
		wantErr  bool
		sleepTS  int
	}{
		// TODO: Add test cases.
		{
			args: args{
				bboardID:              "test",
				articleID:             "test",
				ownerID:               "cheinshin",
				commentDBCS:           testTheRestCommentsDBCS11,
				firstCommentsLastTime: types.NanoTS(1261396020004000000),
				updateNanoTS:          types.NanoTS(1334567890000000000),
			},
			expected: testTheRestComments11,
		},
		{
			args: args{
				bboardID:              "test",
				articleID:             "test",
				ownerID:               "cheinshin",
				commentDBCS:           testTheRestCommentsDBCS11,
				firstCommentsLastTime: types.NanoTS(1261396020004000000),
				updateNanoTS:          types.NanoTS(1434567890000000000),
			},
			expected: testTheRestComments11,
			sleepTS:  1,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			if tt.sleepTS > 0 {
				time.Sleep(time.Duration(tt.sleepTS) * time.Second)
			}
			if err := QueueCommentDBCS(tt.args.bboardID, tt.args.articleID, tt.args.ownerID, tt.args.commentDBCS, tt.args.firstCommentsLastTime, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("QueueCommentDBCS() error = %v, wantErr %v", err, tt.wantErr)

			}

			time.Sleep(time.Duration(5) * time.Second)

			query := bson.M{
				schema.COMMENT_BBOARD_ID_b:  "test",
				schema.COMMENT_ARTICLE_ID_b: "test",
			}
			var got []*schema.Comment
			_ = schema.Comment_c.Find(query, 0, &got, nil, nil)

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].CreateTime <= got[j].CreateTime
			})
			for _, each := range got {
				each.UpdateNanoTS = types.NanoTS(0)
			}
			testutil.TDeepEqual(t, "got", got, tt.expected)

		})
	}
	wg.Wait()
}
