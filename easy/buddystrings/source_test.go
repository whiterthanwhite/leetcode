package buddystrings

import (
	"testing"
)

func TestBuddyStrings(t *testing.T) {
	type test struct {
		name     string
		input    string
		goal     string
		expected bool
	}
	tests := []test{
		{
			name:     "test 1",
			input:    "ab",
			goal:     "ba",
			expected: true,
		},
		{
			name:     "test 2",
			input:    "ab",
			goal:     "ab",
			expected: false,
		},
		{
			name:     "test 3",
			input:    "aa",
			goal:     "aa",
			expected: true,
		},
		{
			name:     "test 4",
			input:    "abab",
			goal:     "abab",
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := buddyStrings(tt.input, tt.goal)
			if actual != tt.expected {
				t.Errorf("Actual: %v; Expected: %v", actual, tt.expected)
			}
		})
	}
}
