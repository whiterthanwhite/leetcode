package fizzbuzz

import (
	"testing"
	"time"
)

func TestFizzBuzz(t *testing.T) {
	type test struct {
		name     string
		n        int
		expected []string
	}
	tests := []test{
		{
			name:     "test 1",
			n:        3,
			expected: []string{"1", "2", "Fizz"},
		},
		{
			name:     "test 2",
			n:        5,
			expected: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name:     "test 3",
			n:        15,
			expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startTime := time.Now()
			actual := fizzBuzz(tt.n)
			if len(actual) != len(tt.expected) {
				t.Errorf("Difference in arrays")
			} else {
				for i := 0; i < len(actual); i++ {
					if actual[i] != tt.expected[i] {
						t.Errorf("Position: %d, actual: %s, expected: %s", i, actual[i], tt.expected[i])
					}
				}
			}
			t.Logf("Execution time: %v", time.Now().Sub(startTime))
		})
	}
}
