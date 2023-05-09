package spiralmatrix

import "testing"

func TestSpiralOrder(t *testing.T) {
	type test struct {
		name     string
		input    [][]int
		expected []int
	}
	tests := []test{
		{
			name: "test 1",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "test 2",
			input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			expected: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name: "test 3",
			input: [][]int{
				{2, 3},
			},
			expected: []int{2, 3},
		},
		{
			name: "test 4",
			input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			expected: []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := spiralOrder(tt.input)
			if len(actual) != len(tt.expected) {
				t.Logf("Length difference between slices! Actual: %d; Epxected: %d", len(actual), len(tt.expected))
				t.FailNow()
			}
			for i, v := range actual {
				if v != tt.expected[i] {
					t.Errorf("At pos: %d; Actual: %d; Expected: %d", i, v, tt.expected[i])
				}
			}
		})
	}
}
