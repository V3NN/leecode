package __mergeTwoSortedLists

import "testing"

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	t.Logf("[1,2,4] & [1,3,4] merge = %v", mergeTwoLists(l1, l2))
}
