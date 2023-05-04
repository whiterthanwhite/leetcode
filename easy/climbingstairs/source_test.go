package climbingstairs

import "testing"

func TestClimbStairs(t *testing.T) {
	type test struct {
		name            string
		input, expected int
	}
	tests := []test{
		{
			name:     "test 1",
			input:    2,
			expected: 2,
		},
		{
			name:     "test 2",
			input:    3,
			expected: 3,
		},
		{
			name:     "test 3",
			input:    4,
			expected: 5,
		},
		{
			name:     "test 4",
			input:    5,
			expected: 8,
		},
		{
			name:     "test 5",
			input:    6,
			expected: 13,
		},
		{
			name:     "test 6",
			input:    7,
			expected: 21,
		},
		{
			name:     "test 7",
			input:    8,
			expected: 34,
		},
		{
			name:     "test 8",
			input:    9,
			expected: 55,
		},
		{
			name:     "test 9",
			input:    45,
			expected: 1836311903,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := climbStairs(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %d; Expected: %d\n", actual, tt.expected)
			}
		})
	}
}
