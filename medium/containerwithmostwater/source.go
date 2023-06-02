package containerwithmostwater

func maxArea(height []int) int {
	maxVolume := -1
	volume := 0
	length := 0
	for i, j := 0, len(height)-1; i < j; {
		length = j - i
		if height[i] > height[j] {
			volume = height[j] * length
			j--
		} else {
			volume = height[i] * length
			i++
		}
		if maxVolume < volume {
			maxVolume = volume
		}
	}
	return maxVolume
}
