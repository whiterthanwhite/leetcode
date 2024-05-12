package addtwonumbers

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	type test struct {
		List1    *List
		List2    *List
		Expected *List
	}

	tests := []test{
		{
			List1:    generateList(2, 4, 3),
			List2:    generateList(5, 6, 4),
			Expected: generateList(7, 0, 8),
		},
		{
			List1:    generateList(0),
			List2:    generateList(0),
			Expected: generateList(0),
		},
		{
			List1:    generateList(9, 9, 9, 9, 9, 9, 9),
			List2:    generateList(9, 9, 9, 9),
			Expected: generateList(8, 9, 9, 9, 0, 0, 0, 1),
		},
		{
			List1:    generateList(9, 9, 9, 9),
			List2:    generateList(9, 9, 9, 9, 9, 9, 9),
			Expected: generateList(8, 9, 9, 9, 0, 0, 0, 1),
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			actual := addTwoNumbers(tt.List1.First, tt.List2.First)
			printList(actual)
		})
	}
}

func generateList(elements ...int) *List {
	l := List{}
	for _, v := range elements {
		l.Enqueue(v)
	}
	return &l
}

func printList(l *ListNode) {
	if l == nil {
		return
	}

	currList := l
	for {
		fmt.Print(currList.Val)
		if currList.Next == nil {
			break
		}
		currList = currList.Next
	}

	fmt.Println()
}
