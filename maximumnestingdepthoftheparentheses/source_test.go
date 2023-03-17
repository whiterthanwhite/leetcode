package maximumnestingdepthoftheparentheses

import "testing"

func TestMaxDepth(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected int
	}
	tests := []test{
		{
			name:     "test 1",
			input:    "(1+(2*3)+((8)/4))+1",
			expected: 3,
		},
		{
			name:     "test 2",
			input:    "(1)+((2))+(((3)))",
			expected: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := maxDepth(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %d\nExpected: %d\n", actual, tt.expected)
			}
		})
	}
}
