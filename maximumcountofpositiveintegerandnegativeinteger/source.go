package maximumcountofpositiveintegerandnegativeinteger

func maximumCount(nums []int) int {
	var pos, neg int32
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			pos++
		} else if nums[i] < 0 {
			neg++
		}
	}
	if pos < neg {
		return int(neg)
	} else {
		return int(pos)
	}
}
