package validanagram

func isAnagram(s string, t string) bool {
	lettersInAnagram := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		if t[i] != ' ' {
			_, ok := lettersInAnagram[t[i]]
			if !ok {
				lettersInAnagram[t[i]] = 1
			} else {
				lettersInAnagram[t[i]]++
			}
		}
	}
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			_, ok := lettersInAnagram[s[i]]
			if !ok {
				return false
			} else {
				lettersInAnagram[s[i]]--
			}
		}
	}
	for _, v := range lettersInAnagram {
		if v != 0 {
			return false
		}
	}
	return true
}
