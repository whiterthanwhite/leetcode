package validanagram

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	type test struct {
		name     string
		s        string
		t        string
		expected bool
	}
	tests := []test{
		{
			name:     "test 1",
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			name:     "test 2",
			s:        "rat",
			t:        "car",
			expected: false,
		},
		{
			name:     "test 3",
			s:        "a",
			t:        "ab",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := isAnagram(tt.s, tt.t)
			if actual != tt.expected {
				t.Errorf("Actual: %v; Expected: %v\n", actual, tt.expected)
			}
		})
	}
}
