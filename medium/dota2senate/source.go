package dota2senate

import "fmt"

func predictPartyVictory(senate string) string {
	N := len(senate)
	radiantQueue := queue{}
	direQueue := queue{}

	for i := 0; i < N; i++ {
		switch senate[i] {
		case 'R':
			radiantQueue.enqueue('R')
		case 'D':
			direQueue.enqueue('D')
		}
	}

	direSkip := 0
	radiantSkip := 0
	i := 0
	newSenate := make([]byte, 0)
	for !direQueue.isEmpty() && !radiantQueue.isEmpty() {
		switch senate[i] {
		case 'R':
			if radiantSkip > 0 {
				radiantQueue.dequeue()
				radiantSkip--
			} else {
				direSkip++
				newSenate = append(newSenate, 'R')
			}
		case 'D':
			if direSkip > 0 {
				direQueue.dequeue()
				direSkip--
			} else {
				radiantSkip++
				newSenate = append(newSenate, 'D')
			}
		}
		i++
		if i == N {
			fmt.Println(string(newSenate))
			senate = string(newSenate)
			i = 0
			N = len(senate)
			newSenate = make([]byte, 0)
		}
	}

	if !radiantQueue.isEmpty() {
		return "Radiant"
	} else {
		return "Dire"
	}
}

type node struct {
	item byte
	next *node
}

type queue struct {
	first *node
	last  *node
}

func (q queue) isEmpty() bool {
	return q.first == nil
}

func (q *queue) enqueue(item byte) {
	oldLast := q.last
	q.last = &node{
		item: item,
	}
	if q.isEmpty() {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}
}

func (q *queue) dequeue() byte {
	item := q.first.item
	q.first = q.first.next
	if q.isEmpty() {
		q.last = nil
	}
	return item
}
