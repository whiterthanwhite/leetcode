package containsduplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	type test struct {
		name     string
		input    []int
		expected bool
	}
	tests := []test{
		{
			name:     "test 1",
			input:    []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "test 2",
			input:    []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "test 3",
			input:    []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
		{
			name:     "test 4",
			input:    []int{2, 14, 18, 22, 22},
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := containsDuplicate(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %t; Expected: %t", actual, tt.expected)
			}
		})
	}
}
