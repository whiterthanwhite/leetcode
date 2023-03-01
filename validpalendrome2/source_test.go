package validPalindrome2

import (
	"testing"
	"time"
)

func TestValidPalindrome(t *testing.T) {
	type test struct {
		name, input string
		expected    bool
	}
	tests := []test{
		{
			name:     "test 1",
			input:    "aba",
			expected: true,
		},
		{
			name:     "test 2",
			input:    "abca",
			expected: true,
		},
		{
			name:     "test 3",
			input:    "abc",
			expected: false,
		},
		{
			name:     "test 4",
			input:    "abbca",
			expected: true,
		},
		{
			name:     "test 5",
			input:    "acbba",
			expected: true,
		},
		{
			name:     "test 6",
			input:    "cabba",
			expected: true,
		},
		{
			name:     "test 7",
			input:    "abbac",
			expected: true,
		},
		{
			name:     "test 8",
			input:    "atbbga",
			expected: false,
		},
		{
			name:     "test 9",
			input:    "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga",
			expected: true,
		},
		{
			name:     "test 10",
			input:    "deeee",
			expected: true,
		},
		{
			name:     "test 11",
			input:    "aydmda",
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startTime := time.Now()
			actual := validPalindrome(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %t; Expected: %t\n", actual, tt.expected)
			}
			t.Log("Execution time:", time.Now().Sub(startTime))
		})
	}
}
