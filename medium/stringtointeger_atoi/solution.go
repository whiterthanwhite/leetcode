package stringtointegeratoi

import (
	"strconv"
)

type Node struct {
	Val  byte
	Next *Node
}

type Queue struct {
	First *Node
	Last  *Node
	Size  int
}

func (q *Queue) IsEmpty() bool {
	return q.First == nil
}

func (q *Queue) Enqueue(val byte) {
	oldLast := q.Last
	q.Last = &Node{
		Val: val,
	}
	if q.IsEmpty() {
		q.First = q.Last
	} else {
		oldLast.Next = q.Last
	}
	q.Size++
}

func (q *Queue) Dequeue() byte {
	val := q.First.Val
	q.First = q.First.Next
	if q.First == nil {
		q.Last = nil
	}
	q.Size--
	return val
}

func myAtoi(s string) int {
	n := len(s)
	signQueue := &Queue{}
	leadingQueue := &Queue{}
	numberQueue := &Queue{}

	for i := 0; i < n; i++ {
		if s[i] >= '1' && s[i] <= '9' {
			numberQueue.Enqueue(s[i])
		} else if s[i] == '+' || s[i] == '-' {
			if !signQueue.IsEmpty() || !leadingQueue.IsEmpty() || !numberQueue.IsEmpty() {
				break
			}
			signQueue.Enqueue(s[i])
		} else if s[i] == '0' {
			if numberQueue.IsEmpty() {
				leadingQueue.Enqueue(s[i])
				continue
			}
			numberQueue.Enqueue(s[i])
		} else if s[i] == ' ' {
			if !signQueue.IsEmpty() || !leadingQueue.IsEmpty() || !numberQueue.IsEmpty() {
				break
			}
		} else {
			break
		}
	}

	r := make([]byte, signQueue.Size+numberQueue.Size)
	n = 0
	for !signQueue.IsEmpty() {
		r[n] = signQueue.Dequeue()
		n++
	}
	for !numberQueue.IsEmpty() {
		r[n] = numberQueue.Dequeue()
		n++
	}

	v, _ := strconv.ParseInt(string(r), 10, 32)

	return int(v)
}
