package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	First *ListNode
	Last  *ListNode
}

func (l *List) Enqueue(val int) {
	oldLast := l.Last
	l.Last = &ListNode{
		Val: val,
	}
	if l.First == nil {
		l.First = l.Last
	} else {
		oldLast.Next = l.Last
	}
}

func (l *List) Dequeue() int {
	val := l.First.Val
	l.First = l.First.Next
	if l.First == nil {
		l.Last = nil
	}
	return val
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	result := &List{}
	sum, reminder := 0, 0
	for {
		if l1 != nil && l2 != nil {
			sum = l1.Val + l2.Val + reminder

			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			sum = l1.Val + reminder

			l1 = l1.Next
		} else if l2 != nil {
			sum = l2.Val + reminder

			l2 = l2.Next
		} else {
			if reminder != 0 {
				result.Enqueue(reminder)
			}
			break
		}

		reminder = sum / 10
		result.Enqueue(sum % 10)
	}

	return result.First
}
