package matrixdiagonalsum

import "testing"

func TestDiagonalSum(t *testing.T) {
	type test struct {
		name     string
		input    [][]int
		expected int
	}
	tests := []test{
		{
			name: "test 1",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: 25,
		},
		{
			name: "test 2",
			input: [][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			expected: 8,
		},
		{
			name: "test 3",
			input: [][]int{
				{5},
			},
			expected: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := diagonalSum(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %d; Expected: %d", actual, tt.expected)
			}
		})
	}
}
