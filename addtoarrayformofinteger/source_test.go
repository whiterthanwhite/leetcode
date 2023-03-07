package addtoarrayformofinteger

import (
	"testing"
	"time"
)

func TestAddToArrayForm(t *testing.T) {
	type test struct {
		name            string
		input, expected []int
		k               int
	}
	tests := []test{
		{
			name:     "test 1",
			input:    []int{1, 2, 0, 0},
			expected: []int{1, 2, 3, 4},
			k:        34,
		},
		{
			name:     "test 2",
			input:    []int{2, 7, 4},
			expected: []int{4, 5, 5},
			k:        181,
		},
		{
			name:     "test 3",
			input:    []int{2, 1, 5},
			expected: []int{1, 0, 2, 1},
			k:        806,
		},
		{
			name:     "test 4",
			input:    []int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 0, 6, 3},
			expected: []int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 5, 7, 9},
			k:        516,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startTime := time.Now()
			actual := addToArrayForm(tt.input, tt.k)
			if len(actual) != len(tt.expected) {
				t.Errorf("Actual: %v, but expected %v\n", actual, tt.expected)
			} else {
				for i, v := range actual {
					if v != tt.expected[i] {
						t.Errorf("Actual: %v, but expected %v at position %d\n", v, tt.expected[i], i)
					}
				}
			}
			t.Logf("Execution time: %v", time.Now().Sub(startTime))
		})
	}
}
