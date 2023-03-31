package dividearrayintoequalpairs

func divideArray(nums []int) bool {
	N := len(nums)
	h := 1
	for h < N/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j >= h && nums[j] < nums[j-h]; j -= h {
				t := nums[j]
				nums[j] = nums[j-h]
				nums[j-h] = t
			}
		}
		h = h / 3
	}
	for i := 0; i < N-1; i += 2 {
		if nums[i] != nums[i+1] {
			return false
		}
	}
	return true
}
