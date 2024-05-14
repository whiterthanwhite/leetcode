package longestpalindromicsubstring

func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	n := len(s)
	if n == 1 {
		return s
	}
	palindrome := ""
	for i := 1; i <= n; i++ {
		for j := i; j >= 0; j-- {
			str := s[j:i]
			if isPalindrome(str) {
				if len(str) > len(palindrome) {
					palindrome = str
				}
			}
		}
	}
	return palindrome
}
