package dbcs

import (
	"reflect"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/sirupsen/logrus"
)

func TestParseComments(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		ownerID bbs.UUserID

		commentsDBCS    []byte
		allCommentsDBCS []byte
	}
	tests := []struct {
		name             string
		args             args
		expectedComments []*schema.Comment
	}{
		// TODO: Add test cases.
		{
			name: "0_" + testFilename0,
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    testFirstCommentsDBCS0,
				allCommentsDBCS: testComment0,
			},
			expectedComments: testFirstComments0,
		},
		{
			name: "1_" + testFilename1,
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    testFirstCommentsDBCS1,
				allCommentsDBCS: testComment1,
			},
			expectedComments: testFirstComments1,
		},
		{
			name: "2_" + testFilename2,
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    testFirstCommentsDBCS2,
				allCommentsDBCS: testComment2,
			},
			expectedComments: testFirstComments2,
		},
		{
			name: "3_" + testFilename3,
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    testFirstCommentsDBCS3,
				allCommentsDBCS: testComment3,
			},
			expectedComments: testFirstComments3,
		},

		{
			name: "4_" + testFilename4,
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    testFirstCommentsDBCS4,
				allCommentsDBCS: testComment4,
			},
			expectedComments: testFirstComments4,
		},
		{
			name: "5_" + testFilename5,
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    testFirstCommentsDBCS5,
				allCommentsDBCS: testComment5,
			},
			expectedComments: testFirstComments5,
		},
		{
			name: "6_" + testFilename6,
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    testFirstCommentsDBCS6,
				allCommentsDBCS: testComment6,
			},
			expectedComments: testFirstComments6,
		},
		{
			name: "7_" + testFilename7,
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    testFirstCommentsDBCS7,
				allCommentsDBCS: testComment7,
			},
			expectedComments: testFirstComments7,
		},
		{
			name: "8_" + testFilename8,
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    testFirstCommentsDBCS8,
				allCommentsDBCS: testComment8,
			},
			expectedComments: testFirstComments8,
		},
		{
			name: "9_" + testFilename9,
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    testFirstCommentsDBCS9,
				allCommentsDBCS: testComment9,
			},
			expectedComments: testFirstComments9,
		},
		{
			name: "10_" + testFilename10,
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    testFirstCommentsDBCS10,
				allCommentsDBCS: testComment10,
			},
			expectedComments: testFirstComments10,
		},
		{
			name: "11_" + testFilename11,
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    testFirstCommentsDBCS11,
				allCommentsDBCS: testComment11,
			},
			expectedComments: testFirstComments11,
		},
		{
			name: "11_" + testFilename11 + "_the_rest",
			args: args{
				ownerID:         "cheinshin",
				commentsDBCS:    testTheRestCommentsDBCS11,
				allCommentsDBCS: testTheRestCommentsDBCS11,
			},
			expectedComments: testTheRestComments11,
		},
		{
			name: "12_" + testFilename12,
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    testFirstCommentsDBCS12,
				allCommentsDBCS: testComment12,
			},
			expectedComments: testFirstComments12,
		},
		{
			name: "13_" + testFilename13,
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    testFirstCommentsDBCS13,
				allCommentsDBCS: testComment13,
			},
			expectedComments: testFirstComments13,
		},
		{
			name: "14_" + testFilename14,
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    testFirstCommentsDBCS14,
				allCommentsDBCS: testComment14,
			},
			expectedComments: testFirstComments14,
		},
		{
			name: "14_" + testFilename14 + "_the_rest",
			args: args{
				ownerID:         "cheinshin",
				commentsDBCS:    testTheRestCommentsDBCS14,
				allCommentsDBCS: testTheRestCommentsDBCS14,
			},
			expectedComments: testTheRestComments14,
		},
		{
			name: "15_" + testFilename15,
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    testFirstCommentsDBCS15,
				allCommentsDBCS: testComment15,
			},
			expectedComments: testFirstComments15,
		},
		{
			name: "16_" + testFilename16,
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    testFirstCommentsDBCS16,
				allCommentsDBCS: testComment16,
			},
			expectedComments: testFirstComments16,
		},
		{
			name: "17_" + testFilename17,
			args: args{
				ownerID:         "testOwner",
				commentsDBCS:    testFirstCommentsDBCS17,
				allCommentsDBCS: testComment17,
			},
			expectedComments: testFirstComments17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logrus.Infof("%v: to ParseComments", tt.name)
			gotComments := ParseComments(tt.args.ownerID, tt.args.commentsDBCS, tt.args.allCommentsDBCS)

			testutil.TDeepEqual(t, "comments", gotComments, tt.expectedComments)

		})
	}
}

func TestParseFirstComments(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		bboardID                  bbs.BBoardID
		articleID                 bbs.ArticleID
		ownerID                   bbs.UUserID
		articleCreateTime         types.NanoTS
		commentsDBCS              []byte
		origFirstCommentsMD5      string
		origFirstCommentsLastTime types.NanoTS
		updateNanoTS              types.NanoTS
	}
	tests := []struct {
		name                          string
		args                          args
		expectedFirstComments         []*schema.Comment
		expectedFirstCommentsMD5      string
		expectedFirstCommentsLastTime types.NanoTS
		expectedTheRestComments       []byte
	}{
		// TODO: Add test cases.
		/*
			{
				name: "0_" + testFilename0,
				args: args{
					bboardID:          "test",
					articleID:         "test",
					ownerID:           "testOwner",
					articleCreateTime: types.NanoTS(1234567890000000000),
					commentsDBCS:      testComment0,
					updateNanoTS:      types.NanoTS(1334567890000000000),
				},
				expectedFirstComments:         testFirstComments0,
				expectedFirstCommentsMD5:      "lUNLzf4Qpeos8HBS676eWg",
				expectedFirstCommentsLastTime: types.NanoTS(1260647460000000000),
			},
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirstComments, gotFirstCommentsMD5, gotFirstCommentsLastTime, gotTheRestComments := ParseFirstComments(tt.args.bboardID, tt.args.articleID, tt.args.ownerID, tt.args.articleCreateTime, tt.args.commentsDBCS, tt.args.origFirstCommentsMD5, tt.args.origFirstCommentsLastTime, tt.args.updateNanoTS)
			testutil.TDeepEqual(t, "firstComments", gotFirstComments, tt.expectedFirstComments)

			if gotFirstCommentsMD5 != tt.expectedFirstCommentsMD5 {
				t.Errorf("ParseFirstComments() gotFirstCommentsMD5 = %v, want %v", gotFirstCommentsMD5, tt.expectedFirstCommentsMD5)
			}
			if !reflect.DeepEqual(gotFirstCommentsLastTime, tt.expectedFirstCommentsLastTime) {
				t.Errorf("ParseFirstComments() gotFirstCommentsLastTime = %v, want %v", gotFirstCommentsLastTime, tt.expectedFirstCommentsLastTime)
			}
			if !reflect.DeepEqual(gotTheRestComments, tt.expectedTheRestComments) {
				t.Errorf("ParseFirstComments() gotTheRestComments = %v, want %v", gotTheRestComments, tt.expectedTheRestComments)
			}
		})
	}
}
