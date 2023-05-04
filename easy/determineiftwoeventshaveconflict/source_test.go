package determineiftwoeventshaveconflict

import "testing"

func TestHaveConflict(t *testing.T) {
	type test struct {
		name     string
		input1   []string
		input2   []string
		expected bool
	}
	tests := []test{
		{
			name:     "test 1",
			input1:   []string{"01:15", "02:00"},
			input2:   []string{"02:00", "03:00"},
			expected: true,
		},
		{
			name:     "test 2",
			input1:   []string{"01:00", "02:00"},
			input2:   []string{"01:20", "03:00"},
			expected: true,
		},
		{
			name:     "test 3",
			input1:   []string{"10:00", "11:00"},
			input2:   []string{"14:00", "15:00"},
			expected: false,
		},
		{
			name:     "test 4",
			input1:   []string{"16:53", "19:00"},
			input2:   []string{"10:33", "18:15"},
			expected: true,
		},
		{
			name:     "test 5",
			input1:   []string{"15:19", "17:56"},
			input2:   []string{"14:11", "20:02"},
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := haveConflict(tt.input1, tt.input2)
			if actual != tt.expected {
				t.Errorf("Actual: %v; Expected: %v", actual, tt.expected)
			}
		})
	}
}
