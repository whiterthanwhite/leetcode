package numberofsegmentsinstring

import "strings"

func countSegments(s string) int {
	seg := strings.Split(s, " ")
	var c uint16
	for i := 0; i < len(seg); i++ {
		if seg[i] != "" {
			c++
		}
	}
	return int(c)
}
