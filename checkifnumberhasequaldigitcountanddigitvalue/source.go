package checkifnumberhasequaldigitcountanddigitvalue

import (
	"strconv"
)

func digitCount(num string) bool {
	expectCount := make([]int, len(num))
	actualCount := make(map[int]int)

	for i := range num {
		t, _ := strconv.Atoi(string(num[i]))
		expectCount[i] = t
		actualCount[t]++
	}

	for k, v := range expectCount {
		v2, ok := actualCount[k]
		if ok && v != v2 {
			return false
		} else if !ok && v != 0 {
			return false
		}
	}
	return true
}
