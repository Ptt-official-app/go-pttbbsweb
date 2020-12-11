package main

import (
	"net/http"
	"testing"
)

func Test_setRequest(t *testing.T) {
	params := "test"
	headers := make(map[string]string)
	headers["Contetn-Type"] = "application/json"

	type args struct {
		path    string
		params  interface{}
		jwt     string
		headers map[string]string
	}
	tests := []struct {
		name     string
		args     args
		expected *http.Request
	}{
		// TODO: Add test cases.
		{
			args: args{params: params, headers: nil},
		},
		{
			args: args{params: params, headers: headers},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := setRequest(tt.args.path, tt.args.params, tt.args.jwt, tt.args.headers)
			if got == nil {
				t.Errorf("unable to setRequest")
			}
		})
	}
}
