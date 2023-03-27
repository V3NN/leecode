package __baseballGame

import "testing"

func TestCalPoints(t *testing.T) {
	s := []string{"5", "2", "C", "D", "+"}
	t.Logf("%s CalPoints = %d", s, calPoints(s))

	s = []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	t.Logf("%s CalPoints = %d", s, calPoints(s))

	s = []string{"1"}
	t.Logf("%s CalPoints = %d", s, calPoints(s))
}
