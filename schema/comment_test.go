package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func TestUpdateComments(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer Comment_c.Drop()

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

	_ = UpdateComments(testComments0, types.NanoTS(1234567890000000000))

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
