package twosum2inputarrayissorted

func twoSum(numbers []int, target int) []int {
	r := make([]int, 2)
	for ind1, v := range numbers {
		key := target - v
		ind2, found := find(key, numbers[ind1+1:])
		if found {
			r[0] = ind1 + 1
			r[1] = ind2 + 1 + r[0]
			break
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
	return 0, false
}
