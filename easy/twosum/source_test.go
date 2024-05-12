package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	type test struct {
		name     string
		nums     []int
		target   int
		expected []int
	}
	tests := []test{
		{
			name:     "test 1",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "test 2",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "test 3",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := twoSum(tt.nums, tt.target)
			if actual == nil {
				t.Log("actual is nil")
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
