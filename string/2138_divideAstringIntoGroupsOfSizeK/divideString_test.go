package __divideAstringIntoGroupsOfSizeK

import "testing"

func TestDivideString(t *testing.T) {
	s, k, fill := "'a'", 3, 'x'
	t.Logf("%s divide a string into groups of size %v = %v", s, k, divideString(s, k, byte(fill)))

	s, k, fill = "'abcd'", 3, 'x'
	t.Logf("%s divide a string into groups of size %v = %v", s, k, divideString(s, k, byte(fill)))

	s, k, fill = "'abcdefghi'", 3, 'x'
	t.Logf("%s divide a string into groups of size %v = %v", s, k, divideString(s, k, byte(fill)))

	s, k, fill = "'xxxx'", 10, '-'
	t.Logf("%s divide a string into groups of size %v = %v", s, k, divideString(s, k, byte(fill)))
}
