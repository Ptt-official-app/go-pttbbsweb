package schema

import (
	"reflect"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func TestGetUserInfoSummary(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer User_c.Drop()

	testUserDetail0.UpdateNanoTS = testUserInfoSummary0.UpdateNanoTS
	UpdateUserDetail(testUserDetail0)

	type args struct {
		userID bbs.UUserID
	}
	tests := []struct {
		name           string
		args           args
		expectedResult *UserInfoSummary
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			args:           args{userID: "test_userid"},
			expectedResult: testUserInfoSummary0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := GetUserInfoSummary(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserInfoSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.expectedResult) {
				t.Errorf("GetUserInfoSummary() = %v, want %v", gotResult, tt.expectedResult)
			}
		})
	}
}
