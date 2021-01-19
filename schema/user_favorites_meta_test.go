package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func TestUpdateUserFavoritesMeta(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer UserFavoritesMeta_c.Drop()

	type args struct {
		meta *UserFavoritesMeta
	}
	tests := []struct {
		name     string
		args     args
		expected *UserFavoritesMeta
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{testFavMeta0},
			expected: testFavMeta0,
		},
		{
			args:     args{testFavMeta1},
			expected: testFavMeta1,
		},
		{
			args:     args{testFavMeta0},
			expected: testFavMeta1,
			wantErr:  true,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if err := UpdateUserFavoritesMeta(tt.args.meta); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserFavoritesMeta() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, _ := GetUserFavoritesMeta(tt.args.meta.UserID)
			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
		wg.Wait()
	}
}
