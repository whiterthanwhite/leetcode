package sortthestudentsbytheirkthscore

import (
	"testing"
	"time"
)

func TestSortTheStudents(t *testing.T) {
	type test struct {
		name     string
		input    [][]int
		k        int
		expected [][]int
	}
	tests := []test{
		{
			name:     "test 1",
			input:    [][]int{{10, 6, 9, 1}, {7, 5, 11, 2}, {4, 8, 3, 15}},
			k:        2,
			expected: [][]int{{7, 5, 11, 2}, {10, 6, 9, 1}, {4, 8, 3, 15}},
		},
		{
			name:     "test 2",
			input:    [][]int{{3, 4}, {5, 6}},
			k:        0,
			expected: [][]int{{5, 6}, {3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startTime := time.Now()
			actual := sortTheStudents(tt.input, tt.k)
			if len(actual) != len(tt.expected) {
				t.Errorf("Slices are different")
			} else {
				for i := range actual {
					for j := range actual[i] {
						if actual[i][j] != tt.expected[i][j] {
							t.Errorf("At position: %d %d expected %d but actual %d",
								i, j, actual[i][j], tt.expected[i][j])
						}
					}
				}
			}
			t.Log("Execution time:", time.Now().Sub(startTime))
		})
	}
}
