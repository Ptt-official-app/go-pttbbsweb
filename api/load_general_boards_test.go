package api

import (
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/gin-gonic/gin"
)

func TestLoadGeneralBoards(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer schema.UserReadBoard_c.Drop()

	update0 := &schema.UserReadBoard{UserID: "SYSOP", BBoardID: "1_test1", UpdateNanoTS: types.Time8(1234567891).ToNanoTS()}
	update1 := &schema.UserReadBoard{UserID: "SYSOP", BBoardID: "2_test2", UpdateNanoTS: types.Time8(1234567891).ToNanoTS()}

	_, _ = schema.UserReadBoard_c.Update(update0, update0)
	_, _ = schema.UserReadBoard_c.Update(update1, update1)

	params := &LoadGeneralBoardsParams{}
	expectedResult := &LoadGeneralBoardsResult{
		List: []*types.BoardSummary{
			{
				BBoardID:  "1_test1",
				Brdname:   "test1",
				Title:     "測試1",
				BrdAttr:   0,
				BoardType: "◎",
				Category:  "測試",
				NUser:     100,
				BMs:       []bbs.UUserID{"okcool", "teemo"},
				Read:      true,
				Total:     123,

				LastPostTimeTS_d: 1234567890,
				StatAttr_d:       ptttype.NBRD_BOARD,
			},
			{
				BBoardID:  "2_test2",
				Brdname:   "test2",
				Title:     "測試2",
				BrdAttr:   0,
				BoardType: "◎",
				Category:  "測試",
				NUser:     101,
				BMs:       []bbs.UUserID{"okcool2", "teemo2"},
				Read:      false,
				Total:     124,

				LastPostTimeTS_d: 1300000000,
				StatAttr_d:       ptttype.NBRD_BOARD,
			},
		},
		NextIdx: "textNextIdx",
	}
	type args struct {
		remoteAddr string
		userID     bbs.UUserID
		params     interface{}
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
			args:               args{remoteAddr: "localhost", userID: "SYSOP", params: params, c: nil},
			expectedResult:     expectedResult,
			expectedStatusCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotStatusCode, err := LoadGeneralBoards(tt.args.remoteAddr, tt.args.userID, tt.args.params, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadGeneralBoards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotResultVal := gotResult.(*LoadGeneralBoardsResult)
			expectedResultVal := tt.expectedResult.(*LoadGeneralBoardsResult)
			for idx, each := range gotResultVal.List {
				if idx >= len(expectedResultVal.List) {
					t.Errorf("LoadGeneralBoards() (%v/%v): %v", idx, len(gotResultVal.List), each)

				}

				utils.TDeepEqual(t, each, expectedResultVal.List[idx])
			}
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("LoadGeneralBoards() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
	}
}
