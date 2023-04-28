package sumofuniqueelements

import "testing"

func TestSumOfUnique(t *testing.T) {
	type test struct {
		name     string
		input    []int
		expected int
	}
	tests := []test{
		{
			name:     "test 1",
			input:    []int{1, 2, 3, 2},
			expected: 4,
		},
		{
			name:     "test 2",
			input:    []int{1, 1, 1, 1, 1},
			expected: 0,
		},
		{
			name:     "test 3",
			input:    []int{1, 2, 3, 4, 5},
			expected: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := sumOfUnique(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %d; Expected: %d", actual, tt.expected)
			}
		})
	}
}
