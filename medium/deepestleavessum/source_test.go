package deepestleavessum

import "testing"

func TestDeepestLeavesSum(t *testing.T) {
	type test struct {
		name     string
		root     *TreeNode
		expected int
	}
	tests := []test{
		{
			name:     "test 1",
			root:     test1TreeNode(),
			expected: 15,
		},
		{
			name:     "test 2",
			root:     test2TreeNode(),
			expected: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := deepestLeavesSum(tt.root)
			if actual != tt.expected {
				t.Errorf("Actual: %d;Expected: %d", actual, tt.expected)
			}
		})
	}
}

func test1TreeNode() *TreeNode {
	/*
		a := []int{
			1,
			2, 3,
			4, 5, -1, 6,
			7, -1, -1, -1, -1, 8,
		}
	*/
	root := &TreeNode{}

	return root
}

func test2TreeNode() *TreeNode {
	/*
		a := []int{
			6,
			7, 8,
			2, 7, 1, 3,
			9, -1, 1, 4, -1, -1, -1, 5,
		}
	*/
	root := &TreeNode{}

	return root
}
