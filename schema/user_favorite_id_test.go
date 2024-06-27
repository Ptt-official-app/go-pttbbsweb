package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	pttbbsfav "github.com/Ptt-official-app/go-pttbbs/ptt/fav"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

func TestGetUserFavoriteIDsByPttbids(t *testing.T) {
	setupTest()
	defer teardownTest()

	_ = UpdateUserFavorites("SYSOP", 0, testUserFavorites0, 1234567890000000000, 1234567890000000000)

	pttbids0 := []ptttype.Bid{1, 9, 10}
	userFavIDs0 := []*UserFavoriteID{
		{UserID: "SYSOP", MTime: 1234567890000000000, TheID: 1, TheType: pttbbsfav.FAVT_BOARD},
		{UserID: "SYSOP", MTime: 1234567890000000000, TheID: 9, TheType: pttbbsfav.FAVT_BOARD},
		{UserID: "SYSOP", MTime: 1234567890000000000, TheID: 1, TheType: pttbbsfav.FAVT_BOARD, LevelIdx: ":2"},
	}

	type args struct {
		userID          bbs.UUserID
		doubleBufferIdx int
		pttbids         []ptttype.Bid
		mTime           types.NanoTS
	}
	tests := []struct {
		name                string
		args                args
		wantUserFavoriteIDs []*UserFavoriteID
		wantErr             bool
	}{
		// TODO: Add test cases.
		{
			args:                args{userID: "SYSOP", doubleBufferIdx: 0, pttbids: pttbids0, mTime: 1234567890000000000},
			wantUserFavoriteIDs: userFavIDs0,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotUserFavoriteIDs, err := GetUserFavoriteIDsByPttbids(tt.args.userID, tt.args.doubleBufferIdx, tt.args.pttbids, tt.args.mTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserFavoriteIDsByPttbids() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotUserFavoriteIDs, tt.wantUserFavoriteIDs)
		})
		wg.Wait()
	}
}
