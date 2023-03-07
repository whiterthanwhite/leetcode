package addtoarrayformofinteger

import (
	"fmt"
	"strconv"
)

func addToArrayForm(num []int, k int) []int {
	numString := ""
	numString1 := ""
	numString2 := fmt.Sprint(k)

	for _, v := range num {
		numString1 += fmt.Sprint(v)
	}

	N := len(numString1) - len(numString2)
	if N > 0 {
		t := ""
		for i := 0; i < N; i++ {
			t += "0"
		}
		numString2 = t + numString2
	} else if N < 0 {
		N *= -1
		t := ""
		for i := 0; i < N; i++ {
			t += "0"
		}
		numString1 = t + numString1
	}

	addOne := false
	for i := len(numString2) - 1; i >= 0; i-- {
		n1, _ := strconv.ParseInt(string(numString2[i]), 0, 64)
		n2, _ := strconv.ParseInt(string(numString1[i]), 0, 64)
		sum := n1 + n2
		if addOne {
			sum++
			addOne = false
		}
		if sum > 9 {
			sum -= 10
			addOne = true
		}
		numString = fmt.Sprint(sum) + numString
	}
	if addOne {
		numString = "1" + numString
	}

	result := make([]int, len(numString))
	for i, v := range numString {
		t, _ := strconv.ParseInt(string(v), 0, 64)
		result[i] = int(t)
	}
	return result
}
