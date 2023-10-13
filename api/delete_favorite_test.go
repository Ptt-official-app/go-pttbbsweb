package api

import (
	"reflect"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func TestDeleteFavorite(t *testing.T) {
	setupTest()
	defer teardownTest()

	paramsLoad0 := &LoadGeneralBoardsParams{}

	_, _, _ = LoadGeneralBoardsByClass("localhost", "SYSOP", paramsLoad0, nil)

	params0 := &DeleteFavoriteParams{
		Idx: "0",
	}

	path0 := &DeleteFavoritePath{
		UserID: "SYSOP",
	}

	ret0 := &DeleteFavoriteResult{
		Success:   true,
		TokenUser: "SYSOP",
	}

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
			args:           args{remoteAddr: "localhost", userID: "SYSOP", params: params0, path: path0},
			wantResult:     ret0,
			wantStatusCode: 200,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotResult, gotStatusCode, err := DeleteFavorite(tt.args.remoteAddr, tt.args.userID, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteFavorite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("DeleteFavorite() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("DeleteFavorite() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}
		})
		wg.Wait()
	}
}
