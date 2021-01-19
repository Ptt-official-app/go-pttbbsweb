package schema

import (
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func TestGetUserFavoritesMetaSummary(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer UserFavoritesMeta_c.Drop()

	UpdateUserFavoritesMeta(testFavMeta0)

	metaSummary0 := &UserFavoritesMetaSummary{
		UserID:          "SYSOP",
		DoubleBufferIdx: 0,
		MTime:           1234567890000000000,
	}

	type args struct {
		userID bbs.UUserID
	}
	tests := []struct {
		name           string
		args           args
		expectedResult *UserFavoritesMetaSummary
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			args:           args{userID: "SYSOP"},
			expectedResult: metaSummary0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := GetUserFavoritesMetaSummary(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserFavoritesMetaSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "result", gotResult, tt.expectedResult)
		})
	}
}
