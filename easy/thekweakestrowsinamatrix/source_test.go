package thekweakestrowsinamatrix

import "testing"

func TestKWeakestRows(t *testing.T) {
	type test struct {
		name     string
		mat      [][]int
		k        int
		expected []int
	}
	tests := []test{
		{
			name: "test 1",
			mat: [][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 1},
			},
			k:        3,
			expected: []int{2, 0, 3},
		},
		{
			name: "test 2",
			mat: [][]int{
				{1, 0, 0, 0},
				{1, 1, 1, 1},
				{1, 0, 0, 0},
				{1, 0, 0, 0},
			},
			k:        2,
			expected: []int{0, 2},
		},
		{
			name: "test 3",
			mat: [][]int{
				{1, 1, 0},
				{1, 0, 0},
				{1, 0, 0},
				{1, 1, 1},
				{1, 1, 0},
				{0, 0, 0},
			},
			k:        4,
			expected: []int{5, 1, 2, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := kWeakestRows(tt.mat, tt.k)
			if len(actual) != len(tt.expected) {
				t.Fatal("Slices are different")
			}
			for i, v := range actual {
				if v != tt.expected[i] {
					t.Errorf("Actual: %d; Expected: %d at pos %d", v, tt.expected[i], i)
				}
			}
		})
	}
}
