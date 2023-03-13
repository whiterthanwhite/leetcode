package finalvalueofvariableafterperformingoperations

import (
	"testing"
)

func TestFinalValueAfterOperations(t *testing.T) {
	type test struct {
		name       string
		operations []string
		expected   int
	}
	tests := []test{
		{
			name:       "test 1",
			operations: []string{"--X", "X++", "X++"},
			expected:   1,
		},
		{
			name:       "test 2",
			operations: []string{"++X", "++X", "X++"},
			expected:   3,
		},
		{
			name:       "test 3",
			operations: []string{"X++", "++X", "--X", "X--"},
			expected:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := finalValueAfterOperations(tt.operations)
			if actual != tt.expected {
				t.Errorf("Actual: %d; Expected: %d", actual, tt.expected)
			}
		})
	}
}
