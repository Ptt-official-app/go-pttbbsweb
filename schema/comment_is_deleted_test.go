package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"go.mongodb.org/mongo-driver/bson"
)

func TestRemoveOldFirstComments(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer Comment_c.Drop()

	_ = UpdateComments(testComments0, types.NanoTS(1334567890000000000))
	_ = UpdateComments(testComments1, types.NanoTS(1334567890000000000))

	type args struct {
		bboardID     bbs.BBoardID
		articleID    bbs.ArticleID
		updateNanoTS types.NanoTS
	}
	tests := []struct {
		name     string
		args     args
		expected int
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				bboardID:     bbs.BBoardID("test"),
				articleID:    bbs.ArticleID("test"),
				updateNanoTS: types.NanoTS(1334567890000000000),
			},
			expected: 7,
		},
		{
			args: args{
				bboardID:     bbs.BBoardID("test"),
				articleID:    bbs.ArticleID("test"),
				updateNanoTS: types.NanoTS(1334567890000000001),
			},
			expected: 5,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if err := RemoveOldFirstComments(tt.args.bboardID, tt.args.articleID, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("RemoveOldComments() error = %v, wantErr %v", err, tt.wantErr)
			}

			query := bson.M{
				COMMENT_BBOARD_ID_b:  tt.args.bboardID,
				COMMENT_ARTICLE_ID_b: tt.args.articleID,
				COMMENT_IS_DELETED_b: bson.M{"$exists": false},
			}
			count, _ := Comment_c.Count(query, 0)
			if int(count) != tt.expected {
				t.Errorf("RemoveOldComments: count: %v expected: %v", count, tt.expected)
			}
		})
	}
	wg.Wait()
}

func TestRemoveCommentIDs(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer Comment_c.Drop()

	_ = UpdateComments(testComments0, types.NanoTS(1334567890000000000))

	type args struct {
		bboardID           bbs.BBoardID
		articleID          bbs.ArticleID
		toRemoveCommentIDs []types.CommentID
		updateNanoTS       types.NanoTS
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		expected int
	}{
		// TODO: Add test cases.
		{
			args: args{
				bboardID:  bbs.BBoardID("test"),
				articleID: bbs.ArticleID("test"),
				toRemoveCommentIDs: []types.CommentID{
					"EYFhmOhREkA:Es26f7U0EXdr7Gp4a9N8pQ",
					"EYFhmOhgVIA:gmrKWXE7BjV-1U89GcPqHg",
				},
				updateNanoTS: types.NanoTS(1334567890000000000),
			},
			expected: 5,
		},
		{
			args: args{
				bboardID:  bbs.BBoardID("test"),
				articleID: bbs.ArticleID("test"),
				toRemoveCommentIDs: []types.CommentID{
					"EYFhmOhREkA:Es26f7U0EXdr7Gp4a9N8pQ",
					"EYFhmOhgVIA:gmrKWXE7BjV-1U89GcPqHg",
				},
				updateNanoTS: types.NanoTS(1334567890000000001),
			},
			expected: 3,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if err := RemoveCommentIDs(tt.args.bboardID, tt.args.articleID, tt.args.toRemoveCommentIDs, tt.args.updateNanoTS, "some-reason"); (err != nil) != tt.wantErr {
				t.Errorf("RemoveCommentIDs() error = %v, wantErr %v", err, tt.wantErr)
			}

			query := bson.M{
				COMMENT_BBOARD_ID_b:  tt.args.bboardID,
				COMMENT_ARTICLE_ID_b: tt.args.articleID,
				COMMENT_IS_DELETED_b: bson.M{"$exists": false},
			}

			count, _ := Comment_c.Count(query, 0)
			if int(count) != tt.expected {
				t.Errorf("RemoveCommentIDs: count: %v expected: %v", count, tt.expected)
			}
		})
	}
	wg.Wait()
}
