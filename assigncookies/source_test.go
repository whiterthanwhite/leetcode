package assigncookies

import (
	"testing"
	"time"
)

func TestFindContentChildren(t *testing.T) {
	type test struct {
		name     string
		g, s     []int
		expected int
	}
	tests := []test{
		{
			name:     "test 1",
			g:        []int{1, 2, 3},
			s:        []int{1, 1},
			expected: 1,
		},
		{
			name:     "test 2",
			g:        []int{1, 2},
			s:        []int{1, 2, 3},
			expected: 2,
		},
		{
			name:     "test 3",
			g:        []int{1, 2, 3},
			s:        []int{},
			expected: 0,
		},
		{
			name:     "test 4",
			g:        []int{1, 2, 3},
			s:        []int{3},
			expected: 1,
		},
		{
			name:     "test 5",
			g:        []int{10, 9, 8, 7},
			s:        []int{5, 6, 7, 8},
			expected: 2,
		},
		{
			name:     "test 6",
			g:        []int{10, 9, 8, 7, 10, 9, 8, 7},
			s:        []int{10, 9, 8, 7},
			expected: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startTime := time.Now()
			actual := findContentChildren(tt.g, tt.s)
			if actual != tt.expected {
				t.Errorf("Actual: %d, but expected: %d\n", actual, tt.expected)
			}
			t.Logf("Execution time: %v", time.Now().Sub(startTime))
		})
	}
}
