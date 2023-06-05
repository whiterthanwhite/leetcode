package trappingrainwater

func trap(height []int) int {
	n := len(height)
	r := 0
	skipHeight := 0
	for left, right := 0, n-1; left < right; {
		if height[left] <= skipHeight {
			left++
			continue
		} else if height[right] <= skipHeight {
			right--
			continue
		}

		min := 0
		if height[left] < height[right] {
			min = height[left]
		} else {
			min = height[right]
		}
		if min > skipHeight {
			skipHeight = min
		}

		for mid := left + 1; mid < right; mid++ {
			if height[mid] < min {
				r += skipHeight - height[mid]
				height[mid] = skipHeight
			}
		}

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return r
}
