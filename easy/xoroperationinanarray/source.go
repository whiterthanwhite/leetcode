package xoroperationinanarray

func xorOperation(n, start int) int {
	var r int
	for i := 0; i < n; i++ {
		el := start + 2*i
		if i == 0 {
			r = 0 ^ el
		} else {
			r = r ^ el
		}
	}
	return r
}
