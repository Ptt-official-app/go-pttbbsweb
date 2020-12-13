package mock_http

import (
	"reflect"

	"github.com/Ptt-official-app/go-openbbsmiddleware/backend"
	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/types"

	"testing"
)

func Test_parseResult(t *testing.T) {
	setupTest()
	defer teardownTest()

	result1_b := &api.LoginResult{Jwt: "test_token", TokenType: "bearer"}
	var result1 *backend.LoginResult

	result2_b := &api.RegisterResult{Jwt: "test_token", TokenType: "bearer"}
	var result2 *backend.RegisterResult

	type args struct {
		backendResult interface{}
		httpResult    interface{}
		expected      interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			args: args{backendResult: result1_b, httpResult: &result1, expected: &result1},
		},
		{
			args: args{backendResult: result2_b, httpResult: &result2, expected: &result2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseResult(tt.args.backendResult, tt.args.httpResult)

			httpValue := reflect.ValueOf(tt.args.httpResult).Elem().Interface()
			expectedValue := reflect.ValueOf(tt.args.expected).Elem().Interface()
			types.TDeepEqual(t, httpValue, expectedValue)
		})
	}
}
