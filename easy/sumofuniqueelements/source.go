package sumofuniqueelements

func sumOfUnique(nums []int) int {
	N := len(nums)
	elements := make([]int, 101)
	for i := 0; i < N; i++ {
		elements[nums[i]]++
	}
	sum := 0
	for i, val := range elements {
		if val == 1 {
			sum += i
		}
	}
	return sum
}
