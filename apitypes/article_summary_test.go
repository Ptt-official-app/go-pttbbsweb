package apitypes

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func TestToFTitle(t *testing.T) {
	origServiceMode := types.SERVICE_MODE
	defer func() {
		types.SERVICE_MODE = origServiceMode
	}()

	type args struct {
		title string
	}
	tests := []struct {
		name        string
		args        args
		serviceMode types.ServiceMode
		expected    string
	}{
		// TODO: Add test cases.
		{
			args:        args{title: "某個標題"},
			serviceMode: types.DEV,
			expected:    "[測試]某個標題",
		},
		{
			args:        args{title: "某個標題"},
			serviceMode: types.PRODUCTION,
			expected:    "某個標題",
		},
		{
			args:        args{title: "某個標題"},
			serviceMode: types.INFO,
			expected:    "某個標題",
		},
		{
			args:        args{title: "某個標題"},
			serviceMode: types.DEBUG,
			expected:    "某個標題",
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			types.SERVICE_MODE = tt.serviceMode
			if got := ToFTitle(tt.args.title); got != tt.expected {
				t.Errorf("ToFTitle() = %v, want %v", got, tt.expected)
			}
		})
		wg.Wait()
	}
}
