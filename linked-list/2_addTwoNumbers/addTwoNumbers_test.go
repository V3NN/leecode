package __addTwoNumbers

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 8, Next: &ListNode{Val: 3}}}
	t.Log(addTwoNumbers(l1, l2))
}
