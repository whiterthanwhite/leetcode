package maxtrixcellsindistanceorder

import (
	"testing"
	"time"
)

func TestAllCellsDistOrder(t *testing.T) {
	type test struct {
		name                         string
		rows, cols, rCenter, cCenter int
		expected                     [][]int
	}
	tests := []test{
		{
			name:     "test 1",
			rows:     1,
			cols:     2,
			rCenter:  0,
			cCenter:  0,
			expected: [][]int{{0, 0}, {0, 1}},
		},
		{
			name:     "test 2",
			rows:     2,
			cols:     2,
			rCenter:  0,
			cCenter:  1,
			expected: [][]int{{0, 1}, {0, 0}, {1, 1}, {1, 0}},
		},
		{
			name:     "test 3",
			rows:     2,
			cols:     3,
			rCenter:  1,
			cCenter:  2,
			expected: [][]int{{1, 2}, {0, 2}, {1, 1}, {0, 1}, {1, 0}, {0, 0}},
		},
		{
			name:     "test 4",
			rows:     3,
			cols:     3,
			rCenter:  0,
			cCenter:  2,
			expected: [][]int{{0, 2}, {0, 1}, {1, 2}, {0, 0}, {1, 1}, {2, 2}, {1, 0}, {2, 1}, {2, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startTime := time.Now()
			actual := allCellsDistOrder(tt.rows, tt.cols, tt.rCenter, tt.cCenter)
			if len(actual) != len(tt.expected) {
				t.Error("arrays are different")
			} else {
				for i := range actual {
					for j := range actual[i] {
						if actual[i][j] != tt.expected[i][j] {
							t.Errorf("i: %d; j: %d; actual: %d; expected: %d", i, j, actual[i][j], tt.expected[i][j])
						}
					}
				}
			}
			t.Logf("Execution time: %v", time.Now().Sub(startTime))
		})
	}
}
