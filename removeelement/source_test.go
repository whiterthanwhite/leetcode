package removeelement

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	type test struct {
		name          string
		nums          []int
		val           int
		expected      int
		expectedOrder []int
	}
	tests := []test{
		{
			name:          "test 1",
			nums:          []int{3, 2, 2, 3},
			val:           3,
			expected:      2,
			expectedOrder: []int{2, 2},
		},
		{
			name:          "test 2",
			nums:          []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:           2,
			expected:      5,
			expectedOrder: []int{0, 1, 3, 0, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := removeElement(tt.nums, tt.val)
			t.Log(tt.nums)
			if actual != tt.expected {
				t.Fatalf("Actual: %d; Expected: %d", actual, tt.expected)
			}
			for i := 0; i < actual; i++ {
				if tt.nums[i] != tt.expectedOrder[i] {
					t.Errorf("Pos: %d; Actual: %d; Expected: %d", i, tt.nums[i], tt.expectedOrder[i])
				}
			}
		})
	}
}
