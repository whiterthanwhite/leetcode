package anumberafterdoublereversal

import "testing"

func TestIsSameAfterReversal(t *testing.T) {
	type test struct {
		name     string
		input    int
		expected bool
	}
	tests := []test{
		{
			name:     "test 1",
			input:    526,
			expected: true,
		},
		{
			name:     "test 2",
			input:    1800,
			expected: false,
		},
		{
			name:     "test 3",
			input:    0,
			expected: true,
		},
		{
			name:     "test 4",
			input:    100,
			expected: false,
		},
		{
			name:     "test 5",
			input:    1000,
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := isSameAfterReversals(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %v, Expected: %v", actual, tt.expected)
			}
		})
	}
}
