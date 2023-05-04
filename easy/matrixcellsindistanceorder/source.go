package maxtrixcellsindistanceorder

func allCellsDistOrder(rows, cols, rCenter, cCenter int) [][]int {
	answer := make([][]int, rows*cols)
	for i := range answer {
		answer[i] = make([]int, 2)
	}
	edges := &edges{
		e: make([]*edge, len(answer)-1),
	}
	p0 := &point{
		x: rCenter,
		y: cCenter,
	}
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			if p0.x != j || p0.y != i {
				edges.addPoint(p0, &point{x: j, y: i})
			}
		}
	}
	edges.sort()
	answer[0] = []int{p0.x, p0.y}
	for i := 0; i < edges.s; i++ {
		answer[i+1] = []int{edges.e[i].w.x, edges.e[i].w.y}
	}

	return answer
}

func abs(v int) int {
	if v < 0 {
		v *= -1
	}
	return v
}

type point struct {
	x, y int
}

type edge struct {
	v, w   *point
	weight int
}

type edges struct {
	e []*edge
	s int
}

func (e *edges) addPoint(v, w *point) {
	e.e[e.s] = &edge{
		v:      v,
		w:      w,
		weight: abs(v.x-w.x) + abs(v.y-w.y),
	}
	e.s++
}

func (e *edges) sort() {
	h := 1
	for h < e.s/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < e.s; i++ {
			for j := i; j >= h && e.e[j].weight < e.e[j-h].weight; j -= h {
				t := e.e[j]
				e.e[j] = e.e[j-h]
				e.e[j-h] = t
			}
		}
		h = h / 3
	}
}
