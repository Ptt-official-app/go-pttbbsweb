package dbcs

import (
	"reflect"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func TestParseFirstComments(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		bboardID             bbs.BBoardID
		articleID            bbs.ArticleID
		ownerID              bbs.UUserID
		articleCreateTime    types.NanoTS
		articleMTime         types.NanoTS
		commentsDBCS         []byte
		origFirstCommentsMD5 string
	}
	tests := []struct {
		name                     string
		args                     args
		updateNanoTS             types.NanoTS
		expectedFirstComments    []*schema.Comment
		expectedFirstCommentsMD5 string
		expectedTheRestComments  []byte
		wantErr                  bool
	}{
		// TODO: Add test cases.
		{
			name: "0_" + testFilename0,
			args: args{
				bboardID:          "test",
				articleID:         "test",
				ownerID:           "testOwner",
				articleCreateTime: types.NanoTS(1607202237000000000),
				articleMTime:      types.NanoTS(1607802720000000000),
				commentsDBCS:      testComment0,
			},
			updateNanoTS:             types.NanoTS(1607802730000000000),
			expectedFirstComments:    testFullFirstComments0,
			expectedFirstCommentsMD5: "lUNLzf4Qpeos8HBS676eWg",
		},
		{
			name: "0_" + testFilename0 + ":repeat",
			args: args{
				bboardID:          "test",
				articleID:         "test",
				ownerID:           "testOwner",
				articleCreateTime: types.NanoTS(1607202237000000000),
				articleMTime:      types.NanoTS(1607802720000000000),
				commentsDBCS:      testComment0,
			},
			updateNanoTS:             types.NanoTS(1607802740000000000),
			expectedFirstCommentsMD5: "lUNLzf4Qpeos8HBS676eWg",
		},
		{
			name: "0_" + testFilename0,
			args: args{
				bboardID:          "test",
				articleID:         "test01",
				ownerID:           "testOwner",
				articleCreateTime: types.NanoTS(1607202237000000000),
				articleMTime:      types.NanoTS(1607802690000000000),
				commentsDBCS:      testComment0,
			},
			updateNanoTS:             types.NanoTS(1607802730000000000),
			expectedFirstComments:    testFullFirstComments01,
			expectedFirstCommentsMD5: "lUNLzf4Qpeos8HBS676eWg",
		},
		{
			name: "0_" + testFilename0 + ":2021",
			args: args{
				bboardID:          "test",
				articleID:         "test02",
				ownerID:           "testOwner",
				articleCreateTime: types.NanoTS(1607202237000000000),
				articleMTime:      types.NanoTS(1639338720000000000),
				commentsDBCS:      testComment0,
			},
			updateNanoTS:             types.NanoTS(1607802740000000000),
			expectedFirstComments:    testFullFirstComments02,
			expectedFirstCommentsMD5: "lUNLzf4Qpeos8HBS676eWg",
		},
		{
			name: "1_" + testFilename1,
			args: args{
				bboardID:          "test",
				articleID:         "test1",
				ownerID:           "testOwner",
				articleCreateTime: types.NanoTS(1607250193000000000),
				articleMTime:      types.NanoTS(1607802780000000000),
				commentsDBCS:      testComment1,
			},
			updateNanoTS:             types.NanoTS(1607802790000000000),
			expectedFirstComments:    testFullFirstComments1,
			expectedFirstCommentsMD5: "WVuJu6yziL3Xw0LCXoIXVw",
		},
		{
			name: "2_" + testFilename2,
			args: args{
				bboardID:          "test",
				articleID:         "test2",
				ownerID:           "testOwner",
				articleCreateTime: types.NanoTS(1607983972000000000),
				articleMTime:      types.NanoTS(1607983972000000000),
				commentsDBCS:      testComment2,
			},
			updateNanoTS:          types.NanoTS(1607802790000000000),
			expectedFirstComments: testFullFirstComments2,
		},
		{
			name: "3_" + testFilename3,
			args: args{
				bboardID:          "test",
				articleID:         "test3",
				ownerID:           "testOwner",
				articleCreateTime: types.NanoTS(1608433078000000000),
				articleMTime:      types.NanoTS(1608433078000000000),
				commentsDBCS:      testComment3,
			},
			updateNanoTS:          types.NanoTS(1608433078000000000),
			expectedFirstComments: testFullFirstComments3,
		},
		{
			name: "4_" + testFilename4,
			args: args{
				bboardID:          "test",
				articleID:         "test4",
				ownerID:           "testOwner",
				articleCreateTime: types.NanoTS(1608388504000000000),
				articleMTime:      types.NanoTS(1608388624000000000),
				commentsDBCS:      testComment4,
			},
			updateNanoTS:             types.NanoTS(1608435524000000000),
			expectedFirstComments:    testFullFirstComments4,
			expectedFirstCommentsMD5: "3fjMk__1yvzpuEgq8jfdmg",
		},
		{
			name: "5_" + testFilename5,
			args: args{
				bboardID:          "test",
				articleID:         "test5",
				ownerID:           "testOwner",
				articleCreateTime: types.NanoTS(1048928980000000000),
				articleMTime:      types.NanoTS(1048928980000000000),
				commentsDBCS:      testComment5,
			},
			updateNanoTS:             types.NanoTS(1608435524000000000),
			expectedFirstComments:    testFullFirstComments5,
			expectedFirstCommentsMD5: "gQUbzCzxvt83giSwqT4odw",
		},
		{
			name: "6_" + testFilename6,
			args: args{
				bboardID:          "test",
				articleID:         "test6",
				ownerID:           "testOwner",
				articleCreateTime: types.NanoTS(1581529604000000000),
				articleMTime:      types.NanoTS(1581628710000000000),
				commentsDBCS:      testComment6,
			},
			updateNanoTS:             types.NanoTS(1608435524000000000),
			expectedFirstComments:    testFullFirstComments6,
			expectedFirstCommentsMD5: "rbq5acSoZFqWxpyORdI93Q",
			expectedTheRestComments:  testTheRestCommentsDBCS6,
		},
		{
			name: "7_" + testFilename7,
			args: args{
				bboardID:          "test",
				articleID:         "test7",
				ownerID:           "testOwner",
				articleCreateTime: types.NanoTS(1049092184000000000),
				articleMTime:      types.NanoTS(1049092244000000000),
				commentsDBCS:      testComment7,
			},
			updateNanoTS:             types.NanoTS(1608435524000000000),
			expectedFirstComments:    testFullFirstComments7,
			expectedFirstCommentsMD5: "ohOkfXXoey16NGhqsEvWBg",
		},
		{
			name: "8_" + testFilename8,
			args: args{
				bboardID:          "test",
				articleID:         "test8",
				ownerID:           "testOwner",
				articleCreateTime: types.NanoTS(1419202243000000000),
				articleMTime:      types.NanoTS(1419202834000000000),
				commentsDBCS:      testComment8,
			},
			updateNanoTS:             types.NanoTS(1608435524000000000),
			expectedFirstComments:    testFullFirstComments8,
			expectedFirstCommentsMD5: "jYSd2b4ScDh1JXFhy4yUHw",
		},
		{
			name: "9_" + testFilename9,
			args: args{
				bboardID:          "test",
				articleID:         "test9",
				ownerID:           "testOwner",
				articleCreateTime: types.NanoTS(1608360368000000000),
				articleMTime:      types.NanoTS(1608375447000000000),
				commentsDBCS:      testComment9,
			},
			updateNanoTS:             types.NanoTS(1608435524000000000),
			expectedFirstComments:    testFullFirstComments9,
			expectedFirstCommentsMD5: "suLbbfMSixdaM_j2dZj9Aw",
		},
		{
			name: "10_" + testFilename10,
			args: args{
				bboardID:          "test",
				articleID:         "test10",
				ownerID:           "testOwner",
				articleCreateTime: types.NanoTS(1607503629000000000),
				articleMTime:      types.NanoTS(1607624123000000000),
				commentsDBCS:      testComment10,
			},
			updateNanoTS:             types.NanoTS(1608435524000000000),
			expectedFirstComments:    testFullFirstComments10,
			expectedFirstCommentsMD5: "BLKFYC7nlC3i_GH8jBpf5w",
		},
		{
			name: "11_" + testFilename11,
			args: args{
				bboardID:          "test",
				articleID:         "test11",
				ownerID:           "testOwner",
				articleCreateTime: types.NanoTS(1608551120000000000),
				articleMTime:      types.NanoTS(1608567097000000000),
				commentsDBCS:      testComment11,
			},
			updateNanoTS:             types.NanoTS(1688435524000000000),
			expectedFirstComments:    testFullFirstComments11,
			expectedFirstCommentsMD5: "AYuuWvrS1A7fz6dmcQQwIg",
			expectedTheRestComments:  testTheRestCommentsDBCS11,
		},
		{
			name: "12_" + testFilename12,
			args: args{
				bboardID:          "test",
				articleID:         "test12",
				ownerID:           "testOwner",
				articleCreateTime: types.NanoTS(1608548064000000000),
				articleMTime:      types.NanoTS(1608566424000000000),
				commentsDBCS:      testComment12,
			},
			updateNanoTS:             types.NanoTS(1688435524000000000),
			expectedFirstComments:    testFullFirstComments12,
			expectedFirstCommentsMD5: "hmB8dp37sWvtvooMvgItOA",
			expectedTheRestComments:  testTheRestCommentsDBCS12,
		},
		{
			name: "13_" + testFilename13,
			args: args{
				bboardID:          "test",
				articleID:         "test13",
				ownerID:           "testOwner",
				articleCreateTime: types.NanoTS(1037679701000000000),
				articleMTime:      types.NanoTS(1037679702000000000),
				commentsDBCS:      testComment13,
			},
			updateNanoTS:             types.NanoTS(1688435524000000000),
			expectedFirstComments:    testFullFirstComments13,
			expectedFirstCommentsMD5: "GggVUAiHZqVwe5ck0uOw8w",
			expectedTheRestComments:  testTheRestCommentsDBCS13,
		},
		{
			name: "14_" + testFilename14,
			args: args{
				bboardID:          "test",
				articleID:         "test14",
				ownerID:           "testOwner",
				articleCreateTime: types.NanoTS(1616036734000000000),
				articleMTime:      types.NanoTS(1616119706000000000),
				commentsDBCS:      testComment14,
			},
			updateNanoTS:             types.NanoTS(1688435524000000000),
			expectedFirstComments:    testFullFirstComments14,
			expectedFirstCommentsMD5: "Do-tp_iSyqBDw2iISeo86g",
			expectedTheRestComments:  testTheRestCommentsDBCS14,
		},
		{
			name: "15_" + testFilename15,
			args: args{
				bboardID:          "test",
				articleID:         "test15",
				ownerID:           "testOwner",
				articleCreateTime: types.NanoTS(1466655288000000000),
				articleMTime:      types.NanoTS(1469169588000000000),
				commentsDBCS:      testComment15,
			},
			updateNanoTS:             types.NanoTS(1688435524000000000),
			expectedFirstComments:    testFullFirstComments15,
			expectedFirstCommentsMD5: "i-aQp775L_hMxPVqnaqJRw",
			expectedTheRestComments:  testTheRestCommentsDBCS15,
		},
		{
			name: "16_" + testFilename16,
			args: args{
				bboardID:          "test",
				articleID:         "test16",
				ownerID:           "testOwner",
				articleCreateTime: types.NanoTS(1583484543000000000),
				articleMTime:      types.NanoTS(1586091103000000000),
				commentsDBCS:      testComment16,
			},
			updateNanoTS:             types.NanoTS(1688435524000000000),
			expectedFirstComments:    testFullFirstComments16,
			expectedFirstCommentsMD5: "E812lmz5JJfJsEIFsFOYIw",
			expectedTheRestComments:  testTheRestCommentsDBCS16,
		},
		{
			name: "17_" + testFilename17,
			args: args{
				bboardID:          "test",
				articleID:         "test17",
				ownerID:           "testOwner",
				articleCreateTime: types.NanoTS(1618064338000000000),
				articleMTime:      types.NanoTS(1618064339000000000),
				commentsDBCS:      testComment17,
			},
			updateNanoTS:             types.NanoTS(1688435524000000000),
			expectedFirstComments:    testFullFirstComments17,
			expectedFirstCommentsMD5: "LpgHzpQMcdSiT6sOoX-F-g",
			expectedTheRestComments:  testTheRestCommentsDBCS17,
		},
		{
			name: "18_" + testFilename18,
			args: args{
				bboardID:          "test",
				articleID:         "test18",
				ownerID:           "testOwner",
				articleCreateTime: types.NanoTS(1584668622000000000),
				articleMTime:      types.NanoTS(1584787002000000000),
				commentsDBCS:      testComment18,
			},
			updateNanoTS:             types.NanoTS(1688435524000000000),
			expectedFirstComments:    testFullFirstComments18,
			expectedFirstCommentsMD5: "vMNTfs4ySxG0pglHqsGfzg",
			expectedTheRestComments:  testTheRestCommentsDBCS18,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotFirstComments, gotFirstCommentsMD5, gotTheRestComments, err := ParseFirstComments(tt.args.bboardID, tt.args.articleID, tt.args.ownerID, tt.args.articleCreateTime, tt.args.articleMTime, tt.args.commentsDBCS, tt.args.origFirstCommentsMD5)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFirstComments: err: %v expected: %v", err, tt.wantErr)
			}

			testutil.TDeepEqual(t, "firstComments", gotFirstComments, tt.expectedFirstComments)

			if gotFirstCommentsMD5 != tt.expectedFirstCommentsMD5 {
				t.Errorf("ParseFirstComments() gotFirstCommentsMD5 = %v, want %v", gotFirstCommentsMD5, tt.expectedFirstCommentsMD5)
			}

			if !reflect.DeepEqual(gotTheRestComments, tt.expectedTheRestComments) {
				t.Errorf("ParseFirstComments() gotTheRestComments = %v, want %v", gotTheRestComments, tt.expectedTheRestComments)
			}

			if len(gotFirstComments) == 0 {
				return
			}
			_ = schema.UpdateComments(gotFirstComments, tt.updateNanoTS)
		})
		wg.Wait()
	}
}

