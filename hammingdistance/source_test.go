package hammingdistance

import (
	"testing"
	"time"
)

func TestHammingDistance(t *testing.T) {
	type test struct {
		name           string
		x, y, expected int
	}
	tests := []test{
		{
			name:     "test 1",
			x:        1,
			y:        4,
			expected: 2,
		},
		{
			name:     "test 2",
			x:        3,
			y:        1,
			expected: 1,
		},
		{
			name:     "test 3",
			x:        680142203,
			y:        1111953568,
			expected: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startTime := time.Now()
			actual := hammingDistance(tt.x, tt.y)
			if actual != tt.expected {
				t.Errorf("Actual: %d, but expected: %d\n", actual, tt.expected)
			}
			t.Logf("Execution time: %v", startTime.Sub(time.Now()))
		})
	}
}
