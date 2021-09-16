package schema

import (
	"fmt"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func TestUpdateComments(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		comments     []*Comment
		updateNanoTS types.NanoTS
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{comments: testComments0, updateNanoTS: types.NanoTS(1234567890000000000)},
		},
		{
			args: args{comments: testComments0, updateNanoTS: types.NanoTS(1234567890000000000)},
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if err := UpdateComments(tt.args.comments, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("UpdateComments() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	wg.Wait()
}

func TestComment_CleanReply(t *testing.T) {
	setupTest()
	defer teardownTest()

	tests := []struct {
		name     string
		c        *Comment
		expected *Comment
	}{
		// TODO: Add test cases.
		{
			c:        testReply0,
			expected: testExpectedReply0,
		},
		{
			c:        testReply1,
			expected: testExpectedReply1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.c
			c.CleanReply()

			testutil.TDeepEqual(t, "c", c, tt.expected)
		})
	}
}

func TestCountComments(t *testing.T) {
	setupTest()
	defer teardownTest()

	_ = UpdateComments(testComments0, types.NanoTS(1334567890000000000))

	type args struct {
		boardID   bbs.BBoardID
		articleID bbs.ArticleID
	}
	tests := []struct {
		name              string
		args              args
		expectedNComments int
		wantErr           bool
	}{
		// TODO: Add test cases.
		{
			args:              args{boardID: bbs.BBoardID("test"), articleID: bbs.ArticleID("test")},
			expectedNComments: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNComments, err := CountComments(tt.args.boardID, tt.args.articleID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotNComments != tt.expectedNComments {
				t.Errorf("CountComments() = %v, want %v", gotNComments, tt.expectedNComments)
			}
		})
	}
}

func TestGetComments(t *testing.T) {
	setupTest()
	defer teardownTest()

	_ = UpdateComments(testComments0, types.NanoTS(1334567890000000000))

	expectedComments0 := []*Comment{testComments0[4], testComments0[3]}
	expectedComments1 := []*Comment{testComments0[3], testComments0[2]}
	expectedComments2 := []*Comment{testComments0[0], testComments0[1]}
	expectedComments3 := []*Comment{testComments0[1], testComments0[2]}

	type args struct {
		boardID      bbs.BBoardID
		articleID    bbs.ArticleID
		createNanoTS types.NanoTS
		commentID    types.CommentID
		descending   bool
		limit        int
	}
	tests := []struct {
		name             string
		args             args
		expectedComments []*Comment
		wantErr          bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				boardID:    "test",
				articleID:  "test",
				limit:      2,
				descending: true,
			},
			expectedComments: expectedComments0,
		},
		{
			args: args{
				boardID:      "test",
				articleID:    "test",
				limit:        2,
				createNanoTS: types.NanoTS(1261396680003100000),
				commentID:    types.CommentID("EYFhmOhvlsA:cpqbGyLoF_jIyITF4bv-rQ:R"),
				descending:   true,
			},
			expectedComments: expectedComments1,
		},
		{
			args: args{
				boardID:   "test",
				articleID: "test",
				limit:     2,
			},
			expectedComments: expectedComments2,
		},
		{
			args: args{
				boardID:      "test",
				articleID:    "test",
				limit:        2,
				createNanoTS: types.NanoTS(1261396680002000000),
				commentID:    "EYFhmOhgVIA:gmrKWXE7BjV-1U89GcPqHg",
			},
			expectedComments: expectedComments3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotComments, err := GetComments(tt.args.boardID, tt.args.articleID, tt.args.createNanoTS, tt.args.commentID, tt.args.descending, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			testutil.TDeepEqual(t, "got", gotComments, tt.expectedComments)
		})
	}
}

func TestComment_SetSortTime(t *testing.T) {
	setupTest()
	defer teardownTest()
	c0 := &Comment{
		CommentID: "test0",
		TheType:   ptttype.COMMENT_TYPE_BOO,
		MD5:       "md50",
	}

	expected0 := &Comment{
		CommentID: "ESIQ9HaNtAA:md50",
		TheType:   ptttype.COMMENT_TYPE_BOO,
		MD5:       "md50",
		SortTime:  1234567890000000000,
	}

	c1 := &Comment{
		CommentID: "test1:R",
		TheType:   ptttype.COMMENT_TYPE_REPLY,
		MD5:       "md50",
	}

	expected1 := &Comment{
		CommentID: "test1:R",
		TheType:   ptttype.COMMENT_TYPE_REPLY,
		MD5:       "md50",
		SortTime:  1234567890000000000,
	}

	type args struct {
		nanoTS types.NanoTS
	}
	tests := []struct {
		name     string
		c        *Comment
		args     args
		expected *Comment
	}{
		// TODO: Add test cases.
		{
			c:        c0,
			args:     args{nanoTS: 1234567890000000000},
			expected: expected0,
		},
		{
			c:        c1,
			args:     args{nanoTS: 1234567890000000000},
			expected: expected1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.c
			c.SetSortTime(tt.args.nanoTS)
			testutil.TDeepEqual(t, "c", c, tt.expected)
		})
	}
}

func TestGetCommentMapByCommentIDs(t *testing.T) {
	setupTest()
	defer teardownTest()

	_ = UpdateComments(testComments0, types.NanoTS(1334567890000000000))

	commentIDs0 := []types.CommentID{
		"EYFhmOhREkA:Es26f7U0EXdr7Gp4a9N8pQ",
		"EYFhmOhgVIA:gmrKWXE7BjV-1U89GcPqHg",
	}
	expected0 := map[types.CommentID]*Comment{
		"EYFhmOhREkA:Es26f7U0EXdr7Gp4a9N8pQ": testComments0[0],
		"EYFhmOhgVIA:gmrKWXE7BjV-1U89GcPqHg": testComments0[1],
	}

	type args struct {
		boardID    bbs.BBoardID
		articleID  bbs.ArticleID
		commentIDs []types.CommentID
	}
	tests := []struct {
		name               string
		args               args
		expectedCommentMap map[types.CommentID]*Comment
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			args:               args{boardID: "test", articleID: "test", commentIDs: commentIDs0},
			expectedCommentMap: expected0,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotCommentMap, err := GetCommentMapByCommentIDs(tt.args.boardID, tt.args.articleID, tt.args.commentIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCommentMapByCommentIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for idx, each := range tt.expectedCommentMap {
				got := gotCommentMap[idx]
				testutil.TDeepEqual(t, fmt.Sprintf("got-%v", (idx)), got, each)
			}
		})
		wg.Wait()
	}
}
