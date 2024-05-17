package reverseinteger

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	str := []byte(fmt.Sprint(x))
	n := len(str)
	str2 := make([]byte, n)
	i := 0
	j := 1
	if x < 0 {
		i = 1
		j = 0
		str2[0] = '-'
	}

	for ; i < n; i++ {
		str2[n-i-j] = str[i]
	}

	r, err := strconv.ParseInt(string(str2), 10, 32)
	if err != nil {
		return 0
	}
	return int(r)
}
