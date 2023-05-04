package counthillsandvalleysinanarray

import "testing"

func TestCountHillValley(t *testing.T) {
	type test struct {
		name     string
		input    []int
		expected int
	}
	tests := []test{
		{
			name:     "test 1",
			input:    []int{2, 4, 1, 1, 6, 5},
			expected: 3,
		},
		{
			name:     "test 2",
			input:    []int{6, 6, 5, 5, 4, 1},
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := countHillValley(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %d; Expected: %d", actual, tt.expected)
			}
		})
	}
}
