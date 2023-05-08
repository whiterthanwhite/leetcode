package matrixdiagonalsum

func diagonalSum(mat [][]int) int {
	xLen := len(mat)
	yLen := len(mat)
	used := make([]int, yLen)
	for i := range used {
		used[i] = -1
	}

	x := 0
	y := 0
	sum := 0
	for x < xLen && y < yLen {
		sum += mat[x][y]
		used[y] = x
		x++
		y++
	}

	x = xLen - 1
	y = 0
	for x >= 0 && y < yLen {
		v := used[y]
		if v != x {
			sum += mat[x][y]
		}
		x--
		y++
	}
	return sum
}
