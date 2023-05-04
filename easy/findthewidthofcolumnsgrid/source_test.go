package findthewidthofcolumnsgrid

import "testing"

func TestFindColumnWidth(t *testing.T) {
	type test struct {
		name     string
		input    [][]int
		expected []int
	}
	tests := []test{
		{
			name: "test 1",
			input: [][]int{
				{1},
				{22},
				{333},
			},
			expected: []int{3},
		},
		{
			name: "test 2",
			input: [][]int{
				{-15, 1, 3},
				{15, 7, 12},
				{5, 6, -2},
			},
			expected: []int{3, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := findColumnWidth(tt.input)
			if len(actual) != len(tt.expected) {
				t.Fatal("Slices are different")
			}
			for i, v := range actual {
				if v != tt.expected[i] {
					t.Errorf("Index: %d; Actual: %d; Expected: %d", i, v, tt.expected[i])
				}
			}
		})
	}
}
