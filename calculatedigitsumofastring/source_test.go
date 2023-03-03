package calculatedigitsumofastring

import (
	"testing"
	"time"
)

func TestDigitSum(t *testing.T) {
	type test struct {
		name            string
		k               int
		input, expected string
	}
	tests := []test{
		{
			name:     "test 1",
			input:    "11111222223",
			k:        3,
			expected: "135",
		},
		{
			name:     "test 2",
			input:    "00000000",
			k:        3,
			expected: "000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startTime := time.Now()
			actual := digitSum(tt.input, tt.k)
			if actual != tt.expected {
				t.Errorf("Actual: %s, but expected: %s\n", actual, tt.expected)
			}
			t.Logf("Execution time: %v", time.Now().Sub(startTime))
		})
	}
}
