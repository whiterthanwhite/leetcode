package productofarrayexceptself

import "testing"

func TestProductExceptSelf(t *testing.T) {
	type test struct {
		name     string
		nums     []int
		expected []int
	}
	tests := []test{
		{
			name:     "test 1",
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			name:     "test 2",
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := productExceptSelf(tt.nums)
			if len(actual) != len(tt.expected) {
				t.Logf("Actual len %d differs from expected len %d", len(actual), len(tt.expected))
				t.FailNow()
			}
			for i, v := range actual {
				if v != tt.expected[i] {
					t.Errorf("Actual: %d; Expected: %d; Position: %d", v, tt.expected[i], i)
				}
			}
		})
	}
}
