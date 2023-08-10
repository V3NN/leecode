package __repeatedDnaSequences

import "testing"

func TestFindRepeatedDnaSequences(t *testing.T) {
	t.Log(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))

	t.Log(findRepeatedDnaSequences("AAAAAAAAAAAAA"))
}
