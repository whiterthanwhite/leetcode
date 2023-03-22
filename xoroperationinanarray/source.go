package xoroperationinanarray

import "fmt"

func xorOperation(n, start int) int {
	nums := make([]int, n)
	var r int
	for i := 0; i < n; i++ {
		nums[i] = start + 2*i
		if i == 1 {
			r = nums[i-1] ^ nums[i]
		} else if i > 1 {
			r = r ^ nums[i]
		} else {
			r = 0 ^ nums[i]
		}
	}
	fmt.Println(nums, r)
	return r
}
