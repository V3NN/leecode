package __NimGame

import (
	"testing"
)

func TestCanWinNim(t *testing.T) {
	t.Log(canWinNim(5))
	t.Log(canWinNim(4))
	t.Log(canWinNim(3))
}
