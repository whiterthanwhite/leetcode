package deepestleavessum

func deepestLeavesSum(root *TreeNode) int {
	maxLeft, maxRight := 0, 0
	maxLeft = maxRank(root, maxLeft)
	maxRight = maxRank(root, maxRight)
	return getValue(root, maxLeft)
}

func maxRank(x *TreeNode, rank int) int {
	maxLeft, maxRight := rank, rank
	if x == nil {
		return maxLeft
	}
	if x.Left != nil {
		maxLeft = maxRank(x.Left, rank+1)
	}
	if x.Right != nil {
		maxRight = maxRank(x.Right, rank+1)
	}
	if maxLeft < maxRight {
		return maxRight
	} else {
		return maxLeft
	}
}

func getValue(x *TreeNode, rank int) int {
	leftVal, rightVal := 0, 0
	if rank == 0 {
		return x.Val
	}
	if x.Left != nil {
		leftVal = getValue(x.Left, rank-1)
	}
	if x.Right != nil {
		rightVal = getValue(x.Right, rank-1)
	}
	return leftVal + rightVal
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
