package counthillsandvalleysinanarray

func countHillValley(nums []int) int {
	hillsValleys := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		left, right := 0, 0
		for left = i - 1; left >= 0; left-- {
			if nums[left] != nums[i] {
				break
			}
		}
		for right = i + 1; right < n; right++ {
			if nums[right] != nums[i] {
				break
			}
		}
		if left >= 0 && right < n {
			if nums[i] != nums[i-1] {
				if nums[left] < nums[i] && nums[i] > nums[right] {
					hillsValleys++
				} else if nums[left] > nums[i] && nums[i] < nums[right] {
					hillsValleys++
				}
			}
		}
	}
	return hillsValleys
}
