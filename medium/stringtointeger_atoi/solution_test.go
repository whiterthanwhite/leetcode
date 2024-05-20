package stringtointegeratoi

import "testing"

func TestMyAtoi(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected int
	}
	tests := [...]test{
		{
			name:     "Test 1",
			input:    "42",
			expected: 42,
		},
		{
			name:     "Test 2",
			input:    "-042",
			expected: -42,
		},
		{
			name:     "Test 3",
			input:    "1337c0d3",
			expected: 1337,
		},
		{
			name:     "Test 4",
			input:    "0-1",
			expected: 0,
		},
		{
			name:     "Test 5",
			input:    "words and 987",
			expected: 0,
		},
		{
			name:     "Test 6",
			input:    "  0002147483648",
			expected: 2147483647,
		},
		{
			name:     "Test 7",
			input:    "  -0002147483649",
			expected: -2147483648,
		},
		{
			name:     "Test 8",
			input:    "3.14159",
			expected: 3,
		},
		{
			name:     "Test 9",
			input:    "21474836460",
			expected: 2147483647,
		},
		{
			name:     "Test 10",
			input:    "   +0 123",
			expected: 0,
		},
		{
			name:     "Test 11",
			input:    "0  123",
			expected: 0,
		},
		{
			name:     "Test 12",
			input:    "-5-",
			expected: -5,
		},
		{
			name:     "Test 13",
			input:    "123-",
			expected: 123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := myAtoi(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %v; Expected: %v", actual, tt.expected)
			}
		})
	}
}
