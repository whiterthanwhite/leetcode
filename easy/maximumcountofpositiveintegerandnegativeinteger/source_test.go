package maximumcountofpositiveintegerandnegativeinteger

import "testing"

func TestMaximumCount(t *testing.T) {
	type test struct {
		name     string
		input    []int
		expected int
	}
	tests := []test{
		{
			name:     "test 1",
			input:    []int{-2, -1, -1, 1, 2, 3},
			expected: 3,
		},
		{
			name:     "test 2",
			input:    []int{-3, -2, -1, 0, 0, 1, 2},
			expected: 3,
		},
		{
			name:     "test 3",
			input:    []int{5, 20, 66, 1314},
			expected: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := maximumCount(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %d; Expected: %d", actual, tt.expected)
			}
		})
	}
}
