package maximumnestingdepthoftheparentheses

func maxDepth(s string) int {
	N := len(s)
	p := make([]byte, 0)
	maxDepth := 0
	for i := 0; i < N; i++ {
		switch s[i] {
		case '(':
			p = append(p, s[i])
		case ')':
			l := len(p)
			if l > maxDepth {
				maxDepth = l
			}
			p = p[:l-1]

		}
	}
	return maxDepth
}
