package spiralmatrix

func spiralOrder(matrix [][]int) []int {
	xStart := 0
	yStart := 0
	xEnd := len(matrix[0]) - 1
	yEnd := len(matrix) - 1
	x := xStart
	y := yStart
	r := make([]int, 0)
	firstHalf := true
	for xStart <= xEnd && yStart <= yEnd {
		if firstHalf {
			if x < xEnd {
				r = append(r, matrix[y][x])
				x++
			} else if y < yEnd {
				r = append(r, matrix[y][x])
				y++
			} else {
				yStart++
				xEnd--
				firstHalf = false
			}
		} else {
			if x > xStart {
				r = append(r, matrix[y][x])
				x--
			} else if y > yStart {
				r = append(r, matrix[y][x])
				y--
			} else {
				xStart++
				yEnd--
				firstHalf = true
			}
		}
	}
	r = append(r, matrix[y][x])
	return r
}
