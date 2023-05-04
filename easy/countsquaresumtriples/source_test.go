package countsquaresumtriples

import "testing"

func TestCountTriples(t *testing.T) {
	type test struct {
		name     string
		input    int
		expected int
	}
	tests := []test{
		{
			name:     "test 1",
			input:    5,
			expected: 2,
		},
		{
			name:     "test 2",
			input:    10,
			expected: 4,
		},
		{
			name:     "test 3",
			input:    55,
			expected: 48,
		},
		{
			name:     "test 4",
			input:    110,
			expected: 118,
		},
		{
			name:     "test 5",
			input:    116,
			expected: 126,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := countTriples(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %d; Expected: %d", actual, tt.expected)
			}
		})
	}
}
