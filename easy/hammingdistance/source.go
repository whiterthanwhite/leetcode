package hammingdistance

import "fmt"

func hammingDistance(x int, y int) int {
	xstr := toBinaryString(x)
	ystr := toBinaryString(y)
	distanceCounter := 0
	for i := 0; i < len(xstr); i++ {
		if xstr[i] != ystr[i] {
			distanceCounter++
		}
	}
	return distanceCounter
}

func toBinaryString(x int) string {
	s := ""
	padding := 64
	paddingSymb := "0"
	for n := x; n > 0; n /= 2 {
		s = fmt.Sprint(n%2) + s
	}
	for i := len(s); i < padding; i++ {
		s = paddingSymb + s
	}
	return s
}
