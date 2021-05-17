package dbcs

import (
	"reflect"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func Test_calcEditDistance(t *testing.T) {
	setupTest()
	defer teardownTest()

	comments0 := []*schema.Comment{
		{MD5: "h"},
		{MD5: "o"},
		{MD5: "r"},
		{MD5: "s"},
		{MD5: "e"},
	}

	commentMD5s0 := []*schema.CommentMD5{
		{MD5: "r"},
		{MD5: "o"},
		{MD5: "s"},
	}

	scoreMat0 := [][]int{
		//  r  o  s
		{0, 1, 2, 3}, //
		{1, 2, 3, 4}, //h
		{2, 3, 2, 3}, //o
		{3, 2, 3, 4}, //r
		{4, 3, 4, 3}, //s
		{5, 4, 5, 4}, //e
	}

	comments1 := []*schema.Comment{
		{MD5: "a"},
		{MD5: "b"},
		{MD5: "c"},
		{MD5: "d"},
		{MD5: "a"},
		{MD5: "b"},
		{MD5: "c"},
	}

	commentMD5s1 := []*schema.CommentMD5{
		{MD5: "b"},
		{MD5: "c"},
	}

	scoreMat1 := [][]int{
		//  b  c
		{0, 1, 2}, //
		{1, 2, 3}, //a
		{2, 1, 2}, //b
		{3, 2, 1}, //c
		{4, 3, 2}, //d
		{5, 4, 3}, //a
		{6, 5, 4}, //b
		{7, 6, 5}, //c
	}

	comments2 := []*schema.Comment{
		{MD5: "a"},
		{MD5: "b"},
		{MD5: "c"},
		{MD5: "d"},
		{MD5: "a"},
		{MD5: "b"},
		{MD5: "a"},
		{MD5: "c"},
	}

	commentMD5s2 := []*schema.CommentMD5{
		{MD5: "b"},
		{MD5: "e"},
		{MD5: "a"},
		{MD5: "c"},
	}

	scoreMat2 := [][]int{
		//  b  e  a, c
		{0, 1, 2, 3, 4}, //
		{1, 2, 3, 2, 3}, //a
		{2, 1, 2, 3, 4}, //b
		{3, 2, 3, 4, 3}, //c
		{4, 3, 4, 5, 4}, //d
		{5, 4, 5, 4, 5}, //a
		{6, 5, 6, 5, 6}, //b
		{7, 6, 7, 6, 7}, //a
		{8, 7, 8, 7, 6}, //c
	}
	type args struct {
		comments       []*schema.Comment
		allCommentMD5s []*schema.CommentMD5
	}
	tests := []struct {
		name             string
		args             args
		expectedScoreMat [][]int
	}{
		// TODO: Add test cases.
		{
			name:             "horse vs. ros",
			args:             args{comments: comments0, allCommentMD5s: commentMD5s0},
			expectedScoreMat: scoreMat0,
		},
		{
			name:             "abcdabc vs. bc",
			args:             args{comments: comments1, allCommentMD5s: commentMD5s1},
			expectedScoreMat: scoreMat1,
		},
		{
			name:             "abcdabac vs. beac",
			args:             args{comments: comments2, allCommentMD5s: commentMD5s2},
			expectedScoreMat: scoreMat2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotScoreMat := calcEditDistance(tt.args.comments, tt.args.allCommentMD5s); !reflect.DeepEqual(gotScoreMat, tt.expectedScoreMat) {
				t.Errorf("calcEditDistance() = %v, want %v", gotScoreMat, tt.expectedScoreMat)
			}
		})
	}
}

