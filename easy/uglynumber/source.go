package uglynumber

func isUgly(n int) bool {
	switch n {
	case 1, 2, 3, 5:
		return true
	}

	primeFactors := []int{2, 3, 5}
	for n > 0 {
		for _, v := range primeFactors {
			if v*v <= n {
				if n%v == 0 {
					switch n / v {
					case 2, 3, 5:
						return true
					default:
						n = n / v
					}
				}
			}
		}
		t := true
		for _, v := range primeFactors {
			t = t && n%v != 0
		}
		if t {
			break
		}
	}

	return false
}
