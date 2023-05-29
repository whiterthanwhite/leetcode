package longestconsecutivesequence

import "testing"

func TestLongestConsecutive(t *testing.T) {
	type test struct {
		name     string
		nums     []int
		expected int
	}
	tests := []test{
		{
			name:     "test 1",
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4,
		},
		{
			name:     "test 2",
			nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			expected: 9,
		},
		{
			name:     "test 3",
			nums:     []int{1, 0, -1},
			expected: 3,
		},
		{
			name:     "test 4",
			nums:     []int{1, 2, 0, 1},
			expected: 3,
		},
		{
			name:     "test 5",
			nums:     []int{0},
			expected: 1,
		},
		{
			name:     "test 6",
			nums:     []int{},
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := longestConsecutive(tt.nums)
			if actual != tt.expected {
				t.Errorf("Actual: %d; Expected: %d", actual, tt.expected)
			}
		})
	}
}
