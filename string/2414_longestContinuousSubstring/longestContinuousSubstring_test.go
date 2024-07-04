package __longestContinuousSubstring

import "testing"

func TestLongestContinuousSubstring(t *testing.T) {
	s := "abacaba"
	t.Log(longestContinuousSubstring(s))

	s = "abcde"
	t.Log(longestContinuousSubstring(s))
}
