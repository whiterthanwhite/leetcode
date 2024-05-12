package uncrossedlines

import "testing"

func TestMaxUncrossedLines(t *testing.T) {
	type test struct {
		name     string
		nums1    []int
		nums2    []int
		expected int
	}
	tests := []test{
		{
			name:     "test 1",
			nums1:    []int{1, 4, 2},
			nums2:    []int{1, 2, 4},
			expected: 2,
		},
		{
			name:     "test 2",
			nums1:    []int{2, 5, 1, 2, 5},
			nums2:    []int{10, 5, 2, 1, 5, 2},
			expected: 3,
		},
		{
			name:     "test 3",
			nums1:    []int{1, 3, 7, 1, 7, 5},
			nums2:    []int{1, 9, 2, 5, 1},
			expected: 2,
		},
		{
			name:     "test 4",
			nums1:    []int{1, 4, 2},
			nums2:    []int{1, 2, 3, 4},
			expected: 2,
		},
		{
			name:     "test 5",
			nums1:    []int{1, 1, 2, 1, 2},
			nums2:    []int{1, 3, 2, 3, 1},
			expected: 3,
		},
		{
			name:     "test 6",
			nums1:    []int{1},
			nums2:    []int{3},
			expected: 0,
		},
		{
			name:     "test 7",
			nums1:    []int{3, 3},
			nums2:    []int{3},
			expected: 1,
		},
		{
			name:     "test 8",
			nums1:    []int{3, 3, 3, 1},
			nums2:    []int{2, 2, 3, 3},
			expected: 2,
		},
		{
			name:     "test 9",
			nums1:    []int{1, 1, 3, 5, 3, 3, 5, 5, 1, 1},
			nums2:    []int{2, 3, 2, 1, 3, 5, 3, 2, 2, 1},
			expected: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := maxUncrossedLines(tt.nums1, tt.nums2)
			if actual != tt.expected {
				t.Errorf("Actual: %d; Expected: %d", actual, tt.expected)
			}
		})
	}
}
