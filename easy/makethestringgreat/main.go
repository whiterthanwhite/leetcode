package makethestringgreat

import (
	"unicode"
)

func makeGood(s string) string {
	if s == "" {
		return ""
	}
	st := &stack{}

	var isUpper bool
	var currItem byte
	for i := 0; i < len(s); i++ {
		isUpper = unicode.IsUpper(rune(s[i]))
		if isUpper {
			currItem = byte(unicode.ToLower(rune(s[i])))
		} else {
			currItem = s[i]
		}

		if st.isEmpty() {
			st.push(currItem, isUpper)
		} else {
			if isUpper && !st.first.isUpper && currItem == st.first.item {
				st.pop()
			} else if !isUpper && st.first.isUpper && currItem == st.first.item {
				st.pop()
			} else {
				st.push(currItem, isUpper)
			}
		}
	}

	r := make([]byte, st.size())
	var t rune
	var i int
	for !st.isEmpty() {
		i = st.size() - 1
		if st.first.isUpper {
			t = unicode.ToUpper(rune(st.pop()))
			r[i] = byte(t)
		} else {
			r[i] = st.pop()
		}
	}
	return string(r)
}

type node struct {
	item    byte
	isUpper bool
	next    *node
}

type stack struct {
	first *node
	n     int
}

func (s *stack) size() int {
	return s.n
}

func (s *stack) isEmpty() bool {
	return s.n == 0
}

func (s *stack) push(item byte, isUpper bool) {
	oldFirst := s.first
	s.first = &node{
		item:    item,
		next:    oldFirst,
		isUpper: isUpper,
	}
	s.n++
}

func (s *stack) pop() byte {
	item := s.first.item
	s.first = s.first.next
	s.n--
	return item
}
