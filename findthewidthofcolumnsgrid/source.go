package findthewidthofcolumnsgrid

import "fmt"

func findColumnWidth(grid [][]int) []int {
	x, y := len(grid[0]), len(grid)
	r := make([]int, x)
	for j := 0; j < x; j++ {
		for i := 0; i < y; i++ {
			tempWidth := len(fmt.Sprint(grid[i][j]))
			if tempWidth > r[j] {
				r[j] = tempWidth
			}
		}
	}
	return r
}
