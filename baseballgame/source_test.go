package baseballgame

import "testing"

func TestCalPoints(t *testing.T) {
	type test struct {
		name     string
		input    []string
		expected int
	}
	tests := []test{
		{
			name:     "test 1",
			input:    []string{"5", "2", "C", "D", "+"},
			expected: 30,
		},
		{
			name:     "test 2",
			input:    []string{"5", "-2", "4", "C", "D", "9", "+", "+"},
			expected: 27,
		},
		{
			name:     "test 3",
			input:    []string{"1", "C"},
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := calPoints(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %d\nExpected: %d", actual, tt.expected)
			}
		})
	}
}
