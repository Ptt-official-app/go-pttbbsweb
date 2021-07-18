package api

import (
	"sync"
	"testing"
)

func Test_createCSRFToken(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			got, err := createCSRFToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("createCSRFToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !isValidCSRFToken(got) {
				t.Errorf("createCSRFToken: not valid csrf-token: got: %v", got)
			}
		})
	}
	wg.Wait()
}

func Test_isValidCSRFToken(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name     string
		args     args
		expected bool
	}{
		// TODO: Add test cases.
		{
			args:     args{raw: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGkiOiIiLCJleHAiOjEwNTk3MjIyMzQwLCJzdWIiOiIifQ.4Wr9FuxAve-kiHgt51-u4ewtI5CIkRk9tQQXbE5C8HU"},
			expected: true,
		},
		{
			args: args{raw: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGkiOiIiLCJleHAiOjk0OTM0MjI4MTIsInN1YiI6IlNZU09QMiJ9.VbixNBxg4h5FCyTmvhtVzBJ4HsE5_va-MPZzR8TLaY8"},
		},
		{
			args: args{raw: "invalid-raw"},
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if got := isValidCSRFToken(tt.args.raw); got != tt.expected {
				t.Errorf("isValidCSRFToken() = %v, want %v", got, tt.expected)
			}
		})
	}
	wg.Wait()
}
