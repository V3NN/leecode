package __isPalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	t.Logf("['%s'] is palindrome = %v", s, isPalindrome(s))

	s = "race a car"
	t.Logf("['%s'] is palindrome = %v", s, isPalindrome(s))

	s = " "
	t.Logf("['%s'] is palindrome = %v", s, isPalindrome(s))

	s = "aBa"
	t.Logf("['%s'] is palindrome = %v", s, isPalindrome(s))

}
