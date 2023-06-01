// 3sum
package threesum

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	type test struct {
		name     string
		nums     []int
		expected [][]int
	}
	tests := []test{
		{
			name: "test 1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name:     "test 2",
			nums:     []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			name: "test 3",
			nums: []int{0, 0, 0},
			expected: [][]int{
				{0, 0, 0},
			},
		},
		{
			name: "test 4",
			nums: []int{0, 0, 0, 0},
			expected: [][]int{
				{0, 0, 0},
			},
		},
		{
			name: "test 5",
			nums: []int{-2, 0, 1, 1, 2},
			expected: [][]int{
				{-2, 0, 2},
				{-2, 1, 1},
			},
		},
		{
			name: "test 6",
			nums: []int{-1, 0, 1, 0},
			expected: [][]int{
				{-1, 0, 1},
			},
		},
		{
			name: "test 7",
			nums: []int{3, 0, -2, -1, 1, 2},
			expected: [][]int{
				{-2, -1, 3},
				{-2, 0, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "test 8",
			nums: []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4},
			expected: [][]int{
				{-4, 0, 4},
				{-4, 1, 3},
				{-3, -1, 4},
				{-3, 0, 3},
				{-3, 1, 2},
				{-2, -1, 3},
				{-2, 0, 2},
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "test 8",
			nums: []int{34, 55, 79, 28, 46, 33, 2, 48, 31, -3, 84, 71, 52, -3, 93, 15, 21, -43, 57, -6, 86, 56, 94, 74, 83, -14, 28, -66, 46, -49, 62, -11, 43, 65, 77, 12, 47, 61, 26, 1, 13, 29, 55, -82, 76, 26, 15, -29, 36, -29, 10, -70, 69, 17, 49},
			expected: [][]int{
				{-82, -11, 93},
				{-82, 13, 69},
				{-82, 17, 65},
				{-82, 21, 61},
				{-82, 26, 56},
				{-82, 33, 49},
				{-82, 34, 48},
				{-82, 36, 46},
				{-70, -14, 84},
				{-70, -6, 76},
				{-70, 1, 69},
				{-70, 13, 57},
				{-70, 15, 55},
				{-70, 21, 49},
				{-70, 34, 36},
				{-66, -11, 77},
				{-66, -3, 69},
				{-66, 1, 65},
				{-66, 10, 56},
				{-66, 17, 49},
				{-49, -6, 55},
				{-49, -3, 52},
				{-49, 1, 48},
				{-49, 2, 47},
				{-49, 13, 36},
				{-49, 15, 34},
				{-49, 21, 28},
				{-43, -14, 57},
				{-43, -6, 49},
				{-43, -3, 46},
				{-43, 10, 33},
				{-43, 12, 31},
				{-43, 15, 28},
				{-43, 17, 26},
				{-29, -14, 43},
				{-29, 1, 28},
				{-29, 12, 17},
				{-14, -3, 17},
				{-14, 1, 13},
				{-14, 2, 12},
				{-11, -6, 17},
				{-11, 1, 10},
				{-3, 1, 2},
			},
		},
		{
			name: "test 9",
			nums: []int{0, -4, -1, -4, -2, -3, 2},
			expected: [][]int{
				{-2, 0, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := threeSum(tt.nums)
			t.Log(actual)
			if actual == nil {
				t.Log("Actual is nil")
				t.FailNow()
			}
			if len(actual) != len(tt.expected) {
				t.Logf("Actual %d differs from expected %d", len(actual), len(tt.expected))
				t.FailNow()
			}
			for i, triplet := range actual {
				for j, value := range triplet {
					if value != tt.expected[i][j] {
						t.Errorf("Actual: %d; Expected: %d; Position: i = %d, j = %d", value, tt.expected[i][j], i, j)
					}
				}
			}
		})
	}
}
