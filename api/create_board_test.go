package api

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/gin-gonic/gin"
)

func TestCreateBoard(t *testing.T) {
	setupTest()
	defer teardownTest()

	path0 := &CreateBoardPath{
		ClsBid: 2,
	}

	params0 := &CreateBoardParams{
		Brdname:  "test3",
		BrdClass: "測試",
		BrdTitle: "測試1",
		BMs:      []bbs.UUserID{"okcool", "teemo"},
	}

	expected0 := &apitypes.BoardSummary{
		FBoardID:  "test3",
		Brdname:   "test3",
		Title:     "測試1",
		BoardType: "◎",
		Category:  "測試",
		BMs:       []bbs.UUserID{"okcool", "teemo"},
		StatAttr:  ptttype.NBRD_BOARD,
		Gid:       2,
		Idx:       "",
		Read:      true,
	}

	type args struct {
		remoteAddr string
		userID     bbs.UUserID
		params     interface{}
		path       interface{}
		c          *gin.Context
	}
	tests := []struct {
		name               string
		args               args
		expectedResult     interface{}
		expectedStatusCode int
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: params0, path: path0},
			expectedResult:     expected0,
			expectedStatusCode: 200,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotResult, gotStatusCode, err := CreateBoard(tt.args.remoteAddr, tt.args.userID, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotResult, tt.expectedResult)
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("CreateBoard() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}
