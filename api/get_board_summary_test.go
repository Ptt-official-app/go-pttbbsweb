package api

import (
	"reflect"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/gin-gonic/gin"
)

func TestGetBoardSummary(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer schema.UserReadBoard_c.Drop()

	update0 := &schema.UserReadBoard{UserID: "SYSOP", BBoardID: "1_test1", UpdateNanoTS: types.Time8(1234567891).ToNanoTS()}

	_, _ = schema.UserReadBoard_c.Update(update0, update0)

	params0 := &GetBoardSummaryParams{}
	path0 := &GetBoardSummaryPath{FBoardID: "test1"}

	result0 := GetBoardSummaryResult(&apitypes.BoardSummary{
		FBoardID:  "test1",
		Brdname:   "test1",
		Title:     "測試1",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "測試",
		NUser:     100,
		BMs:       []bbs.UUserID{"okcool", "teemo"},
		Read:      true,
		Total:     123,

		LastPostTime: 1234567890,
		StatAttr:     ptttype.NBRD_BOARD,
	})

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
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotResult, gotStatusCode, err := GetBoardSummary(tt.args.remoteAddr, tt.args.userID, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBoardSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.expectedResult) {
				t.Errorf("GetBoardSummary() gotResult = %v, want %v", gotResult, tt.expectedResult)
			}
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("GetBoardSummary() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}
