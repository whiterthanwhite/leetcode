package slowestkey

import "testing"

func TestSlowestKey(t *testing.T) {
	type test struct {
		name         string
		releaseTimes []int
		keysPressed  string
		expected     byte
	}
	tests := []test{
		{
			name:         "test 1",
			releaseTimes: []int{9, 29, 49, 50},
			keysPressed:  "cbcd",
			expected:     'c',
		},
		{
			name:         "test 2",
			releaseTimes: []int{12, 23, 36, 46, 62},
			keysPressed:  "spuda",
			expected:     'a',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := slowestKey(tt.releaseTimes, tt.keysPressed)
			if actual != tt.expected {
				t.Errorf("Actual %d; Expected: %d\n", actual, tt.expected)
			}
		})
	}
}
