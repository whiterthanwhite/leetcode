package zigzagconversion

import "testing"

func TestConvert(t *testing.T) {
	type test struct {
		name     string
		str      string
		numRows  int
		expected string
	}
	tests := [...]test{
		{
			name:     "Test 1",
			str:      "PAYPALISHIRING",
			numRows:  3,
			expected: "PAHNAPLSIIGYIR",
		},
		/*
			P   A   H   N
			A P L S I I G
			Y   I   R
		*/
		{
			name:     "Test 2",
			str:      "PAYPALISHIRING",
			numRows:  4,
			expected: "PINALSIGYAHRPI",
		},
		/*
			P     I     N
			A   L S   I G
			Y A   H R
			P     I
		*/
		{
			name:     "Test 3",
			str:      "A",
			numRows:  1,
			expected: "A",
		},
		{
			name:     "Test 4",
			str:      "PAYPALISHIRING",
			numRows:  5,
			expected: "PHASIYIRPLIGAN",
		},
		/*
			P       H
			A     S I
			Y   I   R
			P L     I G
			A       N
		*/
		{
			name:     "Test 5",
			str:      "ABCDEF",
			numRows:  6,
			expected: "ABCDEF",
		},
		{
			name:     "Test 6",
			str:      "A",
			numRows:  2,
			expected: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := convert(tt.str, tt.numRows)
			if actual != tt.expected {
				t.Errorf("Expected: \"%s\"; Actual: \"%s\"", tt.expected, actual)
			}
		})
	}
}
