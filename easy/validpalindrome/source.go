package validpalindrome

import (
	"strings"
)

func isPalindrome(s string) bool {
	n := len(s)
	s = strings.ToLower(s)
	a := make([]byte, 0)
	for i := 0; i < n; i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9') {
			a = append(a, s[i])
		}
	}

	n = len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if a[i] != a[j] {
			return false
		}
	}

	return true
}
