package validparentheses

func isValid(s string) bool {
	n := len(s)
	brackets := stack{}
	oppositeBracket := map[byte]byte{')': '(', '}': '{', ']': '['}
	for i := 0; i < n; i++ {
		switch s[i] {
		case '(', '{', '[':
			brackets.Push(s[i])
		default:
			if brackets.IsEmpty() {
				return false
			}
			if oppositeBracket[s[i]] != brackets.Pop() {
				return false
			}
		}
	}
	return brackets.IsEmpty()
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
