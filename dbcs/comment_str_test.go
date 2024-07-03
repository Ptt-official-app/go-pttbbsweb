package dbcs

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/go-pttbbsweb/schema"
)

func TestParseCommentsStr(t *testing.T) {
	setupTest()
	defer teardownTest()
	type args struct {
		ownerID         bbs.UUserID
		commentsDBCS    string
		allCommentsDBCS string
	}
	tests := []struct {
		name             string
		args             args
		expectedComments []*schema.Comment
	}{
		// TODO: Add test cases.
		{
			name: "0_" + testUtf8Filename0 + "_first",
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    string(testUtf8FirstCommentsDBCS0),
				allCommentsDBCS: string(testUtf8Comment0),
			},
			expectedComments: testUtf8FirstComments0,
		},
		{
			name: "1_" + testUtf8Filename1 + "_first",
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    string(testUtf8FirstCommentsDBCS1),
				allCommentsDBCS: string(testUtf8Comment1),
			},
			expectedComments: testUtf8FirstComments1,
		},
		{
			name: "2_" + testUtf8Filename2 + "_first",
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    string(testUtf8FirstCommentsDBCS2),
				allCommentsDBCS: string(testUtf8Comment2),
			},
			expectedComments: testUtf8FirstComments2,
		},
		{
			name: "3_" + testUtf8Filename3 + "_first",
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    string(testUtf8FirstCommentsDBCS3),
				allCommentsDBCS: string(testUtf8Comment3),
			},
			expectedComments: testUtf8FirstComments3,
		},
		{
			name: "4_" + testUtf8Filename4 + "_first",
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    string(testUtf8FirstCommentsDBCS4),
				allCommentsDBCS: string(testUtf8Comment4),
			},
			expectedComments: testUtf8FirstComments4,
		},
		{
			name: "5_" + testUtf8Filename5 + "_first",
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    string(testUtf8FirstCommentsDBCS5),
				allCommentsDBCS: string(testUtf8Comment5),
			},
			expectedComments: testUtf8FirstComments5,
		},
		{
			name: "6_" + testUtf8Filename6 + "_first",
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    string(testUtf8FirstCommentsDBCS6),
				allCommentsDBCS: string(testUtf8Comment6),
			},
			expectedComments: testUtf8FirstComments6,
		},
		{
			name: "6_" + testUtf8Filename6 + "_the_rest",
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    string(testUtf8TheRestCommentsDBCS6),
				allCommentsDBCS: string(testUtf8TheRestCommentsDBCS6),
			},
			expectedComments: testUtf8TheRestComments6,
		},
		{
			name: "7_" + testUtf8Filename7 + "_first",
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    string(testUtf8FirstCommentsDBCS7),
				allCommentsDBCS: string(testUtf8Comment7),
			},
			expectedComments: testUtf8FirstComments7,
		},
		{
			name: "8_" + testUtf8Filename8 + "_first",
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    string(testUtf8FirstCommentsDBCS8),
				allCommentsDBCS: string(testUtf8Comment8),
			},
			expectedComments: testUtf8FirstComments8,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotComments := ParseCommentsStr(tt.args.ownerID, tt.args.commentsDBCS, tt.args.allCommentsDBCS)

			testutil.TDeepEqual(t, "got", gotComments, tt.expectedComments)
		})
		wg.Wait()
	}
}
