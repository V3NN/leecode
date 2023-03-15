package __lengthOfLastWord

import (
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	s := "Hello World"
	t.Logf("[%s] = %d", s, lengthOfLastWord(s))

	s = "   fly me   to   the moon  "
	t.Logf("[%s] = %d", s, lengthOfLastWord(s))

	s = "luffy is still joyboy"
	t.Logf("[%s] = %d", s, lengthOfLastWord(s))

	s = " a b "
	t.Logf("[%s] = %d", s, lengthOfLastWord(s))
}
