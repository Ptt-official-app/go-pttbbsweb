package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
	"go.mongodb.org/mongo-driver/bson"
)

func TestUpdateAccessToken(t *testing.T) {
	setupTest()
	defer teardownTest()

	accessToken0 := NewAccessToken(bbs.UUserID("SYSOP"), "test_jwt", types.NanoTS(1234567890000000000))
	expected0 := &AccessToken{
		AccessToken: "test_jwt",
		UserID:      bbs.UUserID("SYSOP"),
	}

	type args struct {
		accessToken *AccessToken
	}
	tests := []struct {
		name     string
		args     args
		expected *AccessToken
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{accessToken: accessToken0},
			expected: expected0,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if err := UpdateAccessToken(tt.args.accessToken); (err != nil) != tt.wantErr {
				t.Errorf("UpdateAccessToken() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := &AccessToken{}
			query := bson.M{ACCESS_TOKEN_USER_ID_b: bbs.UUserID("SYSOP")}
			_ = AccessToken_c.FindOne(query, got, nil)

			tt.expected.UpdateNanoTS = got.UpdateNanoTS
			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
		wg.Wait()
	}
}
