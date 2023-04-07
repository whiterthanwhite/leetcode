package decodethemessage

import "testing"

func TestDecodeMessage(t *testing.T) {
	type test struct {
		name     string
		key      string
		message  string
		expected string
	}
	tests := []test{
		{
			name:     "test 1",
			key:      "the quick brown fox jumps over the lazy dog",
			message:  "vkbs bs t suepuv",
			expected: "this is a secret",
		},
		{
			name:     "test 2",
			key:      "eljuxhpwnyrdgtqkviszcfmabo",
			message:  "zwx hnfx lqantp mnoeius ycgk vcnjrdb",
			expected: "the five boxing wizards jump quickly",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := decodeMessage(tt.key, tt.message)
			if actual != tt.expected {
				t.Errorf("\nActual: %s\nExpected: %s\n", actual, tt.expected)
			}
		})
	}
}
