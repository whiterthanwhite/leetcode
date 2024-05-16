package zigzagconversion

type Node struct {
	Value byte
	Next  *Node
}

type Queue struct {
	First *Node
	Last  *Node
}

func (q *Queue) IsEmpty() bool {
	return q.First == nil
}

func (q *Queue) Enqueue(val byte) {
	oldLast := q.Last
	q.Last = &Node{
		Value: val,
	}
	if q.First == nil {
		q.First = q.Last
	} else {
		oldLast.Next = q.Last
	}
}

func (q *Queue) Dequeue() byte {
	r := q.First.Value
	q.First = q.First.Next
	if q.First == nil {
		q.Last = nil
	}
	return r
}

func convert(s string, numRows int) string {
	n := len(s)
	a := [1000]*Queue{}

	i, j := 0, numRows-2
	for y := 0; y < n; y++ {
		if a[i] == nil {
			a[i] = &Queue{}
		}
		if i < numRows {
			a[i].Enqueue(s[y])
			i++
		} else if 1 <= j {
			a[j].Enqueue(s[y])
			j--
		} else {
			i, j = 0, numRows-2
			a[i].Enqueue(s[y])
			i++
		}
	}

	r := make([]byte, n)
	for i, j = 0, 0; i < numRows; i++ {
		if a[i] == nil {
			break
		}
		for !a[i].IsEmpty() {
			r[j] = a[i].Dequeue()
			j++
		}
	}

	return string(r)
}
