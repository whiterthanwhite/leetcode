package containerwithmostwater

import "testing"

func TestMaxArea(t *testing.T) {
	type test struct {
		name     string
		height   []int
		expected int
	}
	tests := []test{
		{
			name:     "test 1",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "test 2",
			height:   []int{1, 1},
			expected: 1,
		},
		{
			name:     "test 3",
			height:   []int{1, 2, 1},
			expected: 2,
		},
		{
			name:     "test 4",
			height:   []int{1, 2},
			expected: 1,
		},
		{
			name:     "test 5",
			height:   []int{1, 2, 4, 3},
			expected: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := maxArea(tt.height)
			if actual != tt.expected {
				t.Errorf("Actual %d; Expected: %d", actual, tt.expected)
			}
		})
	}
}
