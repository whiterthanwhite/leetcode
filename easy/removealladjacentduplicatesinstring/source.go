package removealladjacentduplicatesinstring

func removeDuplicates(s string) string {
	if s == "" {
		return s
	}

	arrStr := []byte(s)
	N := len(arrStr)
	for i := 0; i < N-1; i++ {
		if arrStr[i] == arrStr[i+1] {
			newStr := arrStr[:i]
			newStr = append(newStr, arrStr[i+2:]...)
			arrStr = newStr
			N = len(arrStr)
			i = -1
		}
	}
	return string(arrStr)
}
