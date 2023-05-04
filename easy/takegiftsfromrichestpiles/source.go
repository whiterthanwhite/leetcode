package takegiftsfromrichestpiles

import "math"

func pickGifts(gifts []int, k int) int64 {
	N := len(gifts)
	if N == 0 {
		return 0
	}
	pq := &maxpq{
		pq: make([]int, N+1),
	}
	for i := 0; i < N; i++ {
		pq.insert(gifts[i])
	}
	for i := 0; i < k; i++ {
		max := pq.delMax()
		t := math.Sqrt(float64(max))
		pq.insert(int(t))
	}
	var sum int
	for pq.n >= 0 {
		sum += pq.delMax()
	}
	return int64(sum)
}

type maxpq struct {
	pq []int
	n  int
}

func (m *maxpq) insert(key int) {
	m.n++
	m.pq[m.n] = key
	m.swim(m.n)
}

func (m *maxpq) delMax() int {
	max := m.pq[1]
	m.exch(1, m.n)
	m.n--
	m.pq[m.n+1] = 0
	m.sink(1)
	return max
}

func (m *maxpq) swim(k int) {
	for k > 1 && m.less(k/2, k) {
		m.exch(k/2, k)
		k = k / 2
	}
}

func (m *maxpq) sink(k int) {
	for 2*k <= m.n {
		j := 2 * k
		if j < m.n && m.less(j, j+1) {
			j++
		}
		if !m.less(k, j) {
			break
		}
		m.exch(k, j)
		k = j
	}
}

func (m *maxpq) exch(i, j int) {
	t := m.pq[i]
	m.pq[i] = m.pq[j]
	m.pq[j] = t
}

func (m *maxpq) less(i, j int) bool {
	return m.pq[i] < m.pq[j]
}