func Test_calcEditDistanceBacktracking(t *testing.T) {
	setupTest()
	defer teardownTest()

	scoreMat0 := [][]int{
		//  r  o  s
		{0, 1, 2, 3}, //
		{1, 2, 3, 4}, //h
		{2, 3, 2, 3}, //o
		{3, 2, 3, 4}, //r
		{4, 3, 4, 3}, //s
		{5, 4, 5, 4}, //e
	}

	comments0 := []*schema.Comment{
		{MD5: "h"},
		{MD5: "o"},
		{MD5: "r"},
		{MD5: "s"},
		{MD5: "e", SortTime: 1234567894000000000},
	}

	commentMD5s0 := []*schema.CommentMD5{
		{MD5: "r", SortTime: 1234567891000000000},
		{MD5: "o", SortTime: 1234567892000000000},
		{MD5: "s", SortTime: 1234567893000000000},
	}

	expected0 := []*EDInfo{
		{
			Op:          ED_OP_DELETE,
			OrigComment: &schema.CommentMD5{MD5: "r", SortTime: 1234567891000000000},
			SortTime:    1234567891000000000,
		},
		{
			Op:         ED_OP_ADD,
			NewComment: &schema.Comment{MD5: "h"},
		},
		{
			Op:          ED_OP_SAME,
			NewComment:  &schema.Comment{MD5: "o"},
			OrigComment: &schema.CommentMD5{MD5: "o", SortTime: 1234567892000000000},
			SortTime:    1234567892000000000,
		},
		{
			Op:         ED_OP_ADD,
			NewComment: &schema.Comment{MD5: "r"},
		},
		{
			Op:          ED_OP_SAME,
			NewComment:  &schema.Comment{MD5: "s"},
			OrigComment: &schema.CommentMD5{MD5: "s", SortTime: 1234567893000000000},
			SortTime:    1234567893000000000,
		},
		{
			Op:         ED_OP_ADD,
			NewComment: &schema.Comment{MD5: "e", SortTime: 1234567894000000000},
			SortTime:   1234567894000000000,
		},
	}

	comments1 := []*schema.Comment{
		{MD5: "a"},
		{MD5: "b"},
		{MD5: "c"},
		{MD5: "d"},
		{MD5: "a"},
		{MD5: "b"},
		{MD5: "c"},
	}

	commentMD5s1 := []*schema.CommentMD5{
		{MD5: "b"},
		{MD5: "c"},
	}

	scoreMat1 := [][]int{
		//  b  c
		{0, 1, 2}, //
		{1, 2, 3}, //a
		{2, 1, 2}, //b
		{3, 2, 1}, //c
		{4, 3, 2}, //d
		{5, 4, 3}, //a
		{6, 5, 4}, //b
		{7, 6, 5}, //c
	}

	expected1 := []*EDInfo{
		{
			Op:         ED_OP_ADD,
			NewComment: &schema.Comment{MD5: "a"},
		},
		{
			Op:          ED_OP_SAME,
			NewComment:  &schema.Comment{MD5: "b"},
			OrigComment: &schema.CommentMD5{MD5: "b"},
		},
		{
			Op:          ED_OP_SAME,
			NewComment:  &schema.Comment{MD5: "c"},
			OrigComment: &schema.CommentMD5{MD5: "c"},
		},
		{
			Op:         ED_OP_ADD,
			NewComment: &schema.Comment{MD5: "d"},
		},
		{
			Op:         ED_OP_ADD,
			NewComment: &schema.Comment{MD5: "a"},
		},
		{
			Op:         ED_OP_ADD,
			NewComment: &schema.Comment{MD5: "b"},
		},
		{
			Op:         ED_OP_ADD,
			NewComment: &schema.Comment{MD5: "c"},
		},
	}

	comments2 := []*schema.Comment{
		{MD5: "a"},
		{MD5: "b"},
		{MD5: "c"},
		{MD5: "d"},
		{MD5: "a"},
		{MD5: "b"},
		{MD5: "a"},
		{MD5: "c"},
	}

	commentMD5s2 := []*schema.CommentMD5{
		{MD5: "b"},
		{MD5: "e"},
		{MD5: "a"},
		{MD5: "c"},
	}

	scoreMat2 := [][]int{
		//  b  e  a, c
		{0, 1, 2, 3, 4}, //
		{1, 2, 3, 2, 3}, //a
		{2, 1, 2, 3, 4}, //b
		{3, 2, 3, 4, 3}, //c
		{4, 3, 4, 5, 4}, //d
		{5, 4, 5, 4, 5}, //a
		{6, 5, 6, 5, 6}, //b
		{7, 6, 7, 6, 7}, //a
		{8, 7, 8, 7, 6}, //c
	}

	expected2 := []*EDInfo{
		{
			Op:         ED_OP_ADD,
			NewComment: &schema.Comment{MD5: "a"},
		},
		{
			Op:          ED_OP_SAME,
			NewComment:  &schema.Comment{MD5: "b"},
			OrigComment: &schema.CommentMD5{MD5: "b"},
		},
		{
			Op:          ED_OP_DELETE,
			OrigComment: &schema.CommentMD5{MD5: "e"},
		},
		{
			Op:         ED_OP_ADD,
			NewComment: &schema.Comment{MD5: "c"},
		},
		{
			Op:         ED_OP_ADD,
			NewComment: &schema.Comment{MD5: "d"},
		},
		{
			Op:          ED_OP_SAME,
			NewComment:  &schema.Comment{MD5: "a"},
			OrigComment: &schema.CommentMD5{MD5: "a"},
		},
		{
			Op:         ED_OP_ADD,
			NewComment: &schema.Comment{MD5: "b"},
		},
		{
			Op:         ED_OP_ADD,
			NewComment: &schema.Comment{MD5: "a"},
		},
		{
			Op:          ED_OP_SAME,
			NewComment:  &schema.Comment{MD5: "c"},
			OrigComment: &schema.CommentMD5{MD5: "c"},
		},
	}

	type args struct {
		scoreMat       [][]int
		comments       []*schema.Comment
		allCommentMD5s []*schema.CommentMD5
	}
	tests := []struct {
		name            string
		args            args
		expectedEdInfos []*EDInfo
	}{
		// TODO: Add test cases.
		{
			args:            args{scoreMat: scoreMat0, comments: comments0, allCommentMD5s: commentMD5s0},
			expectedEdInfos: expected0,
		},
		{
			args:            args{scoreMat: scoreMat1, comments: comments1, allCommentMD5s: commentMD5s1},
			expectedEdInfos: expected1,
		},
		{
			args:            args{scoreMat: scoreMat2, comments: comments2, allCommentMD5s: commentMD5s2},
			expectedEdInfos: expected2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEdInfos := calcEditDistanceBacktracking(tt.args.scoreMat, tt.args.comments, tt.args.allCommentMD5s)

			testutil.TDeepEqual(t, "edInfos", gotEdInfos, tt.expectedEdInfos)
		})
	}
}

