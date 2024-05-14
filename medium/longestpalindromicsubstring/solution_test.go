package longestpalindromicsubstring

import "testing"

func TestLongestPalindrome(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected string
	}
	tests := []test{
		{
			name:     "Test 1",
			input:    "babad",
			expected: "bab",
			// expected: "aba"
		},
		{
			name:     "Test 2",
			input:    "cbbd",
			expected: "bb",
		},
		{
			name:     "Test 3",
			input:    "bnqctsabaq",
			expected: "aba",
		},
		{
			name:     "Test 4",
			input:    "a",
			expected: "a",
		},
		{
			name:     "Test 5",
			input:    "ac",
			expected: "a",
		},
		{
			name:     "Test 6",
			input:    "bb",
			expected: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := longestPalindrome(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %s; Expected: %s", actual, tt.expected)
			}
		})
	}
}
