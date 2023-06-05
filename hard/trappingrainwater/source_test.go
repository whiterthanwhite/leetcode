package trappingrainwater

import "testing"

func TestTrap(t *testing.T) {
	type test struct {
		name     string
		height   []int
		expected int
	}
	tests := []test{
		{
			name:   "test 1",
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			/*
				0 0 0 0 0 0 0 1 0 0 0 0
				0 0 0 1 0 0 0 1 1 0 1 0
				0 1 0 1 1 0 1 1 1 1 1 1
			*/
			expected: 6,
		},
		{
			name:   "test 2",
			height: []int{4, 2, 0, 3, 2, 5},
			/*
				0 0 0 0 0 1
				1 0 0 0 0 1
				1 0 0 1 0 1
				1 1 0 1 1 1
				1 1 0 1 1 1
			*/
			expected: 9,
		},
		{
			name:   "test 3",
			height: []int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6},
			/*
				0 0 0 1 0 0 0 0 1 0
				0 0 0 1 0 0 0 0 1 1
				1 1 0 1 0 0 1 0 1 1
				1 1 0 1 0 0 1 0 1 1
				1 1 0 1 0 0 1 0 1 1
				1 1 0 1 0 0 1 1 1 1
				1 1 1 1 1 1 1 1 1 1
			*/
			expected: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := trap(tt.height)
			if actual != tt.expected {
				t.Errorf("Actual: %d; Expected: %d", actual, tt.expected)
			}
		})
	}
}
