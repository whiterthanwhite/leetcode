package minimumbitflipstoconvertnumber

import "testing"

func TestMinBitFlips(t *testing.T) {
	type test struct {
		name     string
		start    int
		goal     int
		expected int
	}

	tests := []test{
		{
			name:     "test 1",
			start:    10,
			goal:     7,
			expected: 3,
		},
		{
			name:     "test 2",
			start:    3,
			goal:     4,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := minBitFlips(tt.start, tt.goal)
			if actual != tt.expected {
				t.Errorf("Actual: %d; Expected: %d", actual, tt.expected)
			}
		})
	}
}
