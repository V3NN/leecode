package __decodeString

import "testing"

func TestDecodeString(t *testing.T) {
	t.Log(decodeString("100[leetcode]"))
	t.Log(decodeString("3[a]2[bc]"))
}