func Test_calcEDInfoMetas(t *testing.T) {
	setupTest()
	defer teardownTest()

	edInfo0 := []*EDInfo{
		{
			Op:          ED_OP_DELETE,
			OrigComment: &schema.CommentMD5{MD5: "r"},
		},
		{
			Op:         ED_OP_ADD,
			NewComment: &schema.Comment{MD5: "h"},
		},
		{
			Op:          ED_OP_SAME,
			NewComment:  &schema.Comment{MD5: "o"},
			OrigComment: &schema.CommentMD5{MD5: "o", SortTime: 1234567891000000000},
		},
		{
			Op:         ED_OP_ADD,
			NewComment: &schema.Comment{MD5: "r"},
		},
		{
			Op:          ED_OP_SAME,
			NewComment:  &schema.Comment{MD5: "s"},
			OrigComment: &schema.CommentMD5{MD5: "s", SortTime: 1234567892000000000},
		},
		{
			Op:         ED_OP_ADD,
			NewComment: &schema.Comment{MD5: "e"},
		},
	}

	expected0 := []*EDInfoMeta{
		{
			StartIdx:    0,
			EndIdx:      2,
			StartNanoTS: 1234567890000000000,
			EndNanoTS:   1234567891000000000,
		},
		{
			StartIdx:    3,
			EndIdx:      4,
			StartNanoTS: 1234567891000000000,
			EndNanoTS:   1234567892000000000,
		},
		{
			StartIdx:    5,
			EndIdx:      6,
			StartNanoTS: 1234567892000000000,
			EndNanoTS:   1234567900000000000,
		},
	}

	type args struct {
		edInfos     []*EDInfo
		startNanoTS types.NanoTS
		endNanoTS   types.NanoTS
	}
	tests := []struct {
		name     string
		args     args
		expected []*EDInfoMeta
	}{
		// TODO: Add test cases.
		{
			args: args{edInfos: edInfo0, startNanoTS: 1234567890000000000, endNanoTS: 1234567900000000000},

			expected: expected0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcEDInfoMetas(tt.args.edInfos, tt.args.startNanoTS, tt.args.endNanoTS)
			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
	}
}

func TestEDInfoMeta_ToEDBlock(t *testing.T) {
	setupTest()
	defer teardownTest()

	edInfos0 := []*EDInfo{
		{
			Op:          ED_OP_DELETE,
			OrigComment: &schema.CommentMD5{MD5: "r"},
			SortTime:    1234567891000000000,
		},
		{
			Op:         ED_OP_ADD,
			NewComment: &schema.Comment{MD5: "h"},
		},
		{
			Op:          ED_OP_SAME,
			NewComment:  &schema.Comment{MD5: "o"},
			OrigComment: &schema.CommentMD5{MD5: "o", SortTime: 1234567892000000000},
		},
		{
			Op:         ED_OP_ADD,
			NewComment: &schema.Comment{MD5: "r"},
		},
		{
			Op:          ED_OP_SAME,
			NewComment:  &schema.Comment{MD5: "s"},
			OrigComment: &schema.CommentMD5{MD5: "s", SortTime: 1234567893000000000},
		},
		{
			Op:         ED_OP_ADD,
			NewComment: &schema.Comment{MD5: "e"},
		},
	}

	edInfoMeta0 := &EDInfoMeta{
		StartIdx:    0,
		EndIdx:      2,
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1234567892000000000,
	}
	expected0 := &EDBlock{
		NewComments: []*EDInfo{
			{Op: ED_OP_ADD, NewComment: &schema.Comment{MD5: "h"}},
		},
		OrigComments: []*EDInfo{
			{Op: ED_OP_DELETE, OrigComment: &schema.CommentMD5{MD5: "r"}, SortTime: 1234567891000000000},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1234567892000000000,
	}
	edInfoMeta1 := &EDInfoMeta{
		StartIdx:    3,
		EndIdx:      4,
		StartNanoTS: 1234567892000000000,
		EndNanoTS:   1234567893000000000,
	}
	expected1 := &EDBlock{
		NewComments: []*EDInfo{
			{Op: ED_OP_ADD, NewComment: &schema.Comment{MD5: "r"}},
		},
		StartNanoTS: 1234567892000000000,
		EndNanoTS:   1234567893000000000,
	}
	edInfoMeta2 := &EDInfoMeta{
		StartIdx:    5,
		EndIdx:      6,
		StartNanoTS: 1234567893000000000,
		EndNanoTS:   1234567900000000000,
	}
	expected2 := &EDBlock{
		NewComments: []*EDInfo{
			{Op: ED_OP_ADD, NewComment: &schema.Comment{MD5: "e"}},
		},
		StartNanoTS: 1234567893000000000,
		EndNanoTS:   1234567900000000000,
	}

	type fields struct {
		StartNanoTS types.NanoTS
		EndNanoTS   types.NanoTS
		StartIdx    int
		EndIdx      int
	}
	type args struct {
		edInfos []*EDInfo
	}
	tests := []struct {
		name       string
		edInfoMeta *EDInfoMeta
		args       args
		expected   *EDBlock
	}{
		// TODO: Add test cases.
		{
			edInfoMeta: edInfoMeta0,
			args:       args{edInfos: edInfos0},
			expected:   expected0,
		},
		{
			edInfoMeta: edInfoMeta1,
			args:       args{edInfos: edInfos0},
			expected:   expected1,
		},
		{
			edInfoMeta: edInfoMeta2,
			args:       args{edInfos: edInfos0},
			expected:   expected2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			meta := tt.edInfoMeta
			got := meta.ToEDBlock(tt.args.edInfos)
			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
	}
}

func TestCalcEDBlocks(t *testing.T) {
	setupTest()
	defer teardownTest()

	comments0 := []*schema.Comment{
		{MD5: "h"},
		{MD5: "o"},
		{MD5: "r"},
		{MD5: "s"},
		{MD5: "e", SortTime: 1234567894000000000},
	}

	commentMD5s0 := []*schema.CommentMD5{
		{MD5: "r", SortTime: 1234567891000000000},
		{MD5: "o", SortTime: 1234567892000000000},
		{MD5: "s", SortTime: 1234567893000000000},
	}

	expected0 := []*EDBlock{
		{
			NewComments: []*EDInfo{
				{
					Op:         ED_OP_ADD,
					NewComment: &schema.Comment{MD5: "h"},
				},
			},
			OrigComments: []*EDInfo{
				{
					Op: ED_OP_DELETE,
					OrigComment: &schema.CommentMD5{
						MD5:      "r",
						SortTime: 1234567891000000000,
					},
					SortTime: 1234567891000000000,
				},
			},
			StartNanoTS: 1234567890000000000,
			EndNanoTS:   1234567892000000000,
		},
		{
			NewComments: []*EDInfo{
				{
					Op:         ED_OP_ADD,
					NewComment: &schema.Comment{MD5: "r"},
				},
			},
			StartNanoTS: 1234567892000000000,
			EndNanoTS:   1234567893000000000,
		},
		{
			NewComments: []*EDInfo{
				{
					Op: ED_OP_ADD,
					NewComment: &schema.Comment{
						MD5:      "e",
						SortTime: 1234567894000000000,
					},
					SortTime: 1234567894000000000,
				},
			},
			StartNanoTS: 1234567893000000000,
			EndNanoTS:   1234567900000000000,
		},
	}

	type args struct {
		newComments       []*schema.Comment
		origComments      []*schema.CommentMD5
		articleCreateTime types.NanoTS
		articleMTime      types.NanoTS
	}
	tests := []struct {
		name     string
		args     args
		expected []*EDBlock
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				newComments:       comments0,
				origComments:      commentMD5s0,
				articleCreateTime: 1234567890000000000,
				articleMTime:      1234567900000000000,
			},
			expected: expected0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEdBlocks, err := CalcEDBlocks(tt.args.newComments, tt.args.origComments, tt.args.articleCreateTime, tt.args.articleMTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalcEDBlocks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotEdBlocks, tt.expected)
		})
	}
}

