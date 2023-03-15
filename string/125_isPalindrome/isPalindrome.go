package __isPalindrome

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1

	isValid := func(c byte) bool {
		return (c >= 'a' && c < 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
	}

	for l < r {
		for l < r && !isValid(s[l]) {
			l++
		}
		for l < r && !isValid(s[r]) {
			r--
		}
		if l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
	}

	return true
}
