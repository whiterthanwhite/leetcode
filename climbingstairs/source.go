package climbingstairs

func climbStairs(n int) int {
	if n > 3 {
		a1, a2 := 2, 3
		for i := 3; i < n; i++ {
			t := a2
			a2 += a1
			a1 = t
		}
		n = a2
	}

	return n
}
