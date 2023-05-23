package topkfrequentelements

func topKFrequent(nums []int, k int) []int {
	uniqueNums := make(map[int]int)
	for _, num := range nums {
		uniqueNums[num]++
	}

	minPQ := New(len(nums))
	for key, val := range uniqueNums {
		minPQ.Insert(CountType{key, val})
	}

	result := make([]int, k)
	result[0] = minPQ.DelMin()
	i := 1
	for i < k && !minPQ.IsEmpty() {
		t := minPQ.DelMin()
		if t != result[i-1] {
			result[i] = t
			i++
		}
	}
	return result
}

type CountType struct {
	num   int
	count int
}

type minPQ struct {
	pq []CountType
	n  int
}

func New(maxN int) *minPQ {
	return &minPQ{
		pq: make([]CountType, maxN+1),
	}
}

func (m *minPQ) IsEmpty() bool {
	return m.n == 0
}

func (m *minPQ) Size() int {
	return m.n
}

func (m *minPQ) Insert(key CountType) {
	m.n++
	m.pq[m.n] = key
	m.swim(m.n)
}

func (m *minPQ) DelMin() int {
	max := m.pq[1]
	m.exch(1, m.n)
	m.n--
	m.sink(1)
	return max.num
}

func (m *minPQ) swim(k int) {
	for k > 1 && m.less(k/2, k) {
		m.exch(k/2, k)
		k = k / 2
	}
}

func (m *minPQ) sink(k int) {
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

func (m *minPQ) less(i, j int) bool {
	return m.pq[i].count < m.pq[j].count
}

func (m *minPQ) exch(i, j int) {
	t := m.pq[i]
	m.pq[i] = m.pq[j]
	m.pq[j] = t
}
