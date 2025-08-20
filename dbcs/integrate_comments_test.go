package dbcs

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/sirupsen/logrus"
)

func TestIntegrateComments(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		boardID              bbs.BBoardID
		articleID            bbs.ArticleID
		firstCommentDBCS     []byte
		commentDBCS          []byte
		articleCreateTime    types.NanoTS
		articleMTime         types.NanoTS
		isForwardOnly        bool
		isLastAlignEndNanoTS bool
	}
	tests := []struct {
		name                     string
		args                     args
		expectedNewComments      []*schema.Comment
		expectedToDeleteComments []*schema.CommentMD5
		wantErr                  bool
	}{
		// TODO: Add test cases.
		{
			name: "6_" + testFilename6,
			args: args{
				boardID:              "test",
				articleID:            "test6",
				commentDBCS:          testComment6,
				articleCreateTime:    types.NanoTS(1581529604000000000),
				articleMTime:         types.NanoTS(1581628710000000000),
				isForwardOnly:        false,
				isLastAlignEndNanoTS: true,
			},
			expectedNewComments: testFullTheRestComments6,
		},
		{
			name: "11_" + testFilename11,
			args: args{
				boardID:              "test",
				articleID:            "test11",
				commentDBCS:          testComment11,
				articleCreateTime:    types.NanoTS(1608551120000000000),
				articleMTime:         types.NanoTS(1608567097000000000),
				isForwardOnly:        false,
				isLastAlignEndNanoTS: true,
			},
			expectedNewComments: testFullTheRestComments11,
		},
		{
			name: "14_" + testFilename14,
			args: args{
				boardID:              "test",
				articleID:            "test14",
				commentDBCS:          testComment14,
				articleCreateTime:    types.NanoTS(1616036734000000000),
				articleMTime:         types.NanoTS(1616123306000000000),
				isForwardOnly:        false,
				isLastAlignEndNanoTS: true,
			},
			expectedNewComments: testFullTheRestComments14,
		},
		{
			name: "18_" + testFilename18,
			args: args{
				boardID:              "test",
				articleID:            "test18",
				commentDBCS:          testComment18,
				articleCreateTime:    types.NanoTS(1584668622000000000),
				articleMTime:         types.NanoTS(1584787002000000000),
				isForwardOnly:        false,
				isLastAlignEndNanoTS: true,
			},
			expectedNewComments: testFullTheRestComments18,
		},
		{
			name: "19_" + testFilename19,
			args: args{
				boardID:              "test",
				articleID:            "test19",
				commentDBCS:          testComment19,
				articleCreateTime:    types.NanoTS(1516958302000000000),
				articleMTime:         types.NanoTS(1517273782000000000),
				isForwardOnly:        false,
				isLastAlignEndNanoTS: true,
			},
			expectedNewComments: testFullTheRestComments19,
		},
		{
			name: "20_" + testFilename20,
			args: args{
				boardID:              "test",
				articleID:            "test20",
				commentDBCS:          testComment20,
				articleCreateTime:    types.NanoTS(1620276637000000000),
				articleMTime:         types.NanoTS(1620276637000000000),
				isForwardOnly:        false,
				isLastAlignEndNanoTS: true,
			},
			expectedNewComments: testFullTheRestComments20,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			firstComments, _, _, _ := ParseFirstComments(tt.args.boardID, tt.args.articleID, "testOwner", tt.args.articleCreateTime, tt.args.articleMTime, tt.args.commentDBCS, "")

			logrus.Infof("IntegrateComments: firstComments: %v", len(firstComments))

			schema.UpdateComments(firstComments, 1688888888000000000)

			comments := ParseComments("testOwner", tt.args.commentDBCS, tt.args.commentDBCS)

			logrus.Infof("IntegrateComments: comments: %v", len(comments))

			gotNewComments, gotToDeleteComments, err := IntegrateComments(tt.args.boardID, tt.args.articleID, comments, tt.args.articleCreateTime, tt.args.articleMTime, tt.args.isForwardOnly, tt.args.isLastAlignEndNanoTS)
			if (err != nil) != tt.wantErr {
				t.Errorf("IntegrateComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "new", gotNewComments, tt.expectedNewComments)
			testutil.TDeepEqual(t, "toDelete", gotToDeleteComments, tt.expectedToDeleteComments)
		})
		wg.Wait()
	}
}
