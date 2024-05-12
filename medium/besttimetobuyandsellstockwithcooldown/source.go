package besttimetobuyandsellstockwithcooldown

import "fmt"

func maxProfit(prices []int) int {
	n := len(prices)
	pricesHeap := make([]int, 1)
	pricesHeap = append(pricesHeap, prices...)
	n = len(pricesHeap)
	edges := make([]*DirectedEdge, 0)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges, &DirectedEdge{
				v:      i,
				w:      j,
				weight: pricesHeap[j],
			})
		}
	}
	for _, edge := range edges {
		fmt.Println(edge)
	}
	return -1
}

type DirectedEdge struct {
	v      int
	w      int
	weight int
}
