package reverseinteger

import "testing"

func TestReverse(t *testing.T) {
	type test struct {
		Name     string
		Input    int
		Expected int
	}
	tests := [...]*test{
		{
			Name:     "Test 1",
			Input:    123,
			Expected: 321,
		},
		{
			Name:     "Test 2",
			Input:    -123,
			Expected: -321,
		},
		{
			Name:     "Test 3",
			Input:    120,
			Expected: 21,
		},
		{
			Name:     "Test 4",
			Input:    1534236469,
			Expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			actual := reverse(tt.Input)
			if actual != tt.Expected {
				t.Errorf("Actual: %v; Expected: %v", actual, tt.Expected)
			}
		})
	}
}
