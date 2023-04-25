package removeelement

func removeElement(nums []int, val int) int {
	N := len(nums)
	var r []int
	for i := 0; i < N; i++ {
		if nums[i] != val {
			r = append(r, nums[i])
		}
	}
	for i, v := range r {
		nums[i] = v
	}
	return len(r)
}
