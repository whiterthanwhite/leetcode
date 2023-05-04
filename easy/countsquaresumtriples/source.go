package countsquaresumtriples

import "math"

func countTriples(n int) int {
	counter := 0
	for n >= 5 {
		t3 := n * n
		for i := 2; i < n; i++ {
			t1 := i * i
			t2 := t3 - t1
			b := int(math.Sqrt(float64(t2)))
			if b*b == t2 {
				counter++
			}
		}
		n--
	}
	return counter
}
