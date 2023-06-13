package generateparentheses

import "fmt"

func generateParenthesis(n int) []string {
	r := make([]string, 0)
	s := new(stack)

	var backtrack func(int, int)
	backtrack = func(openN, closedN int) {
		if openN == n && closedN == n && openN == closedN {
			t := s.String()
			fmt.Println(t)
			r = append(r, t)
			return
		}
		if closedN < n {
			s.Push(')')
			backtrack(openN, closedN+1)
			s.Pop()
		}
		if openN < closedN {
			s.Push('(')
			backtrack(openN+1, closedN)
			s.Pop()
		}
	}
	backtrack(0, 0)
	return r
}

type node struct {
	item byte
	next *node
}

type stack struct {
	first *node
}

func (s *stack) IsEmpty() bool {
	return s.first == nil
}

func (s *stack) Push(item byte) {
	oldFirst := s.first
	s.first = &node{
		item: item,
		next: oldFirst,
	}
}

func (s *stack) Pop() byte {
	item := s.first.item
	s.first = s.first.next
	return item
}

func (s *stack) String() string {
	b := make([]byte, 0)
	for n := s.first; n != nil; n = n.next {
		b = append(b, n.item)
	}
	return string(b)
}
