package __strStr

import "testing"

func TestStrStr(t *testing.T) {
	h, n := "sadbutsad", "sad"
	t.Logf("[%s] , [%s] strstr = %d", h, n, strStr(h, n))

	h, n = "leetcode", "leeto"
	t.Logf("[%s] , [%s] strstr = %d", h, n, strStr(h, n))

	h, n = "abcdefg", "acd"
	t.Logf("[%s] , [%s] strstr = %d", h, n, strStr(h, n))

	h, n = "aabbccd", "d"
	t.Logf("[%s] , [%s] strstr = %d", h, n, strStr(h, n))
}
