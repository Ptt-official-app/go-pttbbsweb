package api

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetUserVisitCount(t *testing.T) {
	setupTest()
	defer teardownTest()
	type args struct {
		remoteAddr string
		params     interface{}
		c          *gin.Context
	}
	tests := []struct {
		name           string
		args           args
		wantResult     interface{}
		wantStatusCode int
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			name:           "test get user visit count is available",
			args:           args{testIP, nil, nil},
			wantResult:     &GetUserVisitCountResult{},
			wantStatusCode: 200,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotStatusCode, err := GetUserVisitCount(tt.args.remoteAddr, tt.args.params, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserVisitCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("GetUserVisitCount() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("GetUserVisitCount() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}
		})
	}
}