func TestEDBlock_AlignEndNanoTS(t *testing.T) {
	setupTest()
	defer teardownTest()

	edBlock0 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_DELETED,
					MD5:        "m4",
					SortTime:   1234576900000010000,
					CreateTime: 1234576900000010000,
					RefIDs:     []types.CommentID{"temp:m3"},
					CommentID:  "ESIZJkRnTxA:m4",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m0",
					TheDate: "02/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_BOO,
					MD5:     "m1",
					TheDate: "02/14 10:31",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_COMMENT,
					MD5:     "m2",
					TheDate: "02/14 10:32",
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1234578720000000000,
	}

	expected0 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_DELETED,
					MD5:        "m4",
					SortTime:   1234576900000010000,
					CreateTime: 1234576900000010000,
					RefIDs:     []types.CommentID{"temp:m3"},
					CommentID:  "ESIZJkRnTxA:m4",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m0",
					TheDate: "02/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_BOO,
					MD5:     "m1",
					TheDate: "02/14 10:31",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_COMMENT,
					MD5:        "m2",
					TheDate:    "02/14 10:32",
					SortTime:   1234578720000000000,
					CreateTime: 1234578720000000000,
					CommentID:  "ESIazgTbQAA:m2",
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1234578720000000000,
	}

	type fields struct {
		NewComments  []*EDInfo
		OrigComments []*EDInfo
		StartNanoTS  types.NanoTS
		EndNanoTS    types.NanoTS
	}
	tests := []struct {
		name     string
		edBlock  *EDBlock
		expected *EDBlock
	}{
		// TODO: Add test cases.
		{
			edBlock:  edBlock0,
			expected: expected0,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			ed := tt.edBlock
			ed.AlignEndNanoTS()

			testutil.TDeepEqual(t, "got", ed, tt.expected)
		})
	}
	wg.Wait()
}

