package evaluatereversepolishnotation

import "testing"

func TestEvalRPN(t *testing.T) {
	type test struct {
		name     string
		tokens   []string
		expected int
	}
	tests := [...]test{
		{
			name:     "test 1",
			tokens:   []string{"2", "1", "+", "3", "*"},
			expected: 9,
		},
		{
			name:     "test 2",
			tokens:   []string{"4", "13", "5", "/", "+"},
			expected: 6,
		},
		{
			name:     "test 3",
			tokens:   []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			expected: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := evalRPN(tt.tokens)
			if actual != tt.expected {
				t.Errorf("Actual: %d; Expected: %d", actual, tt.expected)
			}
		})
	}
}
