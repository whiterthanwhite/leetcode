package sortthestudentsbytheirkthscore

func sortTheStudents(score [][]int, k int) [][]int {
	N := len(score)
	h := 1
	for h < N/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j >= h && score[j][k] > score[j-h][k]; j -= h {
				t := score[j]
				score[j] = score[j-h]
				score[j-h] = t
			}
		}
		h = h / 3
	}

	return score
}
