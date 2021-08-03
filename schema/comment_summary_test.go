package schema

import (
	"reflect"
	"sort"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func TestGetCommentSummaries(t *testing.T) {
	setupTest()
	defer teardownTest()

	_ = UpdateComments(testComments0, types.NanoTS(1234567890000000000))
	_ = UpdateComments(testComments1, types.NanoTS(1234567890000000000))

	type args struct {
		bboardID    bbs.BBoardID
		articleID   bbs.ArticleID
		startNanoTS types.NanoTS
		endNanoTS   types.NanoTS
	}
	tests := []struct {
		name                     string
		args                     args
		expectedCommentSummaries []*CommentSummary
		wantErr                  bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				bboardID:    bbs.BBoardID("test"),
				articleID:   bbs.ArticleID("test"),
				startNanoTS: types.NanoTS(1261396500000000000),
				endNanoTS:   types.NanoTS(1261396680003100000) + types.NanoTS(1), // excluding endNanoTS
			},
			expectedCommentSummaries: testCommentSummaries0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCommentSummaries, err := GetCommentSummaries(tt.args.bboardID, tt.args.articleID, tt.args.startNanoTS, tt.args.endNanoTS)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCommentSummaries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			sort.SliceStable(gotCommentSummaries, func(i int, j int) bool {
				return gotCommentSummaries[i].CreateTime <= gotCommentSummaries[j].CreateTime
			})

			testutil.TDeepEqual(t, "got", gotCommentSummaries, tt.expectedCommentSummaries)
		})
	}
}

func TestUpdateCommentSummaries(t *testing.T) {
	setupTest()
	defer teardownTest()

	_ = UpdateComments(testComments0, types.NanoTS(1234567890000000000))
	_ = UpdateComments(testComments1, types.NanoTS(1234567890000000000))

	type args struct {
		bboardID         bbs.BBoardID
		articleID        bbs.ArticleID
		commentSummaries []*CommentSummary
		updateNanoTS     types.NanoTS
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		expected []*CommentSummary
	}{
		// TODO: Add test cases.
		{
			args: args{
				bboardID:         "test",
				articleID:        "test",
				commentSummaries: testCommentSummaries0,
				updateNanoTS:     types.NanoTS(1334567890000000000),
			},
			expected: testCommentSummaries0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateCommentSummaries(tt.args.bboardID, tt.args.articleID, tt.args.commentSummaries, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("UpdateCommentSummaries() error = %v, wantErr %v", err, tt.wantErr)
			}

			gotCommentSummaries, _ := GetCommentSummaries(tt.args.bboardID, tt.args.articleID, types.NanoTS(1261396500000000000), types.NanoTS(1261396680003100001))

			sort.SliceStable(gotCommentSummaries, func(i int, j int) bool {
				return gotCommentSummaries[i].CreateTime <= gotCommentSummaries[j].CreateTime
			})

			testutil.TDeepEqual(t, "got", gotCommentSummaries, tt.expected)
		})
	}
}

func TestGetCommentSummariesByOwnerID(t *testing.T) {
	setupTest()
	defer teardownTest()

	_ = UpdateComments(testComments0, types.NanoTS(1234567890000000000))
	_ = UpdateComments(testComments1, types.NanoTS(1234567890000000000))

	testCommentSummaries0 := []*CommentSummary{testCommentSummaries0[3]}

	type args struct {
		ownerID       bbs.UUserID
		startSortTime types.NanoTS
		descending    bool
		limit         int
	}
	tests := []struct {
		name           string
		args           args
		expectedResult []*CommentSummary
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				ownerID: "lockeyman",
				limit:   200,
			},
			expectedResult: testCommentSummaries0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := GetBasicCommentSummariesByOwnerID(tt.args.ownerID, tt.args.startSortTime, tt.args.descending, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCommentSummariesByOwnerID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.expectedResult) {
				t.Errorf("GetCommentSummariesByOwnerID() = %v, want %v", gotResult, tt.expectedResult)
			}
		})
	}
}
