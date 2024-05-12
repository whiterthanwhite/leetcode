package solvingquestionswithbrainpower

import "testing"

func TestMostPoints(t *testing.T) {
	type test struct {
		name     string
		input    [][]int
		expected int64
	}
	tests := []test{
		{
			name: "test 1",
			input: [][]int{
				{3, 2}, {4, 3}, {4, 4}, {2, 5},
			},
			expected: 5,
		},
		{
			name: "test 2",
			input: [][]int{
				{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5},
			},
			expected: 7,
		},
		{
			name: "test 3",
			input: [][]int{
				{21, 5}, {92, 3}, {74, 2}, {39, 4}, {58, 2}, {5, 5}, {49, 4}, {65, 3},
			},
			expected: 157,
		},
		{
			name: "test 4",
			input: [][]int{
				{21, 2}, {1, 2}, {12, 5}, {7, 2}, {35, 3}, {32, 2},
				{80, 2}, {91, 5}, {92, 3}, {27, 3}, {19, 1}, {37, 3},
				{85, 2}, {33, 4}, {25, 1}, {91, 4}, {44, 3}, {93, 3},
				{65, 4}, {82, 3}, {85, 5}, {81, 3}, {29, 2}, {25, 1},
				{74, 2}, {58, 1}, {85, 1}, {84, 2}, {27, 2}, {47, 5},
				{48, 4}, {3, 2}, {44, 3}, {60, 5}, {19, 2}, {9, 4},
				{29, 5}, {15, 3}, {1, 3}, {60, 2}, {63, 3}, {79, 3},
				{19, 1}, {7, 1}, {35, 1}, {55, 4}, {1, 4}, {41, 1},
				{58, 5},
			},
			expected: 781,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := mostPoints(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %d; Expected: %d", actual, tt.expected)
			}
		})
	}
}
