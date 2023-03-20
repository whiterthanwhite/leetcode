package baseballgame

import (
	"strconv"
)

func calPoints(operations []string) int {
	stack := stack{}
	for _, v := range operations {
		switch v {
		case "C":
			stack.pop()
		case "D":
			x := stack.pop()
			y := x * 2
			stack.push(x)
			stack.push(y)
		case "+":
			x := stack.pop()
			y := stack.pop()
			stack.push(y)
			stack.push(x)
			stack.push(x + y)
		default:
			t, _ := strconv.Atoi(v)
			stack.push(t)
		}
	}
	sum := 0
	for stack.first != nil {
		sum += stack.pop()
	}
	return sum
}

type node struct {
	item int
	next *node
}

type stack struct {
	first *node
	n     int
}

func (s *stack) push(item int) {
	oldFirst := s.first
	s.first = &node{
		item: item,
		next: oldFirst,
	}
	s.n++
}

func (s *stack) pop() int {
	item := s.first.item
	s.first = s.first.next
	s.n--
	return item
}
