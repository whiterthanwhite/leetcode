package insertationoftwoarrays

import "testing"

func TestIntersection(t *testing.T) {
	type test struct {
		name         string
		nums1, nums2 []int
		expected     []int
	}
	tests := []test{
		{
			name:     "test 1",
			nums1:    []int{1, 2, 2, 1},
			nums2:    []int{2, 2},
			expected: []int{2},
		},
		{
			name:     "test 2",
			nums1:    []int{4, 9, 5},
			nums2:    []int{9, 4, 9, 8, 4},
			expected: []int{4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := intersection(tt.nums1, tt.nums2)
			if len(actual) != len(tt.expected) {
				t.Fatal("Arrays sizes are different!")
			}
			for i := 0; i < len(actual); i++ {
				if actual[i] != tt.expected[i] {
					t.Errorf("Actual: %d; Expected: %d", actual[i], tt.expected[i])
				}
			}
		})
	}
}
