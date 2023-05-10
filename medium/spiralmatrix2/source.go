package spiralmatrix2

import "fmt"

func spiralMatrix(n int) [][]int {
	spiralM := make([][]int, n)
	for i := 0; i < n; i++ {
		spiralM[i] = make([]int, n)
	}
	xStart, yStart := 0, 0
	xEnd, yEnd := len(spiralM[0])-1, n-1
	val := 1
	x, y := xStart, yStart
	firstHalf := true
	for xStart <= xEnd && yStart <= yEnd {
		if firstHalf {
			if x < xEnd {
				spiralM[y][x] = val
				x++
				val++
			} else if y < yEnd {
				spiralM[y][x] = val
				y++
				val = val + n
			} else {
				yStart++
				xEnd--
				firstHalf = false
			}
		} else {
			if x > xStart {
				spiralM[y][x] = val
				x--
				val--
			} else if y > yStart {
				spiralM[y][x] = val
				y--
				val = val - n
			} else {
				xStart++
				yEnd--
				firstHalf = true
			}
		}
	}
	spiralM[y][x] = val
	return spiralM
}

func spiralMatrix2(n int) [][]int {
	spiralM := make([][]int, n)
	for i := 0; i < n; i++ {
		spiralM[i] = make([]int, n)
	}
	xStart, yStart := 0, 0
	xEnd, yEnd := len(spiralM[0])-1, n-1
	val := 1
	x, y := xStart, yStart
	firstHalf := true
	for xStart <= xEnd && yStart <= yEnd {
		if firstHalf {
			if x < xEnd {
				spiralM[y][x] = val
				x++
				val++
			} else if y < yEnd {
				spiralM[y][x] = val
				y++
				val++
			} else {
				yStart++
				xEnd--
				firstHalf = false
			}
		} else {
			if x > xStart {
				spiralM[y][x] = val
				x--
				val++
			} else if y > yStart {
				spiralM[y][x] = val
				y--
				val++
			} else {
				xStart++
				yEnd--
				firstHalf = true
			}
		}
	}
	spiralM[y][x] = val
	fmt.Println(spiralM)
	return spiralM
}