func TestEDBlock_ForwardInferTS(t *testing.T) {
	setupTest()
	defer teardownTest()

	edBlock0 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_DELETED,
					MD5:     "m4",
					RefIDs:  []types.CommentID{"temp:m3"},
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m0",
					TheDate: "02/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_BOO,
					MD5:     "m1",
					TheDate: "03/14 10:31",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_COMMENT,
					MD5:     "m2",
					TheDate: "04/14 10:32",
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1240200000000000000,
	}

	expected0 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_DELETED,
					MD5:        "m4",
					SortTime:   1234567890001000000,
					CreateTime: 1234567890001000000,
					RefIDs:     []types.CommentID{"temp:m3"},
					CommentID:  "ESIQ9Hac9kA:m4",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m0",
					TheDate:    "02/14 10:30",
					CreateTime: 1234578600000000000,
					SortTime:   1234578600000000000,
					CommentID:  "ESIashRMkAA:m0",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_BOO,
					MD5:        "m1",
					TheDate:    "03/14 10:31",
					CommentID:  "ESqy__E36AA:m1",
					CreateTime: 1236997860000000000,
					SortTime:   1236997860000000000,
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_COMMENT,
					MD5:        "m2",
					TheDate:    "04/14 10:32",
					CommentID:  "ETQ3C4IQQAA:m2",
					CreateTime: 1239676320000000000,
					SortTime:   1239676320000000000,
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1240200000000000000,
	}

	edBlock1 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m0",
					TheDate: "02/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_BOO,
					MD5:     "m1",
					TheDate: "03/14 10:31",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_COMMENT,
					MD5:     "m2",
					TheDate: "02/14 10:32",
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1297686600000000000,
	}

	expected1 := &EDBlock{
		NewComments: []*EDInfo{

			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m0",
					TheDate:    "02/14 10:30",
					CreateTime: 1234578600000000000,
					SortTime:   1234578600000000000,
					CommentID:  "ESIashRMkAA:m0",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_BOO,
					MD5:        "m1",
					TheDate:    "03/14 10:31",
					CommentID:  "ESqy__E36AA:m1",
					CreateTime: 1236997860000000000,
					SortTime:   1236997860000000000,
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_COMMENT,
					MD5:     "m2",
					TheDate: "02/14 10:32",
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1297686600000000000,
	}

	edBlock2 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m0",
					TheDate: "02/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_BOO,
					MD5:     "m1",
					TheDate: "02/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_EDIT,
					MD5:        "m2",
					TheDate:    "02/14/2009 10:32",
					CommentID:  "temp:m2",
					SortTime:   1234578720000000000,
					CreateTime: 1234578720000000000,
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1297686600000000000,
	}

	expected2 := &EDBlock{
		NewComments: []*EDInfo{

			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m0",
					TheDate:    "02/14 10:30",
					CreateTime: 1234578600000000000,
					SortTime:   1234578600000000000,
					CommentID:  "ESIashRMkAA:m0",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_BOO,
					MD5:        "m1",
					TheDate:    "02/14 10:30",
					CommentID:  "ESIashRb0kA:m1",
					CreateTime: 1234578600000000000,
					SortTime:   1234578600001000000,
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_EDIT,
					MD5:        "m2",
					TheDate:    "02/14/2009 10:32",
					CommentID:  "temp:m2",
					CreateTime: 1234578720000000000,
					SortTime:   1234578720000000000,
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1297686600000000000,
	}

	edBlock3 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m0",
					TheDate: "02/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_BOO,
					MD5:     "m1",
					TheDate: "03/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m2",
					TheDate: "05/14 10:32",
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1240000000000000000,
	}

	expected3 := &EDBlock{
		NewComments: []*EDInfo{

			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m0",
					TheDate:    "02/14 10:30",
					CreateTime: 1234578600000000000,
					SortTime:   1234578600000000000,
					CommentID:  "ESIashRMkAA:m0",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_BOO,
					MD5:        "m1",
					TheDate:    "03/14 10:30",
					CommentID:  "ESqy8fjwkAA:m1",
					CreateTime: 1236997800000000000,
					SortTime:   1236997800000000000,
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m2",
					TheDate:    "05/14 10:32",
					CommentID:  "ETVdbiF7_Bg:m2",
					CreateTime: 1242268320000000000,
					SortTime:   1239999999999999000,
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1240000000000000000,
	}

	edBlock4 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_REPLY,
					MD5:        "m0",
					TheDate:    "04/14/2009 10:30",
					CommentID:  "temp:m0",
					CreateTime: 1239676320000000000,
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_BOO,
					MD5:     "m1",
					TheDate: "03/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_REPLY,
					MD5:        "m2",
					TheDate:    "04/14/2009 10:30",
					CommentID:  "temp:m2",
					CreateTime: 1239676320000000000,
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1240000000000000000,
	}

	expected4 := &EDBlock{
		NewComments: []*EDInfo{

			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_REPLY,
					MD5:        "m0",
					TheDate:    "04/14/2009 10:30",
					CreateTime: 1239676320000000000,
					SortTime:   1234567890000100000,
					CommentID:  "ESIQ9HaPOqA:m0",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_BOO,
					MD5:        "m1",
					TheDate:    "03/14 10:30",
					CommentID:  "ESqy8fjwkAA:m1",
					CreateTime: 1236997800000000000,
					SortTime:   1236997800000000000,
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_REPLY,
					MD5:        "m2",
					TheDate:    "04/14/2009 10:30",
					CommentID:  "ESqy8fjyFqA:m2",
					CreateTime: 1239676320000000000,
					SortTime:   1236997800000100000,
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1240000000000000000,
	}

	type fields struct {
		NewComments  []*EDInfo
		OrigComments []*EDInfo
		StartNanoTS  types.NanoTS
		EndNanoTS    types.NanoTS
	}
	type args struct {
		startNanoTS types.NanoTS
	}
	tests := []struct {
		name            string
		edBlock         *EDBlock
		args            args
		expected        *EDBlock
		expectedNextIdx int
	}{
		// TODO: Add test cases.
		{
			args:            args{startNanoTS: 1234567890000000000},
			edBlock:         edBlock0,
			expected:        expected0,
			expectedNextIdx: 4,
		},
		{
			name:            "exceeding 1-year",
			args:            args{startNanoTS: 1234567890000000000},
			edBlock:         edBlock1,
			expected:        expected1,
			expectedNextIdx: 2,
		},
		{
			name:            "same year",
			args:            args{startNanoTS: 1234567890000000000},
			edBlock:         edBlock2,
			expected:        expected2,
			expectedNextIdx: 3,
		},
		{
			name:            "exceeding endNanoTS",
			args:            args{startNanoTS: 1234567890000000000},
			edBlock:         edBlock3,
			expected:        expected3,
			expectedNextIdx: 3,
		},
		{
			name:            "reply",
			args:            args{startNanoTS: 1234567890000000000},
			edBlock:         edBlock4,
			expected:        expected4,
			expectedNextIdx: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ed := tt.edBlock
			if gotNextIdx := ed.ForwardInferTS(tt.args.startNanoTS); gotNextIdx != tt.expectedNextIdx {
				t.Errorf("EDBlock.ForwardInferTS() = %v, want %v", gotNextIdx, tt.expectedNextIdx)
			}

			testutil.TDeepEqual(t, "got", tt.edBlock, tt.expected)
		})
	}
}

