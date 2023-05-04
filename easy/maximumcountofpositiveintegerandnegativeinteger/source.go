package maximumcountofpositiveintegerandnegativeinteger

func maximumCount(nums []int) int {
	N := len(nums)
	pos, neg := 0, 0
	for N > 0 && nums[N-1] > 0 {
		pos++
		N--
	}
	for i := 0; i < N && nums[i] < 0; i++ {
		neg++
	}
	if pos < neg {
		return neg
	} else {
		return pos
	}
}
