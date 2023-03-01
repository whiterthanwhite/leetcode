package containsduplicate2

func containsNearbyDuplicate(nums []int, k int) bool {
	N := len(nums)
	m := make(map[int]int, N)
	for i := 0; i < N; i++ {
		j := nums[i]
		v, ok := m[j]
		if ok && i-v <= k {
			return true
		} else {
			m[j] = i
		}
	}
	return false
}
