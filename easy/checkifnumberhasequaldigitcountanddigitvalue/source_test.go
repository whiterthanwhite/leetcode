package checkifnumberhasequaldigitcountanddigitvalue

import "testing"

func TestDigitCount(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected bool
	}
	tests := []test{
		{
			name:     "test 1",
			input:    "1210",
			expected: true,
		},
		{
			name:     "test 2",
			input:    "030",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := digitCount(tt.input)
			if actual != tt.expected {
				t.Errorf("\nActual: %v\nExpected: %v", actual, tt.expected)
			}
		})
	}
}
