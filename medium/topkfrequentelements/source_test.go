package topkfrequentelements

import "testing"

func TestTopKFrequent(t *testing.T) {
	type test struct {
		name     string
		nums     []int
		k        int
		expected []int
	}
	tests := []test{
		{
			name:     "test 1",
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			name:     "test 2",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "test 3",
			nums:     []int{4, 1, -1, 2, -1, 2, 3},
			k:        2,
			expected: []int{-1, 2},
		},
		{
			name:     "test 4",
			nums:     []int{1, 2},
			k:        2,
			expected: []int{1, 2},
		},
		{
			name:     "test 5",
			nums:     []int{5, 2, 5, 3, 5, 3, 1, 1, 3},
			k:        2,
			expected: []int{3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := topKFrequent(tt.nums, tt.k)
			if len(actual) != len(tt.expected) {
				t.Logf("Actual len %d less than expected %d", len(actual), len(tt.expected))
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
