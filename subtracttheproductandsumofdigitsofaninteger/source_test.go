package subtracttheproductandsumofdigitsofaninteger

import (
	"testing"
)

func TestSubtractProductAndSum(t *testing.T) {
	type test struct {
		name     string
		input    int
		expected int
	}
	tests := []test{
		{
			name:     "test 1",
			input:    234,
			expected: 15,
		},
		{
			name:     "test 2",
			input:    4421,
			expected: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := subtractProductAndSum(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %d; Expected: %d", actual, tt.expected)
			}
		})
	}
}
