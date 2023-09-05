package __longestValidParentheses

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	t.Log(longestValidParentheses("()()"))  // 4
	t.Log(longestValidParentheses("())))")) // 2
	t.Log(longestValidParentheses(")((((")) // 0
}
