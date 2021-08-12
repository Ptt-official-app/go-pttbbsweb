package api

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/gin-gonic/gin"
)

func TestLoadFavoriteBoards(t *testing.T) {
	setupTest()
	defer teardownTest()

	params0 := NewLoadFavoriteBoardsParams()
	path0 := &LoadFavoriteBoardsPath{
		UserID: "SYSOP",
	}

	result0 := &LoadFavoriteBoardsResult{
		List:    testFavoriteBoards0,
		NextIdx: "",
	}

	params1 := &LoadFavoriteBoardsParams{
		StartIdx:  "2",
		Max:       2,
		Ascending: true,
	}

	result1 := &LoadFavoriteBoardsResult{
		List:    testFavoriteBoards0[2:4],
		NextIdx: "4",
	}

	params2 := &LoadFavoriteBoardsParams{
		StartIdx:  "4",
		Max:       2,
		Ascending: false,
	}

	list2 := []*apitypes.BoardSummary{
		testFavoriteBoards0[4],
		testFavoriteBoards0[3],
	}
	result2 := &LoadFavoriteBoardsResult{
		List:    list2,
		NextIdx: "2",
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
			expectedResult:     result0,
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: params1, path: path0},
			expectedResult:     result1,
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: params2, path: path0},
			expectedResult:     result2,
			expectedStatusCode: 200,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotResult, gotStatusCode, err := LoadFavoriteBoards(tt.args.remoteAddr, tt.args.userID, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadFavoriteBoards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "result", gotResult, tt.expectedResult)
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("LoadFavoriteBoards() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}
