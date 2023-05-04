package countthedigitsthatdivideanumber

import (
	"fmt"
	"strconv"
)

func countDigits(num int) int {
	strNum := fmt.Sprint(num)
	var count int
	for i := 0; i < len(strNum); i++ {
		t, _ := strconv.Atoi(string(strNum[i]))
		if num%t == 0 {
			count++
		}
	}
	return count
}