func Test_splitFirstComments(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		commentsDBCS []byte
	}
	tests := []struct {
		name                      string
		args                      args
		expectedFirstCommentsDBCS []byte
		expectedTheRestComments   []byte
	}{
		// TODO: Add test cases.
		{
			name:                      "0_" + testFilename0,
			args:                      args{commentsDBCS: testComment0},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS0,
			expectedTheRestComments:   testTheRestCommentsDBCS0,
		},
		{
			name:                      "1_" + testFilename1,
			args:                      args{commentsDBCS: testComment1},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS1,
			expectedTheRestComments:   testTheRestCommentsDBCS1,
		},
		{
			name:                      "2_" + testFilename2,
			args:                      args{commentsDBCS: testComment2},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS2,
			expectedTheRestComments:   testTheRestCommentsDBCS2,
		},
		{
			name:                      "3_" + testFilename3,
			args:                      args{commentsDBCS: testComment3},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS3,
			expectedTheRestComments:   testTheRestCommentsDBCS3,
		},
		{
			name:                      "4_" + testFilename4,
			args:                      args{commentsDBCS: testComment4},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS4,
			expectedTheRestComments:   testTheRestCommentsDBCS4,
		},
		{
			name:                      "5_" + testFilename5,
			args:                      args{commentsDBCS: testComment5},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS5,
			expectedTheRestComments:   testTheRestCommentsDBCS5,
		},
		{
			name:                      "6_" + testFilename6,
			args:                      args{commentsDBCS: testComment6},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS6,
			expectedTheRestComments:   testTheRestCommentsDBCS6,
		},
		{
			name:                      "7_" + testFilename7,
			args:                      args{commentsDBCS: testComment7},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS7,
			expectedTheRestComments:   testTheRestCommentsDBCS7,
		},
		{
			name:                      "8_" + testFilename8,
			args:                      args{commentsDBCS: testComment8},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS8,
			expectedTheRestComments:   testTheRestCommentsDBCS8,
		},
		{
			name:                      "9_" + testFilename9,
			args:                      args{commentsDBCS: testComment9},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS9,
			expectedTheRestComments:   testTheRestCommentsDBCS9,
		},
		{
			name:                      "10_" + testFilename10,
			args:                      args{commentsDBCS: testComment10},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS10,
			expectedTheRestComments:   testTheRestCommentsDBCS10,
		},
		{
			name:                      "11_" + testFilename11,
			args:                      args{commentsDBCS: testComment11},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS11,
			expectedTheRestComments:   testTheRestCommentsDBCS11,
		},
		{
			name:                      "12_" + testFilename12,
			args:                      args{commentsDBCS: testComment12},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS12,
			expectedTheRestComments:   testTheRestCommentsDBCS12,
		},
		{
			name:                      "13_" + testFilename13,
			args:                      args{commentsDBCS: testComment13},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS13,
			expectedTheRestComments:   testTheRestCommentsDBCS13,
		},
		{
			name:                      "14_" + testFilename14,
			args:                      args{commentsDBCS: testComment14},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS14,
			expectedTheRestComments:   testTheRestCommentsDBCS14,
		},
		{
			name:                      "15_" + testFilename15,
			args:                      args{commentsDBCS: testComment15},
			expectedFirstCommentsDBCS: testFirstCommentsDBCS15,
			expectedTheRestComments:   testTheRestCommentsDBCS15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirstCommentsDBCS, gotTheRestComments := splitFirstComments(tt.args.commentsDBCS)
			if !reflect.DeepEqual(gotFirstCommentsDBCS, tt.expectedFirstCommentsDBCS) {
				t.Errorf("splitFirstComments() gotFirstCommentsDBCS = %v, want %v", gotFirstCommentsDBCS, tt.expectedFirstCommentsDBCS)
			}
			if !reflect.DeepEqual(gotTheRestComments, tt.expectedTheRestComments) {
				t.Errorf("splitFirstComments() gotTheRestComments = %v, want %v", gotTheRestComments, tt.expectedTheRestComments)
			}
		})
	}
}
