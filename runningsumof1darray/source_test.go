package runningsumof1darray

import (
	"testing"
)

func TestRunningSum(t *testing.T) {
	type test struct {
		name     string
		input    []int
		expected []int
	}
	tests := []test{
		{
			name:     "test 1",
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 3, 6, 10},
		},
		{
			name:     "test 2",
			input:    []int{1, 1, 1, 1, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "test 3",
			input:    []int{3, 1, 2, 10, 1},
			expected: []int{3, 4, 6, 16, 17},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := runningSum(tt.input)
			if actual == nil {
				t.Fail()
			}
			for i, v := range actual {
				if v != tt.expected[i] {
					t.Errorf("Actual: %d; Expected: %d; at position: %d", v, tt.expected[i], i)
				}
			}
		})
	}
}
