package calculatedigitsumofastring

import (
	"fmt"
	"strconv"
)

func digitSum(s string, k int) string {
	q1 := &queue{}
	for len(s) > k {
		for i := 0; i < len(s); i++ {
			t, _ := strconv.ParseInt(string(s[i]), 0, 64)
			q1.enqueue(int(t))
		}
		s = ""
		j := 0
		sum := 0
		for !q1.isEmpty() {
			if j < k {
				sum += q1.dequeue()
				j++
			} else {
				j = 0
				s += fmt.Sprint(sum)
				sum = 0
			}
		}
		s += fmt.Sprint(sum)
	}

	return s
}

type node struct {
	item int
	next *node
}

type queue struct {
	first, last *node
	n           int
}

func (q *queue) enqueue(item int) {
	oldLast := q.last
	q.last = &node{
		item: item,
	}
	if q.isEmpty() {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}
	q.n++
}

func (q *queue) dequeue() int {
	item := q.first.item
	q.first = q.first.next
	if q.isEmpty() {
		q.last = nil
	}
	q.n--
	return item
}

func (q *queue) isEmpty() bool {
	return q.first == nil
}
