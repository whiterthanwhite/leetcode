package productofarrayexceptself

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	prefixNums := make([]int, n)
	prefixNums[0] = nums[0]
	for i := 1; i < n-1; i++ {
		prefixNums[i] = prefixNums[i-1] * nums[i]
	}
	suffixNums := make([]int, n)
	suffixNums[n-1] = nums[n-1]
	for i := n - 2; i > 0; i-- {
		suffixNums[i] = suffixNums[i+1] * nums[i]
	}

	result[0] = suffixNums[1]
	result[n-1] = prefixNums[n-2]
	for i := 1; i < n-1; i++ {
		result[i] = prefixNums[i-1] * suffixNums[i+1]
	}

	return result
}
