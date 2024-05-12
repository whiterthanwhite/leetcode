package besttimetobuyandsellstockwithcooldown

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	type test struct {
		name     string
		prices   []int
		expected int
	}
	tests := []test{
		{
			name:     "test 1",
			prices:   []int{1, 2, 3, 0, 2},
			expected: 3,
		},
		{
			name:     "test 2",
			prices:   []int{1},
			expected: 0,
		},
		{
			name:     "test 3",
			prices:   []int{1, 2, 4},
			expected: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := maxProfit(tt.prices)
			if actual != tt.expected {
				t.Errorf("Actual: %d; Expected: %d", actual, tt.expected)
			}
		})
	}
}
