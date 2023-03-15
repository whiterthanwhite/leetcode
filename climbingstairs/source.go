package climbingstairs

func climbStairs(n int) int {
	// 1	- 1
	// 2	- 2
	// 3	- 3
	// 4	- 5
	// 5	- 8
	// 6	- 13
	// 7	- 21
	// 8	- 34
	// 9	- 55
	// 10	- 89
	// 11	- 144

	if n > 3 {
		a1, a2 := 2, 3
		for i := 3; i < n; i++ {
			t := a2
			a2 = a1 + a2
			a1 = t
		}
		n = a2
	}

	return n
}
