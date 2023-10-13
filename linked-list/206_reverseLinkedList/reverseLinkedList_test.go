package __reverseLinkedList

import "testing"

func TestReverseLinkedList(t *testing.T) {
	l := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	t.Log(reverseList(l))
}
