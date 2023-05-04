package subtracttheproductandsumofdigitsofaninteger

import (
	"fmt"
	"strconv"
)

func subtractProductAndSum(n int) int {
	numStr := fmt.Sprint(n)
	N := len(numStr)
	prod, _ := strconv.Atoi(string(numStr[0]))
	sum, _ := strconv.Atoi(string(numStr[0]))
	for i := 1; i < N; i++ {
		num, _ := strconv.Atoi(string(numStr[i]))
		prod *= num
		sum += num
	}
	return prod - sum
}
