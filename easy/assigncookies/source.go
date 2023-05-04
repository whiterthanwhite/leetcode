package assigncookies

func findContentChildren(g []int, s []int) int {
	N := len(g)
	M := len(s)
	if N == 0 || M == 0 {
		return 0
	}

	shellSort(g)
	shellSort(s)
	cookiePair := 0
	j := 0
	for i := 0; i < N; i++ {
		if s[j] >= g[i] {
			cookiePair++
			j++
		}
		if j == M {
			break
		}
	}

	return cookiePair
}

func shellSort(a []int) {
	N := len(a)
	h := 1
	for h < N/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j >= h && a[j] > a[j-h]; j -= h {
				t := a[j]
				a[j] = a[j-h]
				a[j-h] = t
			}
		}
		h = h / 3
	}
}
