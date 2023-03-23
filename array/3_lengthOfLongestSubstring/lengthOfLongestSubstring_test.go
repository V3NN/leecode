package __lengthOfLongestSubstring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	t.Logf("%s lengthOfLongestSubstring = %d", s, lengthOfLongestSubstring(s))

	s = "bbbbb"
	t.Logf("%s lengthOfLongestSubstring = %d", s, lengthOfLongestSubstring(s))

	s = "pwwkew"
	t.Logf("%s lengthOfLongestSubstring = %d", s, lengthOfLongestSubstring(s))
}
