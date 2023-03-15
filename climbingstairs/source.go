package climbingstairs

func climbStairs(n int) int {
	// 1 - 1
	// 2 - 2
	// 3 - 3
	// 4 - 5 - 0
	// 5 - 8 - 3
	// 6 - 13 - 5
	// 7 - 21 - 8
	// 8 - 31 - 10
	// 9 - 43 - 12

	if n > 4 {
		a := 5
		for i := 1; i < n-3; i++ {
			a += 3 + (i-1)*2
		}
		n = a
	} else if n == 4 {
		n = 5
	}
	return n
}
