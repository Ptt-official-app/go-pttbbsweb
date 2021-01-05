package schema

import (
	"reflect"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func TestGetUserNewInfo(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer User_c.Drop()

	UpdateUserDetail(testUserDetail0)

	type args struct {
		userID bbs.UUserID
	}
	tests := []struct {
		name           string
		args           args
		expectedResult *UserNewInfo
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			args:           args{userID: "test_userid"},
			expectedResult: testUserNewInfo0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := GetUserNewInfo(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserNewInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.expectedResult) {
				t.Errorf("GetUserNewInfo() = %v, want %v", gotResult, tt.expectedResult)
			}
		})
	}
}
