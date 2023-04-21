package thekweakestrowsinamatrix

import "fmt"

type rowStrength struct {
	row, strength int
}

func kWeakestRows(mat [][]int, k int) []int {
	N := len(mat)
	rs := make([]rowStrength, N)
	for i := 0; i < N; i++ {
		rank := 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				break
			}
			rank++
		}
		rs[i] = rowStrength{
			row:      i,
			strength: rank,
		}
	}
	fmt.Println(rs)
	shellSort(rs)
	fmt.Println(rs)
	r := make([]int, k)
	for i := 0; i < k; i++ {
		r[i] = rs[i].row
	}
	return r
}

func shellSort(a []rowStrength) {
	N := len(a)
	h := 1
	for h < N/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j >= h && less(a, j, j-h); j -= h {
				exch(a, j, j-h)
			}
		}
		h = h / 3
	}
}

func less(a []rowStrength, i, j int) bool {
	if a[i].strength == a[j].strength {
		return a[i].row < a[j].row
	}
	return a[i].strength < a[j].strength
}

func exch(a []rowStrength, i, j int) {
	t := a[i]
	a[i] = a[j]
	a[j] = t
}
