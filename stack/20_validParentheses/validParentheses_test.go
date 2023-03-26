package __validParentheses

import "testing"

func TestValidParentheses(t *testing.T) {
	s := "({[]})"
	t.Logf("'%s' is valid = %v ", s, isValid(s))

	s = "()[]{}"
	t.Logf("'%s' is valid = %v ", s, isValid(s))

	s = "(]"
	t.Logf("'%s' is valid = %v ", s, isValid(s))

	s = "()"
	t.Logf("'%s' is valid = %v ", s, isValid(s))

}
