package schema

import (
	"testing"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/go-pttbbsweb/types"
)

func TestUpdateUserDetail(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer User_c.Drop()

	type args struct {
		user *UserDetail
	}
	tests := []struct {
		name         string
		args         args
		updateNanoTS types.NanoTS
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			args:         args{user: testUserDetail0},
			updateNanoTS: 1234567890000000001,
		},
		{
			args:         args{user: testUserDetail0},
			updateNanoTS: 1234567890000000002,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.user.UpdateNanoTS = tt.updateNanoTS
			if err := UpdateUserDetail(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserDetail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewUserDetail(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		user_b       pttbbsapi.GetUserResult
		updateNanoTS types.NanoTS
	}
	tests := []struct {
		name         string
		args         args
		expectedUser *UserDetail
	}{
		// TODO: Add test cases.
		{
			args: args{
				user_b:       pttbbsapi.GetUserResult(testUser0),
				updateNanoTS: types.NanoTS(123456790000000000),
			},
			expectedUser: testUserDetail0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUser := NewUserDetail(tt.args.user_b, tt.args.updateNanoTS)

			gotUser.UpdateNanoTS = tt.expectedUser.UpdateNanoTS
			testutil.TDeepEqual(t, "user", gotUser, tt.expectedUser)
		})
	}
}
