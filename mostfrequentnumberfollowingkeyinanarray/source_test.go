package mostfrequentnumberfollowingkeyinanarray

import "testing"

func TestMostFrequent(t *testing.T) {
	type test struct {
		name     string
		nums     []int
		key      int
		expected int
	}
	tests := []test{
		{
			name:     "test 1",
			nums:     []int{1, 100, 200, 1, 100},
			key:      1,
			expected: 100,
		},
		{
			name:     "test 2",
			nums:     []int{2, 2, 2, 2, 3},
			key:      2,
			expected: 2,
		},
		{
			name:     "test 3",
			nums:     []int{1, 1},
			key:      1,
			expected: 1,
		},
		{
			name:     "test 4",
			nums:     []int{1, 1000, 2},
			key:      1000,
			expected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := mostFrequent(tt.nums, tt.key)
			if actual != tt.expected {
				t.Errorf("Actual: %d; Expected: %d", actual, tt.expected)
			}
		})
	}
}
