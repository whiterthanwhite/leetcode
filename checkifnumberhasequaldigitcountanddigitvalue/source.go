package checkifnumberhasequaldigitcountanddigitvalue

import (
	"strconv"
)

func digitCount(num string) bool {
	expectCount := make(map[int]int, len(num))
	actualCount := make(map[int]int, len(num))

	for i := range num {
		t, _ := strconv.Atoi(string(num[i]))
		expectCount[i] = t
		actualCount[t]++
	}

	for k, v := range expectCount {
		t, ok := actualCount[k]
		if ok {
			if v != t {
				return false
			}
		} else if v != 0 {
			return false
		}
	}
	return true
}
