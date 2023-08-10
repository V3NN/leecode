package __removeKdigits

import "testing"

func TestRemoveKdigits(t *testing.T) {
	t.Log(removeKdigits("1234567890", 9))
	t.Log(removeKdigits("1432219", 3))
	t.Log(removeKdigits("10200", 1))

}
