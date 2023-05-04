package takegiftsfromrichestpiles

import "testing"

func TestPickGifts(t *testing.T) {
	type test struct {
		name     string
		input    []int
		k        int
		expected int64
	}
	tests := []test{
		{
			name:     "test 1",
			input:    []int{25, 64, 9, 4, 100},
			k:        4,
			expected: 29,
		},
		{
			name:     "test 2",
			input:    []int{1, 1, 1, 1},
			k:        4,
			expected: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := pickGifts(tt.input, tt.k)
			if actual != tt.expected {
				t.Errorf("Actual: %d; Expeced: %d", actual, tt.expected)
			}
		})
	}
}
