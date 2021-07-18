package api

import (
	"sync"
	"testing"
)

func Test_gen2FAToken(t *testing.T) {
	tests := []struct {
		name     string
		expected int
	}{
		// TODO: Add test cases.
		{
			expected: 6,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			got := gen2FAToken()
			if len(got) != tt.expected {
				t.Errorf("gen2FAToken() = %v, want %v", got, tt.expected)
			}
		})
	}
	wg.Wait()
}
