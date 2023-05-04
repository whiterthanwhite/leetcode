package addtoarrayformofinteger

import (
	"fmt"
)

func addToArrayForm(num []int, k int) []int {
	numString := fmt.Sprint(k)
	var returnSlice []int
	if len(numString) > len(num) {
		returnSlice = make([]int, len(numString)+1)
	} else {
		returnSlice = make([]int, len(num)+1)
	}

	j := len(num) - 1
	for i := len(returnSlice) - 1; i >= 0; i-- {
		returnSlice[i] += k % 10
		k = k / 10
		if j >= 0 {
			returnSlice[i] += num[j]
			j--
		}
		if returnSlice[i] > 9 {
			returnSlice[i] -= 10
			returnSlice[i-1] = 1
		}
	}

	if returnSlice[0] == 0 {
		return returnSlice[1:]
	} else {
		return returnSlice
	}
}
