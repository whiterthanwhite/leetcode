package twosum2inputarrayissorted

import "testing"

func TestTwoSum(t *testing.T) {
	type test struct {
		name     string
		numbers  []int
		target   int
		expected []int
	}
	tests := []test{
		{
			name:     "test 1",
			numbers:  []int{2, 7, 11, 15},
			target:   9,
			expected: []int{1, 2},
		},
		{
			name:     "test 2",
			numbers:  []int{2, 3, 4},
			target:   6,
			expected: []int{1, 3},
		},
		{
			name:     "test 3",
			numbers:  []int{-1, 0},
			target:   -1,
			expected: []int{1, 2},
		},
		{
			name:     "test 4",
			numbers:  []int{1, 2, 3, 4, 4, 9, 56, 90},
			target:   8,
			expected: []int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := twoSum(tt.numbers, tt.target)
			if actual == nil {
				t.Log("Actual is nil")
				t.FailNow()
			}
			if len(actual) != len(tt.expected) {
				t.Log("Actual is defferent from expected")
				t.FailNow()
			}
			for i, v := range actual {
				if v != tt.expected[i] {
					t.Errorf("Actual: %d; Expected: %d; Position: %d", v, tt.expected[i], i)
				}
			}
		})
	}
}
