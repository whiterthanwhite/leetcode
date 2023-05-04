package mostfrequentnumberfollowingkeyinanarray

import "fmt"

func mostFrequent(nums []int, key int) int {
	r := make(map[int]int)
	target := 0
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			r[nums[i+1]]++
			if r[nums[i+1]] > count {
				target = nums[i+1]
				count = r[nums[i+1]]
			}
		}
	}
	fmt.Println(r)
	return target
}
