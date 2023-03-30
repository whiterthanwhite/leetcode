package uglynumber

import (
	"testing"
)

func TestIsUgly(t *testing.T) {
	type test struct {
		name     string
		n        int
		expected bool
	}
	tests := []test{
		{
			name:     "test 1",
			n:        6,
			expected: true,
		},
		{
			name:     "test 2",
			n:        1,
			expected: true,
		},
		{
			name:     "test 3",
			n:        14,
			expected: false,
		},
		{
			name:     "test 4",
			n:        8,
			expected: true,
		},
		{
			name:     "test 5",
			n:        2,
			expected: true,
		},
		{
			name:     "test 6",
			n:        9,
			expected: true,
		},
		{
			name:     "test 7",
			n:        29,
			expected: false,
		},
		{
			name:     "test 8",
			n:        27,
			expected: true,
		},
		{
			name:     "test 9",
			n:        128,
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := isUgly(tt.n)
			if actual != tt.expected {
				t.Errorf("Actual: %v; Expected: %v", actual, tt.expected)
			}
		})
	}
}
