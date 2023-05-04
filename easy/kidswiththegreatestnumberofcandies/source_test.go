package kidswiththegreatestnumberofcandies

import "testing"

func TestKidsWithCandies(t *testing.T) {
	type test struct {
		name         string
		candies      []int
		extraCandies int
		expected     []bool
	}
	tests := []test{
		{
			name:         "test 1",
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			expected:     []bool{true, true, true, false, true},
		},
		{
			name:         "test 2",
			candies:      []int{4, 2, 1, 1, 2},
			extraCandies: 1,
			expected:     []bool{true, false, false, false, false},
		},
		{
			name:         "test 3",
			candies:      []int{12, 1, 12},
			extraCandies: 10,
			expected:     []bool{true, false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := kidsWithCandies(tt.candies, tt.extraCandies)
			if actual == nil {
				t.Fatal("Actual should not be nil")
			}
			for i, v := range actual {
				if v != tt.expected[i] {
					t.Errorf("Actual: %v; Expected: %v", v, tt.expected[i])
				}
			}
		})
	}
}
