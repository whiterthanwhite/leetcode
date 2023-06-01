// 3sum
package threesum

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	r := make([][]int, 0)
	usedTriplets := make(map[string]struct{})
	sort.Ints(nums)
	fmt.Println(nums)

	for i := 0; i < n-2; i++ {
		for j := n - 1; j > i; j-- {
			key := 0 - nums[i] - nums[j]
			ind, found := find(key, nums[i+1:j])
			if found {
				str := fmt.Sprintf("%d%d%d", nums[i], nums[1+i+ind], nums[j])
				if _, ok := usedTriplets[str]; !ok {
					usedTriplets[str] = struct{}{}
					r = append(r, []int{nums[i], nums[1+i+ind], nums[j]})
				}
			}
		}
	}
	return r
}

func find(key int, a []int) (int, bool) {
	lo, hi := 0, len(a)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if key < a[mid] {
			hi = mid - 1
		} else if key > a[mid] {
			lo = mid + 1
		} else {
			return mid, true
		}
	}
	return -1, false
}
