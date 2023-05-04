package removealladjacentduplicatesinstring

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected string
	}
	tests := []test{
		{
			name:     "test 1",
			input:    "abbaca",
			expected: "ca",
		},
		{
			name:     "test 2",
			input:    "azxxzy",
			expected: "ay",
		},
		{
			name:     "test 3",
			input:    "aaaaaaaa",
			expected: "",
		},
		{
			name:     "test 4",
			input:    "aaaaaaaaa",
			expected: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := removeDuplicates(tt.input)
			if actual != tt.expected {
				t.Errorf("Actual: %s; Expected: %s", actual, tt.expected)
			}
		})
	}
}
