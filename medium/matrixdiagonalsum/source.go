package matrixdiagonalsum

func diagonalSum(mat [][]int) int {
	N := len(mat)
	x1 := 0
	x2 := N - 1
	sum := 0
	for y := 0; y < N; y++ {
		sum += mat[x1][y]
		if x1 != x2 {
			sum += mat[x2][y]
		}
		x1++
		x2--
	}
	return sum
}
