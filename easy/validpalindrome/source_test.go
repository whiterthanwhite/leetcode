package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected bool
	}
	tests := []test{
		{
			name:     "test 1",
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "test 2",
			input:    "race a car",
			expected: false,
		},
		{
			name:     "test 3",
			input:    " ",
			expected: true,
		},
		{
			name:     "test 4",
			input:    "0P",
			expected: false,
		},
		{
			name:     "test 5",
			input:    "a",
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := isPalindrome(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %t; Expected: %t", actual, tt.expected)
			}
		})
	}
}
