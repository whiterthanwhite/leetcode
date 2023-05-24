package productofarrayexceptself

func productExceptSelf(nums []int) []int {
	n := len(nums)

	suffixNums := make([]int, n)
	suffixNums[n-1] = nums[n-1]
	for i := n - 2; i > 0; i-- {
		suffixNums[i] = suffixNums[i+1] * nums[i]
	}

	result := make([]int, n)
	prefixProd := nums[0]
	result[0] = suffixNums[1]
	for i := 1; i < n-1; i++ {
		result[i] = prefixProd * suffixNums[i+1]
		prefixProd *= nums[i]
	}
	result[n-1] = prefixProd

	return result
}
