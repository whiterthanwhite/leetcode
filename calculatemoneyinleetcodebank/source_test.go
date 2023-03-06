package calculatemoneyinleetcodebank

import (
	"testing"
	"time"
)

func TestTotalMoney(t *testing.T) {
	type test struct {
		name        string
		n, expected int
	}
	tests := []test{
		{
			name:     "test 1",
			n:        4,
			expected: 10,
		},
		{
			name:     "test 2",
			n:        10,
			expected: 37,
		},
		{
			name:     "test 3",
			n:        20,
			expected: 96,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startTime := time.Now()
			actual := totalMoney(tt.n)
			if actual != tt.expected {
				t.Errorf("Actual: %d; but expected: %d", actual, tt.expected)
			}
			t.Log("Execution time:", time.Now().Sub(startTime))
		})
	}
}
