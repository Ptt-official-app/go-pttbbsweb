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
		bboardID        bbs.BBoardID
		articleID       bbs.ArticleID
		ownerID         bbs.UUserID
		lastTimeNanoTS  types.NanoTS
		commentsDBCS    []byte
		allCommentsDBCS []byte
		updateNanoTS    types.NanoTS
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
				bboardID:        "test",
				articleID:       "test",
				ownerID:         "testOwner",
				commentsDBCS:    testFirstCommentsDBCS0,
				allCommentsDBCS: testComment0,
				lastTimeNanoTS:  types.NanoTS(1234567890000000000),
				updateNanoTS:    types.NanoTS(1334567890000000000),
			},
			expectedComments: testFirstComments0,
		},
		/*
			{
				name: "1_" + testFilename1,
				args: args{
					bboardID:        "test",
					articleID:       "test",
					ownerID:         "testOwner",
					commentsDBCS:    testFirstCommentsDBCS1,
					allCommentsDBCS: testComment1,
					lastTimeNanoTS:  types.NanoTS(1234567890000000000),
					updateNanoTS:    types.NanoTS(1334567890000000000),
				},
				expectedComments: testFirstComments1,
			},
			{
				name: "2_" + testFilename2,
				args: args{
					bboardID:        "test",
					articleID:       "test",
					ownerID:         "testOwner",
					commentsDBCS:    testFirstCommentsDBCS2,
					allCommentsDBCS: testComment2,
					lastTimeNanoTS:  types.NanoTS(1234567890000000000),
					updateNanoTS:    types.NanoTS(1334567890000000000),
				},
				expectedComments: testFirstComments2,
			},
			{
				name: "3_" + testFilename3,
				args: args{
					bboardID:        "test",
					articleID:       "test",
					ownerID:         "testOwner",
					commentsDBCS:    testFirstCommentsDBCS3,
					allCommentsDBCS: testComment3,
					lastTimeNanoTS:  types.NanoTS(1234567890000000000),
					updateNanoTS:    types.NanoTS(1334567890000000000),
				},
				expectedComments: testFirstComments3,
			},
			{
				name: "4_" + testFilename4,
				args: args{
					bboardID:        "test",
					articleID:       "test",
					ownerID:         "testOwner",
					commentsDBCS:    testFirstCommentsDBCS4,
					allCommentsDBCS: testComment4,
					lastTimeNanoTS:  types.NanoTS(1234567890000000000),
					updateNanoTS:    types.NanoTS(1334567890000000000),
				},
				expectedComments: testFirstComments4,
			},
			{
				name: "5_" + testFilename5,
				args: args{
					bboardID:        "test",
					articleID:       "test",
					ownerID:         "testOwner",
					commentsDBCS:    testFirstCommentsDBCS5,
					allCommentsDBCS: testComment5,
					lastTimeNanoTS:  types.NanoTS(1234567890000000000),
					updateNanoTS:    types.NanoTS(1334567890000000000),
				},
				expectedComments: testFirstComments5,
			},
			{
				name: "6_" + testFilename6,
				args: args{
					bboardID:        "test",
					articleID:       "test",
					ownerID:         "testOwner",
					commentsDBCS:    testFirstCommentsDBCS6,
					allCommentsDBCS: testComment6,
					lastTimeNanoTS:  types.NanoTS(1234567890000000000),
					updateNanoTS:    types.NanoTS(1334567890000000000),
				},
				expectedComments: testFirstComments6,
			},
			{
				name: "7_" + testFilename7,
				args: args{
					bboardID:        "test",
					articleID:       "test",
					ownerID:         "testOwner",
					commentsDBCS:    testFirstCommentsDBCS7,
					allCommentsDBCS: testComment7,
					lastTimeNanoTS:  types.NanoTS(1234567890000000000),
					updateNanoTS:    types.NanoTS(1334567890000000000),
				},
				expectedComments: testFirstComments7,
			},
			{
				name: "8_" + testFilename8,
				args: args{
					bboardID:        "test",
					articleID:       "test",
					ownerID:         "testOwner",
					commentsDBCS:    testFirstCommentsDBCS8,
					allCommentsDBCS: testComment8,
					lastTimeNanoTS:  types.NanoTS(1234567890000000000),
					updateNanoTS:    types.NanoTS(1334567890000000000),
				},
				expectedComments: testFirstComments8,
			},
			{
				name: "9_" + testFilename9,
				args: args{
					bboardID:        "test",
					articleID:       "test",
					ownerID:         "testOwner",
					commentsDBCS:    testFirstCommentsDBCS9,
					allCommentsDBCS: testComment9,
					lastTimeNanoTS:  types.NanoTS(1234567890000000000),
					updateNanoTS:    types.NanoTS(1334567890000000000),
				},
				expectedComments: testFirstComments9,
			},
			{
				name: "10_" + testFilename10,
				args: args{
					bboardID:        "test",
					articleID:       "test",
					ownerID:         "testOwner",
					commentsDBCS:    testFirstCommentsDBCS10,
					allCommentsDBCS: testComment10,
					lastTimeNanoTS:  types.NanoTS(1234567890000000000),
					updateNanoTS:    types.NanoTS(1334567890000000000),
				},
				expectedComments: testFirstComments10,
			},
			{
				name: "11_" + testFilename11,
				args: args{
					bboardID:        "test",
					articleID:       "test",
					ownerID:         "testOwner",
					commentsDBCS:    testFirstCommentsDBCS11,
					allCommentsDBCS: testComment11,
					lastTimeNanoTS:  types.NanoTS(1234567890000000000),
					updateNanoTS:    types.NanoTS(1334567890000000000),
				},
				expectedComments: testFirstComments11,
			},
			{
				name: "11_" + testFilename11 + "_the_rest",
				args: args{
					bboardID:        "test",
					articleID:       "test",
					ownerID:         "cheinshin",
					commentsDBCS:    testTheRestCommentsDBCS11,
					allCommentsDBCS: testComment11,
					lastTimeNanoTS:  types.NanoTS(1261396020004000000),
					updateNanoTS:    types.NanoTS(1334567890000000000),
				},
				expectedComments: testTheRestComments11,
			},
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logrus.Infof("%v: to ParseComments", tt.name)
			gotComments, _ := ParseComments(tt.args.bboardID, tt.args.articleID, tt.args.ownerID, tt.args.lastTimeNanoTS, tt.args.commentsDBCS, tt.args.allCommentsDBCS, tt.args.updateNanoTS, false)

			for _, each := range tt.expectedComments {
				each.UpdateNanoTS = tt.args.updateNanoTS
			}

			if len(gotComments) != len(tt.expectedComments) {
				t.Errorf("ParseComments: len: %v expected: %v", len(gotComments), len(tt.expectedComments))
			}

			testutil.TDeepEqual(t, "comments", gotComments, tt.expectedComments)
			for idx, each := range gotComments {
				if idx >= len(tt.expectedComments) {
					t.Errorf("comments: (%v/%v): %v", idx, len(gotComments), each.CreateTime.ToTime())
					continue
				}
				if each.CreateTime != tt.expectedComments[idx].CreateTime {
					t.Errorf("comments: (%v/%v): %v expected: %v", idx, len(gotComments), each.CreateTime.ToTime(), tt.expectedComments[idx].CreateTime.ToTime())
				}
			}
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
