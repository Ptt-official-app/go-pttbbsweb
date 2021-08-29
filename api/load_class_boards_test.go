package api

import (
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/mockhttp"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/gin-gonic/gin"
)

func TestLoadClassBoards(t *testing.T) {
	setupTest()
	defer teardownTest()

	ret := mockhttp.LoadGeneralBoards(nil)
	updateNanoTS := types.NowNanoTS()

	boardSummaries0 := make([]*schema.BoardSummary, len(ret.Boards))
	for idx, each_b := range ret.Boards {
		boardSummaries0[idx] = schema.NewBoardSummary(each_b, updateNanoTS)
	}

	update0 := &schema.UserReadBoard{UserID: "SYSOP", BBoardID: "1_test1", UpdateNanoTS: types.Time8(1234567891).ToNanoTS()}
	update1 := &schema.UserReadBoard{UserID: "SYSOP", BBoardID: "2_test2", UpdateNanoTS: types.Time8(1234567891).ToNanoTS()}

	_, _ = schema.UserReadBoard_c.Update(update0, update0)
	_, _ = schema.UserReadBoard_c.Update(update1, update1)

	_ = schema.UpdateBoardSummaries(boardSummaries0, updateNanoTS)

	classSummary0 := &schema.BoardSummary{
		BBoardID:  "3_3..........",
		Brdname:   "3..........",
		Title:     "市民廣場     報告  站長  ㄜ！",
		BoardType: "Σ",
		Category:  "....",

		Gid:        1,
		Bid:        3,
		IdxByName:  "3..........",
		IdxByClass: "....@..........",
	}

	_ = schema.UpdateBoardSummaries([]*schema.BoardSummary{classSummary0}, updateNanoTS)

	params0 := &LoadClassBoardsParams{
		SortBy:    ptttype.BSORT_BY_NAME,
		Ascending: true,
		Max:       10,
	}

	path0 := &LoadClassBoardsPath{
		ClsID: 3,
	}

	expected0 := &LoadClassBoardsResult{
		List: []*apitypes.BoardSummary{
			{
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

				Gid: 3,
				Bid: 1,

				Idx: "test1",
			},
			{
				FBoardID:  "test2",
				Brdname:   "test2",
				Title:     "測試2",
				BrdAttr:   0,
				BoardType: "◎",
				Category:  "測試",
				NUser:     101,
				BMs:       []bbs.UUserID{"okcool2", "teemo2"},
				Read:      false,
				Total:     124,

				LastPostTime: 1300000000,
				StatAttr:     ptttype.NBRD_BOARD,

				Gid: 3,
				Bid: 2,
				Idx: "test2",
			},
		},
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotStatusCode, err := LoadClassBoards(tt.args.remoteAddr, tt.args.userID, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadClassBoards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotResult, tt.expectedResult)
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("LoadClassBoards() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
	}
}

func Test_loadClassBoards(t *testing.T) {
	type args struct {
		userID   bbs.UUserID
		clsID    ptttype.Bid
		startIdx string
		asc      bool
		max      int
		sortBy   ptttype.BSortBy
		c        *gin.Context
	}
	tests := []struct {
		name                      string
		args                      args
		expectedBoardSummaries_db []*schema.BoardSummary
		expectedUserBoardInfoMap  map[bbs.BBoardID]*apitypes.UserBoardInfo
		expectedNextIdx           string
		wantErr                   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBoardSummaries_db, gotUserBoardInfoMap, gotNextIdx, err := loadClassBoards(tt.args.userID, tt.args.clsID, tt.args.startIdx, tt.args.asc, tt.args.max, tt.args.sortBy, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadClassBoards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "boardSummaries", gotBoardSummaries_db, tt.expectedBoardSummaries_db)
			testutil.TDeepEqual(t, "userBoardInfo", gotUserBoardInfoMap, tt.expectedUserBoardInfoMap)
			if gotNextIdx != tt.expectedNextIdx {
				t.Errorf("loadClassBoards() gotNextIdx = %v, want %v", gotNextIdx, tt.expectedNextIdx)
			}
		})
	}
}
