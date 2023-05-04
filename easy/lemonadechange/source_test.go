package lemonadechange

import "testing"

func TestLemonadeChange(t *testing.T) {
	type test struct {
		name     string
		input    []int
		expected bool
	}
	tests := []test{
		{
			name:     "test 1",
			input:    []int{5, 5, 5, 10, 20},
			expected: true,
		},
		{
			name:     "test 2",
			input:    []int{5, 5, 10, 10, 20},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := lemonadeChange(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %v; Expected: %v", actual, tt.expected)
			}
		})
	}
}
