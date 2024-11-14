package __assignCookies

import "testing"

func TestFindContentChildren(t *testing.T) {
	g := []int{1, 2, 3}
	s := []int{1, 2}

	t.Log(findContentChildren(g, s))
}
