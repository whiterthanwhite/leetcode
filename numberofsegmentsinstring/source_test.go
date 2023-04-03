package numberofsegmentsinstring

import "testing"

func TestCountSegments(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected int
	}
	tests := []test{
		{
			name:     "test 1",
			input:    "Hello, my name is John",
			expected: 5,
		},
		{
			name:     "test 2",
			input:    "Hello",
			expected: 1,
		},
		{
			name:     "test 3",
			input:    "",
			expected: 0,
		},
		{
			name:     "test 4",
			input:    "                ",
			expected: 0,
		},
		{
			name:     "test 5",
			input:    "Of all the gin joints in all the towns in all the world,   ",
			expected: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := countSegments(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %d; Expected: %d", actual, tt.expected)
			}
		})
	}
}
