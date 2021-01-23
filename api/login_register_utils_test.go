package api

import "testing"

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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gen2FAToken()
			if len(got) != tt.expected {
				t.Errorf("gen2FAToken() = %v, want %v", got, tt.expected)
			}
		})
	}
}
