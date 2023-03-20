package baseballgame

import (
	"strconv"
)

func calPoints(operations []string) int {
	sum := 0
	n := 0
	stack := make([]int, 0)
	for _, v := range operations {
		switch v {
		case "C":
			n--
			stack = stack[:n]
		case "D":
			t := stack[n-1] * 2
			stack = append(stack, t)
			n++
		case "+":
			t := stack[n-2] + stack[n-1]
			stack = append(stack, t)
			n++
		default:
			t, _ := strconv.Atoi(v)
			stack = append(stack, t)
			n++
		}
	}
	for _, v := range stack {
		sum += v
	}
	return sum
}