func TestEDBlock_BackwardInferTS(t *testing.T) {
	setupTest()
	defer teardownTest()

	edBlock0 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_DELETED,
					MD5:     "m4",
					RefIDs:  []types.CommentID{"temp:m3"},
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m0",
					TheDate: "02/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_BOO,
					MD5:     "m1",
					TheDate: "03/14 10:31",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_COMMENT,
					MD5:     "m2",
					TheDate: "04/14 10:32",
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1240200000000000000,
	}

	expected0 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_DELETED,
					MD5:        "m4",
					SortTime:   1234578600899000000,
					CreateTime: 1234578600899000000,
					RefIDs:     []types.CommentID{"temp:m3"},
					CommentID:  "ESIaskniNsA:m4",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m0",
					TheDate:    "02/14 10:30",
					CreateTime: 1234578600000000000,
					SortTime:   1234578600900000000,
					CommentID:  "ESIasknxeQA:m0",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_BOO,
					MD5:        "m1",
					TheDate:    "03/14 10:31",
					CommentID:  "ESqzACbc0QA:m1",
					CreateTime: 1236997860000000000,
					SortTime:   1236997860900000000,
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_COMMENT,
					MD5:        "m2",
					TheDate:    "04/14 10:32",
					CommentID:  "ETQ3C7e1KQA:m2",
					CreateTime: 1239676320000000000,
					SortTime:   1239676320900000000,
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1240200000000000000,
	}

	edBlock1 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_DELETED,
					MD5:     "m4",
					RefIDs:  []types.CommentID{"temp:m3"},
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m0",
					TheDate: "02/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_BOO,
					MD5:     "m1",
					TheDate: "03/14 10:31",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_COMMENT,
					MD5:     "m2",
					TheDate: "04/14 10:32",
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1271476800000000000,
	}

	expected1 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_DELETED,
					MD5:        "m4",
					SortTime:   1266114600899000000,
					CreateTime: 1266114600899000000,
					RefIDs:     []types.CommentID{"temp:m3"},
					CommentID:  "EZIkhXeFNsA:m4",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m0",
					TheDate:    "02/14 10:30",
					CreateTime: 1266114600000000000,
					SortTime:   1266114600900000000,
					CommentID:  "EZIkhXeUeQA:m0",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_BOO,
					MD5:        "m1",
					TheDate:    "03/14 10:31",
					CommentID:  "EZq801R_0QA:m1",
					CreateTime: 1268533860000000000,
					SortTime:   1268533860900000000,
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_COMMENT,
					MD5:        "m2",
					TheDate:    "04/14 10:32",
					CommentID:  "EaRA3uVYKQA:m2",
					CreateTime: 1271212320000000000,
					SortTime:   1271212320900000000,
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1271476800000000000,
	}

	edBlock2 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m4",
					TheDate: "05/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m0",
					TheDate: "02/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_BOO,
					MD5:     "m1",
					TheDate: "03/14 10:31",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_COMMENT,
					MD5:     "m2",
					TheDate: "04/14 10:32",
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1271476800000000000,
	}

	expected2 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m4",
					TheDate:    "05/14 10:30",
					CreateTime: 1242268200000000000,
					SortTime:   1242268200900000000,
					CommentID:  "ET1sWM5oeQA:m4",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m0",
					TheDate:    "02/14 10:30",
					CreateTime: 1266114600000000000,
					SortTime:   1266114600900000000,
					CommentID:  "EZIkhXeUeQA:m0",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_BOO,
					MD5:        "m1",
					TheDate:    "03/14 10:31",
					CommentID:  "EZq801R_0QA:m1",
					CreateTime: 1268533860000000000,
					SortTime:   1268533860900000000,
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_COMMENT,
					MD5:        "m2",
					TheDate:    "04/14 10:32",
					CommentID:  "EaRA3uVYKQA:m2",
					CreateTime: 1271212320000000000,
					SortTime:   1271212320900000000,
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1271476800000000000,
	}

	edBlock3 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m4",
					TheDate: "05/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m0",
					TheDate: "02/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_BOO,
					MD5:     "m1",
					TheDate: "03/14 10:31",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_COMMENT,
					MD5:     "m2",
					TheDate: "04/17 12:00",
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1271476800000000000,
	}

	expected3 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m4",
					TheDate:    "05/14 10:30",
					CreateTime: 1242268200000000000,
					SortTime:   1242268200900000000,
					CommentID:  "ET1sWM5oeQA:m4",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m0",
					TheDate:    "02/14 10:30",
					CreateTime: 1266114600000000000,
					SortTime:   1266114600900000000,
					CommentID:  "EZIkhXeUeQA:m0",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_BOO,
					MD5:        "m1",
					TheDate:    "03/14 10:31",
					CommentID:  "EZq801R_0QA:m1",
					CreateTime: 1268533860000000000,
					SortTime:   1268533860900000000,
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_COMMENT,
					MD5:        "m2",
					TheDate:    "04/17 12:00",
					CommentID:  "EaUxabwXPcA:m2",
					CreateTime: 1271476800000000000,
					SortTime:   1271476799999000000,
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1271476800000000000,
	}

	edBlock4 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m4",
					TheDate: "05/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m0",
					TheDate: "02/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_BOO,
					MD5:     "m1",
					TheDate: "03/14 10:31",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_COMMENT,
					MD5:     "m2",
					TheDate: "04/17 12:00",
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1271476830000000000,
	}

	expected4 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m4",
					TheDate:    "05/14 10:30",
					CreateTime: 1242268200000000000,
					SortTime:   1242268200900000000,
					CommentID:  "ET1sWM5oeQA:m4",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m0",
					TheDate:    "02/14 10:30",
					CreateTime: 1266114600000000000,
					SortTime:   1266114600900000000,
					CommentID:  "EZIkhXeUeQA:m0",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_BOO,
					MD5:        "m1",
					TheDate:    "03/14 10:31",
					CommentID:  "EZq801R_0QA:m1",
					CreateTime: 1268533860000000000,
					SortTime:   1268533860900000000,
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_COMMENT,
					MD5:        "m2",
					TheDate:    "04/17 12:00",
					CommentID:  "EaUxafHLaQA:m2",
					CreateTime: 1271476800000000000,
					SortTime:   1271476800900000000,
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1271476830000000000,
	}

	edBlock5 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m4",
					TheDate: "05/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m0",
					TheDate: "02/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_BOO,
					MD5:     "m1",
					TheDate: "03/14 10:31",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_COMMENT,
					MD5:     "m2",
					TheDate: "04/17 12:00",
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1271476830000000000,
	}

	expected5 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m4",
					TheDate:    "05/14 10:30",
					CreateTime: 1242268200000000000,
					SortTime:   1242268200900000000,
					CommentID:  "ET1sWM5oeQA:m4",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m0",
					TheDate:    "02/14 10:30",
					CreateTime: 1266114600000000000,
					SortTime:   1266114600900000000,
					CommentID:  "EZIkhXeUeQA:m0",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_BOO,
					MD5:        "m1",
					TheDate:    "03/14 10:31",
					CommentID:  "EZq801R_0QA:m1",
					CreateTime: 1268533860000000000,
					SortTime:   1268533860900000000,
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_COMMENT,
					MD5:        "m2",
					TheDate:    "04/17 12:00",
					CommentID:  "EaUxcLhKLAA:m2",
					CreateTime: 1271476830000000000,
					SortTime:   1271476830000000000,
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1271476830000000000,
	}

	edBlock6 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m4",
					TheDate: "01/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m0",
					TheDate: "02/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_BOO,
					MD5:     "m1",
					TheDate: "03/14 10:31",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_COMMENT,
					MD5:     "m2",
					TheDate: "04/14 10:32",
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1240200000000000000,
	}

	expected6 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m4",
					TheDate:    "01/14 10:30",
					CreateTime: 1231900200000000000,
					SortTime:   1234567890000001000,
					CommentID:  "ESIQ9HaNt-g:m4",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m0",
					TheDate:    "02/14 10:30",
					CreateTime: 1234578600000000000,
					SortTime:   1234578600900000000,
					CommentID:  "ESIasknxeQA:m0",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_BOO,
					MD5:        "m1",
					TheDate:    "03/14 10:31",
					CommentID:  "ESqzACbc0QA:m1",
					CreateTime: 1236997860000000000,
					SortTime:   1236997860900000000,
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_COMMENT,
					MD5:        "m2",
					TheDate:    "04/14 10:32",
					CommentID:  "ETQ3C7e1KQA:m2",
					CreateTime: 1239676320000000000,
					SortTime:   1239676320900000000,
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1240200000000000000,
	}

	edBlock7 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_REPLY,
					MD5:     "m4",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m0",
					TheDate: "02/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_BOO,
					MD5:     "m1",
					TheDate: "03/14 10:31",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_REPLY,
					MD5:     "m1",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_COMMENT,
					MD5:     "m2",
					TheDate: "04/14 10:32",
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1271476800000000000,
	}

	expected7 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_REPLY,
					MD5:        "m4",
					SortTime:   1234567890000100000,
					CreateTime: 1234567890000100000,
					CommentID:  "ESIQ9HaPOqA:m4",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m0",
					TheDate:    "02/14 10:30",
					CreateTime: 1266114600000000000,
					SortTime:   1266114600900000000,
					CommentID:  "EZIkhXeUeQA:m0",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_BOO,
					MD5:        "m1",
					TheDate:    "03/14 10:31",
					CommentID:  "EZq801R_0QA:m1",
					CreateTime: 1268533860000000000,
					SortTime:   1268533860900000000,
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_REPLY,
					MD5:        "m1",
					CommentID:  "EZq801SBV6A:m1",
					CreateTime: 1268533860900100000,
					SortTime:   1268533860900100000,
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_COMMENT,
					MD5:        "m2",
					TheDate:    "04/14 10:32",
					CommentID:  "EaRA3uVYKQA:m2",
					CreateTime: 1271212320000000000,
					SortTime:   1271212320900000000,
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1271476800000000000,
	}

	type fields struct {
		NewComments  []*EDInfo
		OrigComments []*EDInfo
		StartNanoTS  types.NanoTS
		EndNanoTS    types.NanoTS
	}
	type args struct {
		nextIdx          int
		isAlignEndNanoTS bool
	}
	tests := []struct {
		name     string
		edBlock  *EDBlock
		expected *EDBlock
		args     args
	}{
		// TODO: Add test cases.
		{
			args:     args{nextIdx: 0, isAlignEndNanoTS: false},
			edBlock:  edBlock0,
			expected: expected0,
		},
		{
			name:     "2010",
			args:     args{nextIdx: 0, isAlignEndNanoTS: false},
			edBlock:  edBlock1,
			expected: expected1,
		},
		{
			name:     "exceeding 1-year",
			args:     args{nextIdx: 0, isAlignEndNanoTS: false},
			edBlock:  edBlock2,
			expected: expected2,
		},
		{
			name:     "same with endNanoTS",
			args:     args{nextIdx: 0, isAlignEndNanoTS: false},
			edBlock:  edBlock3,
			expected: expected3,
		},
		{
			name:     "same with endNanoTS-2",
			args:     args{nextIdx: 0, isAlignEndNanoTS: false},
			edBlock:  edBlock4,
			expected: expected4,
		},
		{
			name:     "align endNanoTS",
			args:     args{nextIdx: 0, isAlignEndNanoTS: true},
			edBlock:  edBlock5,
			expected: expected5,
		},
		{
			name:     "exceeding startNanoTS",
			args:     args{nextIdx: 0, isAlignEndNanoTS: false},
			edBlock:  edBlock6,
			expected: expected6,
		},
		{
			name:     "reply",
			args:     args{nextIdx: 0, isAlignEndNanoTS: false},
			edBlock:  edBlock7,
			expected: expected7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			edBlock := tt.edBlock
			edBlock.BackwardInferTS(tt.args.nextIdx, tt.args.isAlignEndNanoTS)
			testutil.TDeepEqual(t, "got", tt.edBlock, tt.expected)
		})
	}
}

