package api

import (
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/gin-gonic/gin"
)

func TestGetUserInfo(t *testing.T) {
	setupTest()
	defer teardownTest()

	path0 := &GetUserInfoPath{UserID: "SYSOP"}

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
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: nil, path: path0, c: nil},
			expectedResult:     testUserInfoResult0,
			expectedStatusCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotStatusCode, err := GetUserInfo(tt.args.remoteAddr, tt.args.userID, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			result := gotResult.(*GetUserInfoResult)
			result.UpdateTS = 0
			result.EmailTS = 0
			testutil.TDeepEqual(t, "result", result, tt.expectedResult)

			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("GetUserInfo() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
	}
}
