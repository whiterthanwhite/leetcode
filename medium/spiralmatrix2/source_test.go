package spiralmatrix2

import "testing"

func TestSpiralMatrix(t *testing.T) {
	type test struct {
		name     string
		input    int
		expected [][]int
	}
	tests := []test{
		{
			name:  "test 1",
			input: 3,
			expected: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
		{
			name:  "test 2",
			input: 1,
			expected: [][]int{
				{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := spiralMatrix(tt.input)
			if len(actual) != len(tt.expected) {
				t.Log("Actual and expected slices are different")
				t.FailNow()
			}
			for y := range actual {
				for x, v := range actual[y] {
					if v != tt.expected[y][x] {
						t.Errorf("At pos: x = %d, y = %d; Actual: %d; Expected: %d", x, y, v, tt.expected[y][x])
					}
				}
			}
		})
	}
}

func TestSpiralMatrix2(t *testing.T) {
	type test struct {
		name     string
		input    int
		expected [][]int
	}
	tests := []test{
		{
			name:  "test 1",
			input: 3,
			expected: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			name:  "test 2",
			input: 1,
			expected: [][]int{
				{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := spiralMatrix2(tt.input)
			if len(actual) != len(tt.expected) {
				t.Log("Actual and expected slices are different")
				t.FailNow()
			}
			for y := range actual {
				for x, v := range actual[y] {
					if v != tt.expected[y][x] {
						t.Errorf("At pos: x = %d, y = %d; Actual: %d; Expected: %d", x, y, v, tt.expected[y][x])
					}
				}
			}
		})
	}
}
