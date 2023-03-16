package __romanToInt

import "testing"

func TestRomanToInt(t *testing.T) {
	s := "MCMXCIV"
	t.Logf("[%s] to int = %d", s, romanToInt(s))

	s = "III"
	t.Logf("[%s] to int = %d", s, romanToInt(s))

	s = "IV"
	t.Logf("[%s] to int = %d", s, romanToInt(s))

	s = "IX"
	t.Logf("[%s] to int = %d", s, romanToInt(s))

}
