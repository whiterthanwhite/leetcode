package distringmatch

func diStringMatch(s string) []int {
	n := len(s)
	perm := make([]int, n+1)
	min, max := 0, n
	var i int
	for i = 0; i < n; i++ {
		switch s[i] {
		case 'I':
			perm[i] = min
			min++
		case 'D':
			perm[i] = max
			max--
		}
	}
	perm[n] = min
	return perm
}
