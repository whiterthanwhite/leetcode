package containsduplicate

func containsDuplicate(nums []int) bool {
	n := len(nums)
	duplicates := make(map[int]struct{}, n)
	for i := 0; i < n; i++ {
		_, ok := duplicates[nums[i]]
		if ok {
			return true
		} else {
			duplicates[nums[i]] = struct{}{}
		}
	}
	return false
}
