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
	type fields struct {
		BBoardID           bbs.BBoardID
		ArticleID          bbs.ArticleID
		CommentID          types.CommentID
		TheType            types.CommentType
		RefIDs             []types.CommentID
		IsDeleted          bool
		DeleteReason       string
		CreateTime         types.NanoTS
		Owner              bbs.UUserID
		Content            [][]*types.Rune
		IP                 string
		Host               string
		MD5                string
		FirstCreateTime    types.NanoTS
		InferredCreateTime types.NanoTS
		NewCreateTime      types.NanoTS
		SortTime           types.NanoTS
		TheDate            string
		DBCS               []byte
		EditNanoTS         types.NanoTS
		UpdateNanoTS       types.NanoTS
	}
	type args struct {
		nanoTS types.NanoTS
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Comment{
				BBoardID:           tt.fields.BBoardID,
				ArticleID:          tt.fields.ArticleID,
				CommentID:          tt.fields.CommentID,
				TheType:            tt.fields.TheType,
				RefIDs:             tt.fields.RefIDs,
				IsDeleted:          tt.fields.IsDeleted,
				DeleteReason:       tt.fields.DeleteReason,
				CreateTime:         tt.fields.CreateTime,
				Owner:              tt.fields.Owner,
				Content:            tt.fields.Content,
				IP:                 tt.fields.IP,
				Host:               tt.fields.Host,
				MD5:                tt.fields.MD5,
				FirstCreateTime:    tt.fields.FirstCreateTime,
				InferredCreateTime: tt.fields.InferredCreateTime,
				NewCreateTime:      tt.fields.NewCreateTime,
				SortTime:           tt.fields.SortTime,
				TheDate:            tt.fields.TheDate,
				DBCS:               tt.fields.DBCS,
				EditNanoTS:         tt.fields.EditNanoTS,
				UpdateNanoTS:       tt.fields.UpdateNanoTS,
			}
			c.SetSortTime(tt.args.nanoTS)
		})
	}
}
