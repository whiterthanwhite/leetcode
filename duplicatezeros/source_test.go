package duplicatezeros

import "testing"

func TestDuplicateZeros(t *testing.T) {
	type test struct {
		name     string
		input    []int
		expected []int
	}
	tests := []test{
		{
			name:     "test 1",
			input:    []int{1, 0, 2, 3, 0, 4, 5, 0},
			expected: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			name:     "test 2",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duplicateZeros(tt.input)
			for i, v := range tt.input {
				if v != tt.expected[i] {
					t.Errorf("Actual: %d; Expected: %d; At position: %d", v, tt.expected[i], i)
				}
			}
		})
	}
}
