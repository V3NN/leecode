package __faultyKeyboard

import "testing"

func TestFinalString(t *testing.T) {
	t.Log(finalString("string"))
	t.Log(finalString("poiinter"))
	t.Log(finalString("abcii"))
}
