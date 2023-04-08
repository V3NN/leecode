package __palindromeNumber

import "testing"

func TestPalindromeNumber(t *testing.T) {
	x := 123
	t.Logf("%d is palindrome number = %v", x, isPalindrome(x))

	x = 1234321
	t.Logf("%d is palindrome number = %v", x, isPalindrome(x))

	x = 0
	t.Logf("%d is palindrome number = %v", x, isPalindrome(x))

	x = 1210
	t.Logf("%d is palindrome number = %v", x, isPalindrome(x))
}
