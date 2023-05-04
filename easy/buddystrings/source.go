package buddystrings

func buddyStrings(s string, goal string) bool {
	N := len(s)
	source := []byte(s)
	firstInd, secondInd := -1, -1
	swapCounter := 1
	i := 0
	for i < N {
		if swapCounter == 0 {
			break
		}
		if s[i] == goal[i] {
			i++
			continue
		}
		if firstInd == -1 {
			firstInd = i
			i++
		} else if secondInd == -1 {
			secondInd = i
		} else {
			t := source[firstInd]
			source[firstInd] = source[secondInd]
			source[secondInd] = t
			swapCounter--
			i++
		}
	}
	i, j := 0, 1
	for swapCounter > 0 && i < N-1 {
		if s[i] == goal[j] {
			swapCounter--
			break
		}
		if j == N-1 {
			i++
			j = i + 1
		} else {
			j++
		}
	}

	return string(source) == goal && swapCounter == 0
}
