package dota2senate

func predictPartyVictory(senate string) string {
	N := len(senate)
	votersQueue := queue{}
	for i := 0; i < N; i++ {
		votersQueue.enqueue(senate[i])
	}
	direSkip := 0
	radiantSkip := 0
	for votersQueue.direCount != 0 && votersQueue.radiantCount != 0 {
		voter := votersQueue.dequeue()
		switch voter {
		case 'R':
			if radiantSkip > 0 {
				radiantSkip--
			} else {
				direSkip++
				votersQueue.enqueue(voter)
			}
		case 'D':
			if direSkip > 0 {
				direSkip--
			} else {
				radiantSkip++
				votersQueue.enqueue(voter)
			}
		}
	}

	if votersQueue.radiantCount > 0 {
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
	first        *node
	last         *node
	radiantCount int
	direCount    int
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
	switch item {
	case 'R':
		q.radiantCount++
	case 'D':
		q.direCount++
	}
}

func (q *queue) dequeue() byte {
	item := q.first.item
	switch item {
	case 'R':
		q.radiantCount--
	case 'D':
		q.direCount--
	}
	q.first = q.first.next
	if q.isEmpty() {
		q.last = nil
	}
	return item
}
