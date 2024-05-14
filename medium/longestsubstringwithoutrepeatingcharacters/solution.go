package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	sRunes := []rune(s)

	i, j := 0, 1
	l := 0
	usedSymbols := make(map[rune]int)
	for j <= len(s) {
		r := sRunes[j-1]
		y, ok := usedSymbols[r]
		if ok {
			if l < j-i {
				l = j - i - 1
			}
			usedSymbols = make(map[rune]int)
			i = y
			j = i + 1
			continue
		}
		usedSymbols[r] = j
		j++
	}
	if l < j-i {
		l = j - i - 1
	}

	return l
}
