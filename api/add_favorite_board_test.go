package api

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/gin-gonic/gin"
)

func TestAddFavoriteBoard(t *testing.T) {
	setupTest()
	defer teardownTest()

	paramsLoad0 := &LoadGeneralBoardsParams{
		StartIdx: "vFSt-Q@WhoAmI",
	}

	_, _, _ = LoadGeneralBoardsByClass("localhost", "SYSOP", paramsLoad0, nil)

	params0 := &AddFavoriteBoardParams{
		FBoardID: "WhoAmI",
	}

	path0 := &AddFavoriteBoardPath{
		UserID: "SYSOP",
	}

	ret0 := AddFavoriteBoardResult(testBoardSummary10)

	type args struct {
		remoteAddr string
		userID     bbs.UUserID
		params     interface{}
		path       interface{}
		c          *gin.Context
	}
	tests := []struct {
		name           string
		args           args
		wantResult     interface{}
		wantStatusCode int
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			args:           args{params: params0, path: path0, userID: "SYSOP", remoteAddr: "localhost"},
			wantResult:     ret0,
			wantStatusCode: 200,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotResult, gotStatusCode, err := AddFavoriteBoard(tt.args.remoteAddr, tt.args.userID, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddFavoriteBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotResult, tt.wantResult)
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("AddFavoriteBoard() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}
		})
		wg.Wait()
	}
}
