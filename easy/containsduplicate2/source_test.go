package containsduplicate2

import "testing"

func TestContainNearbyDuplicate(t *testing.T) {
	type test struct {
		name   string
		src    []int
		k      int
		wanted bool
	}
	tests := []test{
		{
			name:   "test 1",
			src:    []int{1, 2, 3, 1},
			k:      3,
			wanted: true,
		},
		{
			name:   "test 2",
			src:    []int{1, 0, 1, 1},
			k:      1,
			wanted: true,
		},
		{
			name:   "test 3",
			src:    []int{1, 2, 3, 1, 2, 3},
			k:      2,
			wanted: false,
		},
		{
			name:   "test 4",
			src:    []int{99, 99},
			k:      2,
			wanted: true,
		},
		{
			name:   "test 5",
			src:    []int{1},
			k:      1,
			wanted: false,
		},
		{
			name:   "test 6",
			src:    []int{1, 2},
			k:      2,
			wanted: false,
		},
		{
			name:   "test 7",
			src:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9},
			k:      3,
			wanted: true,
		},
		{
			name:   "test 8",
			src:    []int{0, 1, 2, 3, 2, 5},
			k:      3,
			wanted: true,
		},
		{
			name:   "test 9",
			src:    []int{1, 2, 1},
			k:      0,
			wanted: false,
		},
		{
			name:   "test 10",
			src:    []int{1, 2, 1},
			k:      1,
			wanted: false,
		},
		{
			name:   "test 11",
			src:    []int{4, 1, 2, 3, 1, 5},
			k:      3,
			wanted: true,
		},
		{
			name:   "test 12",
			src:    []int{-1, -1},
			k:      1,
			wanted: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := containsNearbyDuplicate(tt.src, tt.k)
			if actual != tt.wanted {
				t.Errorf("Actual: %v, but wanted: %v", actual, tt.wanted)
			}
		})
	}
}
