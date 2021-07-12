package api

import (
	"reflect"
	"testing"
)

func TestIndex(t *testing.T) {
	// setupTest moves inside for-loop
	// teardownTest moves inside for-loop

	params := &IndexParams{}

	type args struct {
		params interface{}
	}
	tests := []struct {
		name     string
		args     args
		expected interface{}
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{&IndexParams{}},
			expected: &IndexResult{Data: params},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupTest()
			defer teardownTest()

			got, _, err := Index(testIP, tt.args.params, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("Index() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Index() = %v, want %v", got, tt.expected)
			}
		})
	}
}
