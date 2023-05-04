package bitand2bitcharacters

import "testing"

func TestIsOneBitCharacter(t *testing.T) {
	type test struct {
		name     string
		input    []int
		expected bool
	}
	tests := []test{
		{
			name:     "test 1",
			input:    []int{1, 0, 0},
			expected: true,
		},
		{
			name:     "test 2",
			input:    []int{1, 1, 1, 0},
			expected: false,
		},
		{
			name:     "test 3",
			input:    []int{1, 1, 0},
			expected: true,
		},
		{
			name:     "test 4",
			input:    []int{0, 1, 0},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := isOneBitCharacter(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %v; Expected: %v", actual, tt.expected)
			}
		})
	}
}
