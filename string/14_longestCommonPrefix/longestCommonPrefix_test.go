package __longestCommonPrefix

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	arr := []string{"dog", "racecar", "car"}
	r := longestCommonPrefix(arr)
	t.Log("['dog', 'racecar', 'car] :", r)

	arr = []string{"flower", "flower", "flight"}
	r = longestCommonPrefix(arr)
	t.Log("['flower', 'flow', 'flight] :", r)

	arr = []string{"aaa", "bbb", "abc"}
	r = longestCommonPrefix(arr)
	t.Log("['aaa', 'bbb', 'abc] :", r)

	arr = []string{"aaabb", "aaabbbc", "aaab"}
	r = longestCommonPrefix(arr)
	t.Log("['aaabb', 'aaabbbc', 'aaab] :", r)
}
