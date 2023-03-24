package mostfrequentnumberfollowingkeyinanarray

import "fmt"

func mostFrequent(nums []int, key int) int {
	r := make(map[int]int)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			r[nums[i+1]]++
		}
	}

	count := 0
	value := 0
	for k, v := range r {
		if count < v {
			count = v
			value = k
		}
	}
	fmt.Println(r)
	return value
}
