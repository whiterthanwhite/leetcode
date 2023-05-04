package distringmatch

import "testing"

func TestDIStringMatch(t *testing.T) {
	type test struct {
		name     string
		input    string
		expected []int
	}
	tests := []test{
		{
			input:    "IDID",
			name:     "test 1",
			expected: []int{0, 4, 1, 3, 2},
		},
		{
			name:     "test 2",
			input:    "III",
			expected: []int{0, 1, 2, 3},
		},
		{
			name:     "test 3",
			input:    "DDI",
			expected: []int{3, 2, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := diStringMatch(tt.input)
			if len(actual) != len(tt.expected) {
				t.Fatal("Slices are different")
			}
			for i, v := range actual {
				if v != tt.expected[i] {
					t.Errorf("Actual: %d; Expected: %d", v, tt.expected[i])
				}
			}
		})
	}
}
