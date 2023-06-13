package generateparentheses

import "testing"

func TestGenerateParenthesis(t *testing.T) {
	type test struct {
		name     string
		input    int
		expected []string
	}
	tests := [...]test{
		{
			name:     "test 1",
			input:    3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name:     "test 2",
			input:    1,
			expected: []string{"()"},
		},
		{
			name:     "test 3",
			input:    2,
			expected: []string{"(())", "()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := generateParenthesis(tt.input)
			if len(actual) != len(tt.expected) {
				t.Logf("Actual len is %d; Expected len is %d", len(actual), len(tt.expected))
				t.FailNow()
			}
			for i, v := range actual {
				if v != tt.expected[i] {
					t.Errorf("Actual: %s; Expected: %s", v, tt.expected[i])
				}
			}
		})
	}
}
