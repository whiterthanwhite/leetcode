package numberofvalidwordsinasentense

import (
	"strings"
)

func countValidWords(sentence string) int {
	strs := strings.Split(sentence, " ")
	validWords := 0
	for _, str := range strs {
		if validateWord(str) {
			validWords++
		}
	}
	return validWords
}

func validateWord(str string) bool {
	str = strings.Trim(str, " ")
	if str == "" {
		return false
	}
	N := len(str)
	hasPunctuation := false
	hasHyphen := false
	for i := range str {
		r := str[i]
		if r >= 'a' && r <= 'z' {
			if i > 0 && hasPunctuation {
				return false
			}
		} else if r >= '0' && r <= '9' {
			return false
		} else if r == '!' || r == '.' || r == ',' {
			if hasPunctuation {
				return false
			} else {
				hasPunctuation = true
			}
		} else if r == '-' {
			if hasHyphen || N < 3 {
				return false
			} else {
				if i > 0 && i < N-1 {
					if (str[i-1] >= 'a' && str[i-1] <= 'z') && (str[i+1] >= 'a' && str[i+1] <= 'z') {
						hasHyphen = true
					} else {
						return false
					}
				} else {
					return false
				}
			}
		}
	}
	return true
}
