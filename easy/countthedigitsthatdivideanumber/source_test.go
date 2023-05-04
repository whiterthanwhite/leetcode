package countthedigitsthatdivideanumber

import "testing"

func TestCountDigits(t *testing.T) {
	type test struct {
		name     string
		input    int
		expected int
	}
	tests := []test{
		{
			name:     "test 1",
			input:    7,
			expected: 1,
		},
		{
			name:     "test 2",
			input:    121,
			expected: 2,
		},
		{
			name:     "test 3",
			input:    1248,
			expected: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := countDigits(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %d; Expected: %d\n", actual, tt.expected)
			}
		})
	}
}
