package xoroperationinanarray

import "testing"

func TestXorOperation(t *testing.T) {
	type test struct {
		name     string
		n, start int
		expected int
	}
	tests := []test{
		{
			name:     "test 1",
			n:        5,
			start:    0,
			expected: 8,
		},
		{
			name:     "test 2",
			n:        4,
			start:    3,
			expected: 8,
		},
		{
			name:     "test 3",
			n:        1,
			start:    7,
			expected: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := xorOperation(tt.n, tt.start)
			if actual != tt.expected {
				t.Errorf("\nActual: %d\nExpected: %d", actual, tt.expected)
			}
		})
	}
}
