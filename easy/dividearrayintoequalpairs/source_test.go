package dividearrayintoequalpairs

import "testing"

func TestDivideArray(t *testing.T) {
	type test struct {
		name     string
		input    []int
		expected bool
	}
	tests := []test{
		{
			name:     "test 1",
			input:    []int{3, 2, 3, 2, 2, 2},
			expected: true,
		},
		{
			name:     "test 2",
			input:    []int{1, 2, 3, 4},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := divideArray(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %v; Expected: %v", actual, tt.expected)
			}
		})
	}
}
