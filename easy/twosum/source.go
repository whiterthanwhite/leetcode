package twosum

func twoSum(nums []int, target int) []int {
	vals := make(map[int]int)
	for i, v := range nums {
		if val, ok := vals[target-v]; ok {
			return []int{val, i}
		}
		vals[v] = i
	}
	return nil
}
