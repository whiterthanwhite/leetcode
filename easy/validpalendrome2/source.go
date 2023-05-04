package validPalindrome2

import "fmt"

func validPalindrome(s string) bool {
	return isPalindrome(s, false)
}

func isPalindrome(s string, skipped bool) bool {
	N := len(s)
	fmt.Println(s)
	for i := 0; i < N/2; i++ {
		if s[i] != s[N-1-i] {
			if skipped {
				return false
			}
			if isPalindrome(s[:i]+s[i+1:], true) {
				return true
			} else if isPalindrome(s[:N-1-i]+s[N-i:], true) {
				return true
			} else {
				return false
			}
		}
	}
	return true
}