func TestInferTimestamp(t *testing.T) {
	setupTest()
	defer teardownTest()

	edBlock0 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m4",
					TheDate: "05/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m0",
					TheDate: "02/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_BOO,
					MD5:     "m1",
					TheDate: "03/14 10:31",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_COMMENT,
					MD5:     "m2",
					TheDate: "04/17 12:00",
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1303012830000000000,
	}

	expected0 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m4",
					TheDate:    "05/14 10:30",
					CreateTime: 1242268200000000000,
					SortTime:   1242268200000000000,
					CommentID:  "ET1sWJjDkAA:m4",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m0",
					TheDate:    "02/14 10:30",
					CreateTime: 1297650600000000000,
					SortTime:   1297650600900000000,
					CommentID:  "EgIuWKU3eQA:m0",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_BOO,
					MD5:        "m1",
					TheDate:    "03/14 10:31",
					CommentID:  "EgrGpoIi0QA:m1",
					CreateTime: 1300069860000000000,
					SortTime:   1300069860900000000,
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_COMMENT,
					MD5:        "m2",
					TheDate:    "04/17 12:00",
					CommentID:  "EhU7PR9uaQA:m2",
					CreateTime: 1303012800000000000,
					SortTime:   1303012800900000000,
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1303012830000000000,
	}

	edBlock1 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m4",
					TheDate: "05/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m0",
					TheDate: "02/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_BOO,
					MD5:     "m1",
					TheDate: "03/14 10:31",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_COMMENT,
					MD5:     "m2",
					TheDate: "04/17 12:00",
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1303012830000000000,
	}

	expected1 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m4",
					TheDate:    "05/14 10:30",
					CreateTime: 1242268200000000000,
					SortTime:   1242268200000000000,
					CommentID:  "ET1sWJjDkAA:m4",
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1303012830000000000,
	}

	edBlock2 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m4",
					TheDate: "05/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m0",
					TheDate: "02/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_BOO,
					MD5:     "m1",
					TheDate: "03/14 10:31",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_COMMENT,
					MD5:     "m2",
					TheDate: "04/17 12:00",
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1303012830000000000,
	}

	expected2 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m4",
					TheDate:    "05/14 10:30",
					CreateTime: 1242268200000000000,
					SortTime:   1242268200000000000,
					CommentID:  "ET1sWJjDkAA:m4",
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1303012830000000000,
	}

	edBlock3 := &EDBlock{
		NewComments: []*EDInfo{

			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_RECOMMEND,
					MD5:     "m0",
					TheDate: "02/14 10:30",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_BOO,
					MD5:     "m1",
					TheDate: "03/14 10:31",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType: types.COMMENT_TYPE_COMMENT,
					MD5:     "m2",
					TheDate: "04/17 12:00",
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1303012830000000000,
	}

	expected3 := &EDBlock{
		NewComments: []*EDInfo{
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_RECOMMEND,
					MD5:        "m0",
					TheDate:    "02/14 10:30",
					CreateTime: 1234578600000000000,
					SortTime:   1234578600000000000,
					CommentID:  "ESIashRMkAA:m0",
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_BOO,
					MD5:        "m1",
					TheDate:    "03/14 10:31",
					CommentID:  "ESqy__E36AA:m1",
					CreateTime: 1236997860000000000,
					SortTime:   1236997860000000000,
				},
			},
			{
				Op: ED_OP_ADD,
				NewComment: &schema.Comment{
					TheType:    types.COMMENT_TYPE_COMMENT,
					MD5:        "m2",
					TheDate:    "04/17 12:00",
					CommentID:  "EhU7Q-XtLAA:m2",
					CreateTime: 1303012830000000000,
					SortTime:   1303012830000000000,
				},
			},
		},
		OrigComments: []*EDInfo{
			{
				Op: ED_OP_DELETE,
				OrigComment: &schema.CommentMD5{
					MD5:       "m3",
					TheDate:   "02/14 10:01",
					CommentID: "temp:m3",
					SortTime:  1234576900000000000,
				},
			},
		},
		StartNanoTS: 1234567890000000000,
		EndNanoTS:   1303012830000000000,
	}

	type args struct {
		edBlocks             []*EDBlock
		isForwardOnly        bool
		isLastAlignEndNanoTS bool
	}
	tests := []struct {
		name     string
		args     args
		expected []*EDBlock
	}{
		// TODO: Add test cases.
		{
			args:     args{edBlocks: []*EDBlock{edBlock0}},
			expected: []*EDBlock{expected0},
		},
		{
			name:     "forward only",
			args:     args{edBlocks: []*EDBlock{edBlock1}, isForwardOnly: true},
			expected: []*EDBlock{expected1},
		},
		{
			name:     "forward only, last-align-endNanoTS, no effective",
			args:     args{edBlocks: []*EDBlock{edBlock2}, isForwardOnly: true, isLastAlignEndNanoTS: true},
			expected: []*EDBlock{expected2},
		},
		{
			name:     "forward only, last-align-endNanoTS, effective",
			args:     args{edBlocks: []*EDBlock{edBlock3}, isForwardOnly: true, isLastAlignEndNanoTS: true},
			expected: []*EDBlock{expected3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InferTimestamp(tt.args.edBlocks, tt.args.isForwardOnly, tt.args.isLastAlignEndNanoTS)
			testutil.TDeepEqual(t, "got", tt.args.edBlocks, tt.expected)
		})
	}
}
