package evaluatereversepolishnotation

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	numbers := new(stack)
	for _, token := range tokens {
		switch token {
		case "+":
			val1 := numbers.Pop()
			val2 := numbers.Pop()
			numbers.Push(val2 + val1)
		case "-":
			val1 := numbers.Pop()
			val2 := numbers.Pop()
			numbers.Push(val2 - val1)
		case "*":
			val1 := numbers.Pop()
			val2 := numbers.Pop()
			numbers.Push(val2 * val1)
		case "/":
			val1 := numbers.Pop()
			val2 := numbers.Pop()
			numbers.Push(val2 / val1)
		default:
			t, _ := strconv.Atoi(token)
			numbers.Push(t)
		}
	}
	return numbers.Pop()
}

type node struct {
	item int
	next *node
}

type stack struct {
	first *node
}

func (s *stack) Push(item int) {
	oldFirst := s.first
	s.first = &node{
		item: item,
		next: oldFirst,
	}
}

func (s *stack) Pop() int {
	item := s.first.item
	s.first = s.first.next
	return item
}
