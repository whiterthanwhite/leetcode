package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected bool
	}
	tests := []test{
		{
			name:     "test 1",
			input:    "()",
			expected: true,
		},
		{
			name:     "test 2",
			input:    "()[]{}",
			expected: true,
		},
		{
			name:     "test 3",
			input:    "(]",
			expected: false,
		},
		{
			name:     "test 4",
			input:    "([)]",
			expected: false,
		},
		{
			name:     "test 5",
			input:    "[",
			expected: false,
		},
		{
			name:     "test 6",
			input:    "({{{{}}}))",
			expected: false,
		},
		{
			name:     "test 7",
			input:    "([]){",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := isValid(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %v; Expected: %v", actual, tt.expected)
			}
		})
	}
}
