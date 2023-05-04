package rearrangespacesbetweenwords

import "testing"

func TestReorderSpaces(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected string
	}
	tests := []test{
		{
			name:     "test 1",
			input:    "  this   is  a sentence ",
			expected: "this   is   a   sentence",
		},
		{
			name:     "test 2",
			input:    " practice   makes   perfect",
			expected: "practice   makes   perfect ",
		},
		{
			name:     "test 3",
			input:    "a",
			expected: "a",
		},
		{
			name:     "test 4",
			input:    "  hello",
			expected: "hello  ",
		},
		{
			name:     "test 5",
			input:    "a b   c d",
			expected: "a b c d  ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := reorderSpaces(tt.input)
			if actual != tt.expected {
				t.Errorf("\n\"%s\" - actual\n\"%s\" - expected", actual, tt.expected)
			}
		})
	}
}
