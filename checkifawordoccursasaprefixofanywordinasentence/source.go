package checkifawordoccursasaprefixofanywordinasentence

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	strs := strings.Split(sentence, " ")
	for i, str := range strs {
		if strings.HasPrefix(str, searchWord) {
			return i + 1
		}
	}
	return -1
}
